---
sidebar_position: 6
---

# Contributing

Want to help improve adbt? Contributions are welcome — from documentation fixes to new screens.

## Where to start

1. Read [CONTRIBUTING.md](https://github.com/SakshhamTheCoder/adbt/blob/main/CONTRIBUTING.md) — it covers setup on Windows, macOS and Linux, the checks your change must pass, and how pull requests are reviewed.
2. Look for issues labelled [`good first issue`](https://github.com/SakshhamTheCoder/adbt/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) — they are scoped to be approachable for new contributors.
3. By taking part you agree to follow the [Code of Conduct](https://github.com/SakshhamTheCoder/adbt/blob/main/CODE_OF_CONDUCT.md).

## Quick setup summary

You need Go (see `go.mod` for the version) and a recent `adb` (Android SDK Platform Tools) on your `PATH`. Working on this documentation website also needs Node.js 20 or newer.

```sh
git clone https://github.com/<your-username>/adbt.git
cd adbt
go build ./cmd/adbt
./adbt
```

To preview the documentation website locally:

```sh
cd website
npm install
npm start
```

Before opening a pull request, run `gofmt -l .` (must print nothing), `go vet ./...` and `go test ./...`, and write commit messages in [Conventional Commits](https://www.conventionalcommits.org) style. The full checklist lives in [CONTRIBUTING.md](https://github.com/SakshhamTheCoder/adbt/blob/main/CONTRIBUTING.md).

## Getting help

- **Questions and ideas:** see [SUPPORT.md](https://github.com/SakshhamTheCoder/adbt/blob/main/SUPPORT.md) for where to ask.
- **Bugs:** open an issue with your adbt version (`adbt --version`), operating system, phone model and Android version, and the steps to reproduce.
- **Security problems:** do not post them publicly — follow `SECURITY.md`.
