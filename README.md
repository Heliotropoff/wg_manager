# wg_manager
Simple app for managing wg server through CLI application or REST Api

## Configuration

Create a local `.env` file from `.env.example` and replace the placeholder with
a strong random secret:

```dotenv
API_KEY=replace-with-a-strong-random-secret
```

The local `.env` file is ignored by Git. If it cannot be loaded or `API_KEY` is
empty, the server still starts, but protected endpoints remain unavailable.

## API

`/hello` is public. `/profile` requires the API key in the `X-API-Key` header:

```sh
curl -H "X-API-Key: replace-with-a-strong-random-secret" http://localhost:8080/profile
```

A missing or incorrect key returns `401 Unauthorized`.
