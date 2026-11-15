package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	application.RegisterEvent[SuiteState](suiteStateEvent)
	application.RegisterEvent[AgentState](agentStateEvent)
}

func main() {
	svc := NewApp()

	app := application.New(application.Options{
		Name:        "Mellis",
		Description: "Local E2E runner",
		Services: []application.Service{
			application.NewService(svc),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	svc.setApplication(app)

	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Mellis",
		Width:  1024,
		Height: 768,
		BackgroundColour: application.NewRGB(245, 246, 248),
		URL:              "/",
	})
	window.RegisterHook(events.Common.WindowClosing, func(e *application.WindowEvent) {
		svc.beforeClose()
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
