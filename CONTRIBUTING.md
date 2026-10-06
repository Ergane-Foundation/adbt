# Contributing to adbt

Thanks for helping out. Bug reports, fixes, and small features are all welcome.

## Before you start

- For bugs, open an issue with your adbt version (`adbt --version`), OS, device, and steps to reproduce.
- For new features, open an issue first so we can agree on the approach before you write code.
- By participating you agree to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## Development setup

You need Go (see the version in `go.mod`) and `adb` on your `PATH`.

```bash
go build ./cmd/adbt
./adbt
```

Before opening a pull request, run:

```bash
gofmt -l .
go vet ./...
go test ./...
```

The documentation site lives in `website/` and uses Docusaurus:

```bash
cd website
npm install
npm start
```

## Pull requests

- Branch from `main` and keep each pull request focused on one change.
- Write short, lowercase commit messages that say what the change does, for example `add connect and disconnect to the devices screen`.
- Update the README or the docs in `website/docs` when behavior changes.

## License

By contributing, you agree that your contributions are licensed under the [Apache License 2.0](LICENSE).
