
```md
# go-example-api

A simple example API built with Go.

## Running the API

### Without Hot Reload

```bash
go run ./cmd/api/*.go
```

### With Hot Reload (using Air)

1. **Install Air**

```bash
go install github.com/air-verse/air@latest
```

Make sure `$GOPATH/bin` is in your `PATH`. You can add this to your `~/.zshrc` or `~/.bashrc`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then reload your shell:

```bash
source ~/.zshrc  # or ~/.bashrc
```

2. **Initialize Air**

Run this once to generate the config file:

```bash
air init
```

3. **Start the API with hot reload**

```bash
air
```

This will watch for file changes and automatically restart the server.

## Available Endpoints

### Health Check

```bash
curl http://localhost:8080/v1/health
```
```
