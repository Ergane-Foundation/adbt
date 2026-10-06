# Third-party notices

This project is licensed under the Apache License 2.0. Release binaries are built
with the Go toolchain and statically include the third-party modules below, which
keep their own licenses. The full license texts are in each module's source.

## Go modules compiled into the binary

| Module | License |
|--------|---------|
| Go standard library | BSD 3-Clause (The Go Authors) |
| `golang.org/x/sys`, `golang.org/x/text` | BSD 3-Clause (The Go Authors) |
| `github.com/atotto/clipboard` | BSD 3-Clause |
| `github.com/aymanbagabas/go-osc52/v2` | MIT |
| `github.com/charmbracelet/bubbles` | MIT |
| `github.com/charmbracelet/bubbletea` | MIT |
| `github.com/charmbracelet/colorprofile` | MIT |
| `github.com/charmbracelet/lipgloss` | MIT |
| `github.com/charmbracelet/x/ansi`, `x/cellbuf`, `x/term` | MIT |
| `github.com/clipperhouse/displaywidth` | MIT |
| `github.com/clipperhouse/uax29/v2` | MIT |
| `github.com/erikgeiser/coninput` | MIT |
| `github.com/lucasb-eyer/go-colorful` | MIT |
| `github.com/mattn/go-isatty` | MIT |
| `github.com/mattn/go-localereader` | MIT |
| `github.com/mattn/go-runewidth` | MIT |
| `github.com/muesli/ansi` | MIT |
| `github.com/muesli/cancelreader` | MIT |
| `github.com/muesli/termenv` | MIT |
| `github.com/rivo/uniseg` | MIT |
| `github.com/xo/terminfo` | MIT |

Exact versions are pinned in `go.mod` and `go.sum`.

## Tools adbt runs but does not include

`adb` (Android SDK Platform Tools) and `scrcpy` are installed separately and are
covered by their own licenses.

## Documentation website (`website/`, not part of the binary)

The website is built with Docusaurus and other npm packages, each under its own
license (see `website/package.json`).
