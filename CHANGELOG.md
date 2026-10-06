# Changelog

All notable changes to this project are documented here. The format is based on
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and the project uses
[Semantic Versioning](https://semver.org/spec/v2.0.0.html). Before 1.0, minor
versions may contain breaking changes.

## [Unreleased]

## [0.2.1] - 2026-10-06

### Breaking

- License changed from MIT to Apache-2.0. Releases up to 0.2.0 stay under MIT.

### Added

- Code of conduct, contributing guide, security policy, support guide,
  maintainers list and third-party notices.

### Changed

- Release archives, deb and rpm packages, and the AUR package now include
  `NOTICE` and `THIRD_PARTY_NOTICES.md` next to `LICENSE`.
- Commit messages follow Conventional Commits.

## [0.2.0] - 2026-07-10

### Added

- `--version` flag, and a startup check that exits with install instructions
  when `adb` is missing.
- App Manager: details pane (version, size, target SDK, APK path), multi-select
  for batch uninstall, and APK extract.
- Device Info: screenshots and screen recording saved to the computer.
- Logcat: save to a file, and filter by package or PID.
- Input screen for sending text and key events to the device.
- File Explorer: create directories.
- Performance Monitor: battery level, temperature, voltage and health.
- Timeout on `adb` commands.

### Changed

- Updated to Go 1.26 and newer dependencies.
- Polished layouts and forms, and added autocomplete to the Intent Tester.

### Fixed

- Typing `q` in a form no longer quits the app.
- An error is shown when device stats fail to load.

## [0.1.0] - 2026-02-19

First version.
