# Core

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/matkv/core" alt="Go Version">
  <img src="https://img.shields.io/github/license/matkv/core" alt="License">
  <img src="https://img.shields.io/github/last-commit/matkv/core" alt="Last Commit">
  <a href="https://pkg.go.dev/github.com/matkv/core">
    <img src="https://pkg.go.dev/badge/github.com/matkv/core.svg" alt="Go Reference">
  </a>
</p>

A Go CLI app to automate some personal tasks - mainly a project for learning & experimenting with Go. This is meant to be a little pocket-knife for various random tasks I'd like to run in the terminal so I don't really mind if the features are a bit bloated or unfocused. It uses [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper) for the CLI and [SvelteKit](https://kit.svelte.dev/) for a web UI. I'm trying to write most of the Go code myself, but in the web app some vibe-coding is allowed 😉.

![Screenshot Web UI](./docs/screenshots/screenshot-main.png)

## Features

- [x] CLI & config management with Cobra & Viper
- [x] Embedded SvelteKit web UI
- [x] `core browser` command to open multiple pre-defined URLs in the default browser
- [ ] System tray mode
- [ ] `core pick` command to pick one option from multiple provided arguments
- [ ] `core journal` command to add journal entries to my Obisidan vault from the CLI or web UI


## Installation



## Development