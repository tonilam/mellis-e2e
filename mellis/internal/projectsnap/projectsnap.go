// Package projectsnap snapshots a project tree so agent CLI writes can be
// reviewed, accepted, or discarded.
//
// Assumption (CLARIFICATION_NEEDED until product confirms): agent CLIs write
// in-place in the project path. The runner snapshots before the request, lets
// the CLI write, then restores the project and holds the written tree as a
// proposal until accept or discard. After the request finishes and before
// accept, nothing from the request remains in the project.
//
// Snapshots live under the system temp directory (not inside the project).
// .git is skipped so restore does not thrash git metadata.
package projectsnap

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Kind is how one path differs between before and after.
type Kind string

const (
	KindAdded    Kind = "added"
	KindModified Kind = "modified"
	KindRemoved  Kind = "removed"
)

// Change is one proposed file difference relative to the project root.
type Change struct {
	Path string `json:"path"`
	Kind Kind   `json:"kind"`
}

// Snap is a captured copy of a project tree.
type Snap struct {
	Dir string
}

// Capture copies projectPath into a new temp directory and returns it.
func Capture(projectPath string) (*Snap, error) {
	projectPath = strings.TrimSpace(projectPath)
	if projectPath == "" {
		return nil, fmt.Errorf("project path is empty")
	}
	info, err := os.Stat(projectPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("project path is not a directory: %s", projectPath)
	}
	dir, err := os.MkdirTemp("", "mellis-snap-*")
	if err != nil {
		return nil, err
	}
	if err := copyTree(projectPath, dir); err != nil {
		_ = os.RemoveAll(dir)
		return nil, err
	}
	return &Snap{Dir: dir}, nil
}

// Close removes the snapshot directory.
func (s *Snap) Close() error {
	if s == nil || s.Dir == "" {
		return nil
	}
	err := os.RemoveAll(s.Dir)
	s.Dir = ""
	return err
}

// Restore makes projectPath match the snapshot (files only; .git skipped).
func (s *Snap) Restore(projectPath string) error {
	if s == nil || s.Dir == "" {
		return fmt.Errorf("empty snapshot")
	}
	return replaceTree(s.Dir, projectPath)
}

// DiffAgainst compares the snapshot to the live project tree.
func (s *Snap) DiffAgainst(projectPath string) ([]Change, error) {
	if s == nil || s.Dir == "" {
		return nil, fmt.Errorf("empty snapshot")
	}
	before, err := listFiles(s.Dir)
	if err != nil {
		return nil, err
	}
	after, err := listFiles(projectPath)
	if err != nil {
		return nil, err
	}
	var out []Change
	for rel, beforeInfo := range before {
		afterInfo, ok := after[rel]
		if !ok {
			out = append(out, Change{Path: rel, Kind: KindRemoved})
			continue
		}
		same, err := sameFile(filepath.Join(s.Dir, rel), beforeInfo, filepath.Join(projectPath, rel), afterInfo)
		if err != nil {
			return nil, err
		}
		if !same {
			out = append(out, Change{Path: rel, Kind: KindModified})
		}
	}
	for rel := range after {
		if _, ok := before[rel]; !ok {
			out = append(out, Change{Path: rel, Kind: KindAdded})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Path == out[j].Path {
			return out[i].Kind < out[j].Kind
		}
		return out[i].Path < out[j].Path
	})
	return out, nil
}

// Apply copies proposedDir onto projectPath (proposal → project).
func Apply(proposedDir, projectPath string) error {
	proposedDir = strings.TrimSpace(proposedDir)
	projectPath = strings.TrimSpace(projectPath)
	if proposedDir == "" || projectPath == "" {
		return fmt.Errorf("proposed and project paths are required")
	}
	return replaceTree(proposedDir, projectPath)
}

func replaceTree(srcRoot, dstRoot string) error {
	dstFiles, err := listFiles(dstRoot)
	if err != nil {
		return err
	}
	srcFiles, err := listFiles(srcRoot)
	if err != nil {
		return err
	}
	// Remove files in dst that are not in src.
	for rel := range dstFiles {
		if _, ok := srcFiles[rel]; ok {
			continue
		}
		if err := os.RemoveAll(filepath.Join(dstRoot, rel)); err != nil {
			return err
		}
	}
	return copyTree(srcRoot, dstRoot)
}

func copyTree(srcRoot, dstRoot string) error {
	return filepath.WalkDir(srcRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(srcRoot, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if shouldSkip(rel, d) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		dst := filepath.Join(dstRoot, rel)
		if d.IsDir() {
			return os.MkdirAll(dst, 0o755)
		}
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return err
		}
		return copyFile(path, dst)
	})
}

func listFiles(root string) (map[string]os.FileInfo, error) {
	out := make(map[string]os.FileInfo)
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if shouldSkip(rel, d) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		out[filepath.ToSlash(rel)] = info
		return nil
	})
	return out, err
}

func shouldSkip(rel string, d os.DirEntry) bool {
	base := filepath.Base(rel)
	if base == ".git" {
		return true
	}
	return false
}

func sameFile(pathA string, a os.FileInfo, pathB string, b os.FileInfo) (bool, error) {
	if a.Size() != b.Size() {
		return false, nil
	}
	fa, err := os.Open(pathA)
	if err != nil {
		return false, err
	}
	defer fa.Close()
	fb, err := os.Open(pathB)
	if err != nil {
		return false, err
	}
	defer fb.Close()
	bufA := make([]byte, 32*1024)
	bufB := make([]byte, 32*1024)
	for {
		na, ea := fa.Read(bufA)
		nb, eb := fb.Read(bufB)
		if na > 0 || nb > 0 {
			if na != nb || string(bufA[:na]) != string(bufB[:nb]) {
				return false, nil
			}
		}
		if ea == io.EOF && eb == io.EOF {
			return true, nil
		}
		if ea != nil && ea != io.EOF {
			return false, ea
		}
		if eb != nil && eb != io.EOF {
			return false, eb
		}
		if ea == io.EOF || eb == io.EOF {
			return na == 0 && nb == 0 && ea == eb, nil
		}
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	info, err := in.Stat()
	if err != nil {
		return err
	}
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode().Perm())
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
