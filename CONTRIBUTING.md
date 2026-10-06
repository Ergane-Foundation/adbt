# Contributing

Thank you for helping. This guide covers setup on Windows, macOS and Linux, the
checks your change must pass, and how pull requests are reviewed. By taking part
you agree to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

Security problems are not reported here: see [SECURITY.md](SECURITY.md).

## What you need

| Tool | Version | Notes |
|------|---------|-------|
| Go | see `go.mod` | from go.dev/dl |
| adb | recent | Android SDK Platform Tools, on your `PATH` |
| scrcpy | recent | optional, only for screen mirroring |
| Node.js | 20 or newer | only for the documentation website in `website/` |

You also need an Android phone or emulator with USB debugging enabled.

## Setup

```sh
git clone https://github.com/<your-username>/adbt.git
cd adbt
go build ./cmd/adbt
./adbt
```

On Windows run `.\adbt.exe` instead of `./adbt`.

## Running the documentation website

```sh
cd website
npm install
npm start
```

## Checks

Run these before opening a pull request:

```sh
gofmt -l .        # must print nothing
go vet ./...
go test ./...
```

For changes to `.goreleaser.yaml`, also run `goreleaser check`.

## Commit messages

One short line in lower case, in the imperative mood, with no full stop, saying
what the change does. See `git log` for the style:

- `add connect and disconnect to the devices screen`
- `keep typing q in forms from quitting the app`

## Branches and pull requests

1. Comment on the issue you want to work on and wait to be assigned, so two
   people do not do the same work.
2. Fork the repository and create a branch from `main` with a short descriptive
   name, for example `fix-logcat-pause`.
3. Keep the pull request focused on one issue. Link it in the description
   (`Closes #123`).
4. Say in the description how you tested the change: which screens, which
   operating system, which phone or emulator.
5. Make sure all checks pass. A maintainer will review; please answer comments
   by pushing new commits rather than force-pushing, so the review is easy to follow.

## What makes a good pull request

- It solves the linked issue and nothing else. No unrelated formatting changes.
- It includes tests for new behavior, or explains why a test is not practical.
- It updates the README or `website/docs` when behavior or shortcuts change.
- New dependencies are necessary, actively maintained and have a license
  compatible with Apache-2.0. Mention them in the description and add them to
  [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
- Anything passed to `adb shell` is quoted, so user input cannot run extra commands.
- Documentation is in plain English, ASCII punctuation, and makes no claim that
  is not measured or cited.

## Open-source event rules

During community events (for example Hacktoberfest), we value quality over quantity.

- Only pull requests that address an open issue, or a fix that clearly deserves
  one, are reviewed.
- Pull requests that only change whitespace, reword text without improving it,
  add your name somewhere, or are generated without understanding the code, are
  closed and labelled `invalid` or `spam`.
- Accepted pull requests are labelled `hacktoberfest-accepted`.
- Be patient: maintainers are volunteers. Do not ping repeatedly.

## Getting help

See [SUPPORT.md](SUPPORT.md).

## License

By contributing, you agree that your contributions are licensed under the
[Apache License 2.0](LICENSE), as stated in its section 5.
