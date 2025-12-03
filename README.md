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

#### 1. Production
To build the SvelteKit app, embed it into the Go binary, and run the server:

```bash
make run
```
The server will start at `http://localhost:8080`.

#### 2. Development Mode

In VS Code, pick: **Dev: Go + Svelte (full debug)** in the Run and Debug panel.

This starts the Go API and the Svelte dev server, and opens a Chrome window with debugging enabled. Breakpoints work across routes (including `settings/+page.svelte`).

## Debugging (VS Code)

- **Dev: Go + Svelte (full debug):** Runs the Go server and Svelte dev server, opens Chrome with debugger attached. Use this to debug both backend and frontend together.
- **Debug Core CLI:** Runs the standalone CLI `cmd/core/main.go` under the debugger. Use this to debug the actual CLI commands.
