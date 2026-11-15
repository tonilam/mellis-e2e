<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Home — {{ config('app.name') }}</title>
    <style>
        body { font-family: system-ui, sans-serif; max-width: 36rem; margin: 4rem auto; padding: 0 1rem; }
        button { margin-top: 1rem; padding: 0.5rem 1rem; }
    </style>
</head>
<body>
    <h1 data-testid="home-heading">Welcome, {{ auth()->user()->name }}</h1>
    <p data-testid="authenticated-banner">You are signed in.</p>
    <form method="POST" action="{{ route('logout') }}">
        @csrf
        <button type="submit" data-testid="logout-submit">Sign out</button>
    </form>
</body>
</html>
