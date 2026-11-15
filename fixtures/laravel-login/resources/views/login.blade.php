<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Login — {{ config('app.name') }}</title>
    <style>
        body { font-family: system-ui, sans-serif; max-width: 24rem; margin: 4rem auto; padding: 0 1rem; }
        label { display: block; margin-top: 1rem; font-weight: 600; }
        input { width: 100%; padding: 0.5rem; margin-top: 0.25rem; box-sizing: border-box; }
        button { margin-top: 1.25rem; padding: 0.5rem 1rem; }
        .error { color: #b91c1c; margin-top: 0.5rem; font-size: 0.875rem; }
    </style>
</head>
<body>
    <h1>Login</h1>
    <form method="POST" action="{{ route('login') }}" data-testid="login-form">
        @csrf
        <label for="username">Username</label>
        <input
            id="username"
            name="username"
            type="text"
            value="{{ old('username') }}"
            autocomplete="username"
            required
            autofocus
            data-testid="username"
        >
        @error('username')
            <p class="error" data-testid="login-error">{{ $message }}</p>
        @enderror

        <label for="password">Password</label>
        <input
            id="password"
            name="password"
            type="password"
            autocomplete="current-password"
            required
            data-testid="password"
        >

        <button type="submit" data-testid="login-submit">Sign in</button>
    </form>
</body>
</html>
