# Core – Copilot Onboarding

## What this project does
- `core` is a Go-based personal CLI “pocket knife” that can also serve an embedded SvelteKit dashboard. Cobra powers the command surface (`cmd/`), while Viper loads user config from `~/.config/core/config.yaml`. The `serve` command spins up an HTTP server (`internal/server`) that exposes a few JSON endpoints and serves the bundled frontend assets from `internal/ui/build`.

## Tech stack & tooling
- **Languages:** Go 1.25.x for the CLI/server, TypeScript/Svelte 5 + Tailwind via Vite for the web UI.
- **Key deps:** `spf13/cobra`, `spf13/viper`, `embed` FS for bundling, SvelteKit CLI (`sv`-generated scaffold), ESLint 9 + `typescript-eslint`, Prettier, Tailwind 4.
- **Build tools:** `make` orchestrates `build-web` (runs `npm install && npm run build` then copies to `internal/ui/build`) and `build-go` (`go build -o bin/core .`). `make run` chains both then launches `./bin/core serve`.
- **CI:** `.github/workflows/release-on-tag.yml` builds when a `v*` tag is pushed, using Node 25.2.0 and Go 1.25 to produce Linux/Windows binaries and publish a release.
- **Editor setup:** `.vscode/launch.json` has launch configs for debugging the CLI, the `serve` command, and the web app (Chrome + node-terminal dev server).

## Repository layout
```
cmd/               Individual Cobra commands (`browser`, `random`, `serve`, `version`, ...).
internal/config    Config loading + device enum helpers.
internal/random    Tiny random utilities used by CLI and API.
internal/server    HTTP mux and API endpoints; wires in UI assets.
internal/ui/build  Generated SvelteKit static bundle (copied from `web/build`).
web/               SvelteKit frontend (Vite config, src/lib, routes, Tailwind, ESLint/Prettier).
docs/              Currently only screenshots used in README.
Makefile           Primary build workflow.
main.go            Entrypoint calling `cmd.Execute()`.
```
> `internal/ui/build` is generated output—never hand-edit it. Rebuild via `make build-web`.

## Build, run, and test workflow
1. **Go tooling**
   - Install Go 1.25+.
   - `go test ./...` is the fastest way to verify CLI/server packages (no tests yet but it ensures everything compiles).
   - `go run . <command>` executes the CLI directly. Every command calls `config.Load()`, which will interactively ask to create `~/.config/core/config.yaml` if it is missing and will panic with `panic: Config file creation was declined by the user` if you answer anything but `y/yes` or provide no input. Pre-create the file (copy the snippet from `README.md`), or script it:

     ```bash
     mkdir -p ~/.config/core &&
     cat <<'EOF' > ~/.config/core/config.yaml
     paths:
       obsidianvault: /home/YOU/documents/Obsidian Vault
     device: desktop
     EOF
     ```

     For isolated tests you can set `XDG_CONFIG_HOME` to point at a temp directory that already contains this file.
   - Add new commands by creating a file under `cmd/`, defining a `cobra.Command`, and registering it via `addCommand` in `cmd/root.go`, specifying allowed devices.
2. **Frontend tooling**
   - Requires Node ≥20 (CI uses 25.2.0). `web/.npmrc` has `engine-strict=true`, so mismatched Node versions will abort installs; align with the version listed in CI or `.nvmrc` if you add one.
   - Run `npm install` once inside `web/`; repeated `make build-web` calls already re-run `npm install`, so local dev is faster if you do it manually first.
   - `npm run dev` starts Vite + SvelteKit; `npm run check` runs `svelte-check`; `npm run lint` runs Prettier check + ESLint; `npm run build` outputs to `web/build/`.
   - After frontend changes meant for Go embedding, run `make build-web && git add internal/ui/build`.
3. **Full app run**
   - `make run` (or `make build-web && make build-go && ./bin/core serve`) keeps Go + web outputs in sync and launches the server at `http://localhost:8080`, opening a browser automatically unless the configured device is `wsl`.

## Coding guidelines & conventions
- **General**
  - Keep edits surgical; follow existing file/module boundaries (`cmd` for CLI surface, `internal/*` for reusable logic, `web` for UI). Use `gofmt`/`goimports` and SvelteKit’s Prettier config.
  - Favor returning errors instead of panicking except at the CLI entry (`cmd.Execute` already panics on fatal errors).
  - Use the `config.Device` enum when gating commands (`addCommand(..., config.Desktop, ...)`).
  - When reading/writing config, reuse `internal/config` helpers—never hardcode paths outside them.
- **Go CLI/server**
  - Each command file defines a `var xxxCmd = &cobra.Command{...}`. Shared helpers (e.g., `OpenURL`) live near consumers.
  - Server routes live in `internal/server/server.go`; keep API handlers minimal and push logic into `internal/*` packages for reuse/testing.
- **Web**
  - Use `$lib` exports + `src/lib/state/widgetState.ts` for cross-route persistence (see `settings` and `journal` routes).
  - Styling relies on Tailwind utility classes defined in `app.css`; keep typography consistent with existing tokens.
  - Run `npm run check && npm run lint` before committing frontend work to catch TypeScript or ESLint issues early.
- **Generated artifacts**
  - Do not edit files under `.svelte-kit/` or `internal/ui/build`; rebuild instead.

## Existing tools & resources
- `README.md` (project overview + config instructions), `docs/screenshots` (UI reference), `Makefile`, `.vscode/launch.json`, `web` lint/format configs, GitHub Actions release workflow.
- No `CONTRIBUTING.md` yet, so follow this file plus the README.
- Ripgrep (`rg`) is not installed in this environment; use `grep -R` or install `rg` before relying on it.

## Known gaps / TODOs / gotchas
- `cmd/browser` expects `links.txt` next to the config file and currently lacks URL validation (`// TODO add check for valid URLs`); plan to guard against malformed input if you touch that command.
- `internal/config.EnsureConfigFileExists` prompts on stdin; automation must supply `y/yes` or create the config ahead of time to avoid blocked builds/tests. If you answer “no” you’ll see `panic: Config file creation was declined by the user`.
- Random number seeding uses `time.Now()` inside each call; avoid calling it in tight loops if deterministic results are needed.
- Release tags (`v*`) must include freshly built UI assets (`make build-web`) before tagging, otherwise the embedded frontend will be stale.

## Quick start checklist
1. Clone the repo, install Go 1.25+, Node ≥20, npm.
2. Inside `web/`: `npm install`.
3. Create `~/.config/core/config.yaml` (copy the YAML from `README.md` and adjust paths/device).
4. For backend work: `go test ./...` or `go run . random`.
5. For frontend work: `npm run dev` (frontend only) or `make run` (backend + embedded UI).
6. Keep `internal/ui/build` in sync whenever shipping frontend changes.
