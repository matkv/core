# Core

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/matkv/core" alt="Go Version">
  <img src="https://img.shields.io/github/license/matkv/core" alt="License">
  <img src="https://img.shields.io/github/last-commit/matkv/core" alt="Last Commit">
  <a href="https://pkg.go.dev/github.com/matkv/core">
    <img src="https://pkg.go.dev/badge/github.com/matkv/core.svg" alt="Go Reference">
  </a>
</p>

Re-writing this whole project. The general idea is to have a CLI tool like before, but this time with an additional command that starts web server that runs a little web UI for which I'll probably use SvelteKit.

## Project Structure

```text
core/
├── bin/                   // Compiled binaries
├── cmd/
│   └── core/
│       └── main.go        // Very thin entry point, just calls cli.Execute()
│
├── internal/
│   ├── cli/               // Cobra command definitions
│   │   ├── root.go        // Root command setup
│   │   ├── serve.go       // "serve" command logic
│   │   └── version.go     // "version" command logic
│   │
│   ├── server/            // HTTP Server logic
│   │   ├── server.go      // Server struct and startup
│   │
│   └── ui/                // Handles the embedded frontend
│       └── embed.go       // //go:embed web/build
│
├── web/                   // SvelteKit project
│   ├── src/
│   ├── build/             // The output directory (embedded by Go)
│   └── package.json
│
├── Makefile               // Critical: runs `npm build` inside web/ before `go build`
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js & npm

### Installation

To install the `core` binary globally so you can run it from anywhere:

```bash
go install ./cmd/core
```

This will install the binary to your `$GOPATH/bin` (usually `~/go/bin`). Make sure this directory is in your `PATH`.

After installation, you can run commands like:
```bash
core version
core serve
```

### Running the Application

#### 1. Production Mode (Integrated)
To build the SvelteKit app, embed it into the Go binary, and run the server:

```bash
make run
```
The server will start at `http://localhost:8080`.

#### 2. Development Mode

For trying out changes to the SvelteKit frontend without rebuilding the Go binary each time, run the Go API server and SvelteKit dev server separately:

1. Start the Go API server:
   ```bash
   go run ./cmd/core serve
   ```
   
   Or if you've installed the binary:
   ```bash
   core serve
   ```

   *Note: The API runs on port 8080.*

2. In a separate terminal, start the SvelteKit dev server:
   ```bash
   cd web
   npm run dev
   ```
   *Note: The UI runs on http://localhost:5173.*
