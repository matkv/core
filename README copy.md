# Core

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/matkv/core" alt="Go Version">
  <img src="https://img.shields.io/github/license/matkv/core" alt="License">
  <img src="https://img.shields.io/github/last-commit/matkv/core" alt="Last Commit">
  <a href="https://pkg.go.dev/github.com/matkv/core">
    <img src="https://pkg.go.dev/badge/github.com/matkv/core.svg" alt="Go Reference">
  </a>
</p>

A Go CLI app with a SvelteKit web app. Mainly a project for learning and experimenting with Go - I'm trying to write most of the Go code myself, but in the web app some vibe-coding is allowed 😉.

![Screenshot Web UI](./docs/screenshots/screenshot-main.png)

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js & npm

### Installation

If I want some changes in the web app to also be included in the binary, first build the SvelteKit app & commit the changes from the build directory:

```bash
make build-web
git add internal/ui/build
```

Install the `core` binary globally so you can run it from anywhere:

```bash
go install github.com/matkv/core@latest
```
Here the assets for the web UI are embedded into the Go binary using `embed.FS`. So when you run `core serve`, it serves the embedded SvelteKit app.

This installs the binary to `$GOPATH/bin` (usually `~/go/bin`). Ensure this directory is in your `PATH`.

Alternatively, you can just use the latest binary from the [Releases](https://github.com/matkv/core/releases).

Then run commands like:
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
- **Debug Core CLI:** Runs only the CLI under the debugger. Use this to debug the actual CLI commands.
