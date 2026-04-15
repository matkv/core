# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this project is

`core` is a personal Go CLI "pocket knife" that also serves an embedded SvelteKit dashboard. Cobra powers the command surface (`cmd/`), Viper loads user config from `~/.config/core/config.yaml`, and the `serve` command starts an HTTP server that exposes JSON endpoints and serves the bundled frontend from `internal/ui/build`.

## Build commands

```bash
# Full build (web + Go) and run
make run

# Build steps separately
make build-web   # npm install + npm run build + copy to internal/ui/build
make build-go    # go build -o bin/core .

# Go only
go test ./...    # verify compilation (no tests yet)
go run . <command>

# Frontend only (from web/)
npm install
npm run dev      # Vite dev server at http://localhost:5173
npm run build    # output to web/build/
npm run check    # svelte-check (TypeScript)
npm run lint     # Prettier + ESLint
```

## Config file requirement

Every command calls `config.Load()`, which interactively prompts to create `~/.config/core/config.yaml` if missing and **panics** if declined. Pre-create the file before running any command:

```bash
mkdir -p ~/.config/core && cat > ~/.config/core/config.yaml << 'EOF'
paths:
  obsidianvault: /home/YOU/documents/Obsidian Vault
device: desktop   # desktop | laptop | wsl
EOF
```

Set `XDG_CONFIG_HOME` to override the default config location for isolated testing.

## Architecture

**Command registration** (`cmd/root.go`): Commands are registered via `addCommand(cmd, devices...)` — only commands matching the current `config.Device` are visible. Use the `config.Device` enum (`Desktop`, `Laptop`, `WSL`) when adding new commands.

**Adding a command**: Create a file in `cmd/`, define a `var xxxCmd = &cobra.Command{...}`, and register it in `cmd/root.go` via `addCommand`.

**HTTP server** (`internal/app/server.go`): Standard library `net/http` with three endpoints: `/api/hello`, `/api/random`, `/api/settings`. Push logic into `internal/*` packages; keep handlers thin.

**Embedded UI** (`internal/ui/embed.go`): Go `embed` bundles `internal/ui/build` (copied from `web/build`). Never hand-edit files under `internal/ui/build` — always rebuild via `make build-web`. After frontend changes meant for embedding: `make build-web && git add internal/ui/build`.

**Frontend** (`web/src/`): SvelteKit with static adapter. Use `$lib` imports and `src/lib/state/widgetState.ts` for cross-route state. Styling via Tailwind utility classes in `app.css`.

## Key conventions

- Favor returning errors over panicking, except at `cmd.Execute()`.
- Never hardcode config paths — use `internal/config` helpers.
- Each command file owns its own `var xxxCmd`. Shared utilities live in `internal/*`.
- `gofmt`/`goimports` for Go; Prettier config in `web/` for frontend.
- Run `npm run check && npm run lint` before committing frontend changes.

## Known gotchas

- `cmd/browser` reads `links.txt` from the config directory; it has no URL validation yet.
- Release tags (`v*`) must include freshly built UI assets — run `make build-web` and commit `internal/ui/build` before tagging.
- Node ≥20 required; `web/.npmrc` has `engine-strict=true` and will abort on version mismatch.
- The `serve` command skips auto-opening the browser when `device: wsl`.
