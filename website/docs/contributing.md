---
sidebar_position: 5
---

# Contributing

Thanks for helping improve `adbt`.

This project is a community effort, and we welcome bug fixes, feature work, and documentation improvements.

## Start here

- Read the contributor guide in [CONTRIBUTING.md](https://github.com/SakshhamTheCoder/adbt/blob/main/CONTRIBUTING.md) for setup, checks, and pull-request rules.
- Need help or want to ask a question? See [SUPPORT.md](https://github.com/SakshhamTheCoder/adbt/blob/main/SUPPORT.md).
- Looking for a beginner-friendly place to start? Browse the open [good first issues](https://github.com/SakshhamTheCoder/adbt/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22).
- By taking part you agree to follow the [Code of Conduct](https://github.com/SakshhamTheCoder/adbt/blob/main/CODE_OF_CONDUCT.md).

## Local setup

You need Go (see `go.mod` for the version) and a recent `adb` (Android SDK Platform Tools) on your `PATH`. Working on this documentation website also needs Node.js 20 or newer.

```bash
git clone https://github.com/<your-username>/adbt.git
cd adbt
go build ./cmd/adbt
./adbt
```

To preview the documentation website locally:

```bash
cd website
npm install
npm start
```

Before opening a pull request, run `gofmt -l .` (must print nothing), `go vet ./...` and `go test ./...`, and write commit messages in [Conventional Commits](https://www.conventionalcommits.org) style. The full checklist lives in [CONTRIBUTING.md](https://github.com/SakshhamTheCoder/adbt/blob/main/CONTRIBUTING.md).

## Getting help

- **Bugs:** open an issue with your adbt version (`adbt --version`), operating system, phone model and Android version, and the steps to reproduce.
- **Security problems:** do not post them publicly — follow `SECURITY.md`.