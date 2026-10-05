# ASCII Galaxy Go

A Go rewrite of ASCII Galaxy, built as a learning project for terminal rendering, HTTP webhooks, and concurrency.

The goal is an animated terminal galaxy with a clock and full-screen celebrations showing the repository and branch when a Git push arrives.

## Status

Early work in progress. Configuration loading, terminal size detection, and screen-buffer structures are in place. Animation, clock rendering, and working Gitea/GitLab webhook integrations are planned.

The project currently does not compile: `main.go` calls the undefined `CalculateStarPositions` function. The existing layout helper is named `CalculateFormattingPositions`. The webhook handler and its registration are also commented out.

## Development setup

```sh
git clone git@github.com:fatpotatosalad/ascii-galaxy-go.git
cd ascii-galaxy-go
cp config.toml.example config.toml
go mod download
```

`config.toml` is local configuration and is ignored by Git. Use dummy secrets during development: the current startup code prints the configured secrets, which must be removed before using real credentials.

The configuration contains webhook secrets and a server address and port. These are shared webhook secrets, not API keys. The server currently listens on `:9999`; using the configured host and port is still to be implemented.

After resolving the build blocker, run from the project directory so the program can find `config.toml`:

```sh
go run .
```

To build a binary:

```sh
go build -o ascii-galaxy-go .
```

## Project layout

| File | Purpose |
| --- | --- |
| `main.go` | Application startup and HTTP server setup |
| `config.go` | TOML configuration types and loading |
| `galaxy.go` | Terminal sizing, layout, and screen-buffer foundations |
| `webhook.go` | Draft webhook payload types and handler |
| `config.toml.example` | Example local configuration |

## Roadmap

- [ ] Fix startup and validate configuration.
- [ ] Implement and test signed Gitea push webhooks.
- [ ] Implement and test GitLab push webhooks separately.
- [ ] Render changed terminal cells without clearing every frame.
- [ ] Add galaxy animation, clock, and responsive layout.
- [ ] Send push events to the renderer through a channel.
- [ ] Add explosion effects and full-screen push information.
- [ ] Handle terminal resizing, shutdown, and kiosk deployment.
