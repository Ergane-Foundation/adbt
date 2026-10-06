<p align="center">
  <img src="website/static/img/logoadbt.png" alt="adbt Logo" width="120" />
</p>

<h1 align="center">adbt</h1>

<p align="center">
  <b>A modern, keyboard-driven Terminal User Interface (TUI) for Android Debug Bridge.</b>
</p>

<p align="center">
  <a href="https://adbt-tui.vercel.app"><b>Documentation Website »</b></a>
</p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-blue?style=flat-square" alt="License: Apache-2.0" /></a>
  <img src="https://img.shields.io/github/v/release/SakshhamTheCoder/adbt?style=flat-square&color=purple" alt="Latest Release" />
  <img src="https://img.shields.io/badge/built%20with-Bubble%20Tea-brightgreen?style=flat-square" alt="Built with Bubble Tea" />
</p>

---

**adbt** is a fast, keyboard-driven terminal user interface for interacting with Android devices via ADB. Built in Go using Charm’s **Bubble Tea** framework, it provides a structured and interactive command-line workspace that helps you avoid complex raw shell commands and context-switching.

For detailed guides, shortcuts, and troubleshooting, visit the [adbt Documentation Website](https://adbt-tui.vercel.app).

## Screenshots

<p align="center">
  <img src="website/static/img/screenshots/dashboard.png" alt="adbt dashboard listing every screen, with the connected device shown in the header" width="100%" />
  <br />
  <b>Dashboard</b>
</p>

<table>
  <tr>
    <td width="50%" align="center"><img src="website/static/img/screenshots/device_info.png" alt="Device Info screen" width="100%" /><br /><b>Device Info</b></td>
    <td width="50%" align="center"><img src="website/static/img/screenshots/app_manager.png" alt="App Manager screen" width="100%" /><br /><b>App Manager</b></td>
  </tr>
  <tr>
    <td width="50%" align="center"><img src="website/static/img/screenshots/file_explorer.png" alt="File Explorer screen" width="100%" /><br /><b>File Explorer</b></td>
    <td width="50%" align="center"><img src="website/static/img/screenshots/logcat.png" alt="Logcat Viewer screen" width="100%" /><br /><b>Logcat Viewer</b></td>
  </tr>
  <tr>
    <td width="50%" align="center"><img src="website/static/img/screenshots/performance.png" alt="Performance Monitor screen" width="100%" /><br /><b>Performance Monitor</b></td>
    <td width="50%" align="center"><img src="website/static/img/screenshots/input.png" alt="Input Sender screen" width="100%" /><br /><b>Input Sender</b></td>
  </tr>
  <tr>
    <td width="50%" align="center"><img src="website/static/img/screenshots/intent_tester.png" alt="Intent Tester screen" width="100%" /><br /><b>Intent Tester</b></td>
    <td width="50%" align="center"><img src="website/static/img/screenshots/port_manager.png" alt="Port Forwarding screen" width="100%" /><br /><b>Port Forwarding</b></td>
  </tr>
</table>

## Features

- **Device Management**: Real-time connected device detection, wireless pairing over IP/port/PIN, and automatic selection when a single device is attached.
- **Device Info**: Model, serial, Android version, battery, storage, screen size/density, and IP address, plus Wi-Fi and screen toggles and reboot controls (device, recovery, bootloader).
- **Scrcpy Integration**: Launch high-performance screen mirroring with a single keypress.
- **Screen Capture**: Take screenshots and record the screen straight from Device Info, saved to your host machine.
- **Performance Monitor**: Real-time CPU, memory, and network usage, plus live battery level, temperature, voltage, and health readouts.
- **App Manager**: Browse, search, and filter user/system apps; inspect per-app details (version, size, target SDK, APK path); multi-select for batch uninstall; install or extract APKs; and launch, force-stop, or clear data.
- **File Explorer**: Browse the device filesystem, push and pull files, create directories, and delete with confirmation.
- **Logcat Viewer**: Live stream device logs with severity filters, text search highlighting, filtering by package/PID, pause/resume, clear, and one-key save to a file.
- **Input Sender**: Type text and send key events (Back, Home, Recents, arrows, and more) to the device, with an optional live keystroke-forwarding mode.
- **Intent Tester**: Construct and send custom activity/broadcast intents to test deep links and receiver behavior.
- **Port Forwarding**: Easily configure forward and reverse network connections between your host and device.

---

## Installation

### Homebrew (macOS / Linux)
```bash
brew install --cask SakshhamTheCoder/tap/adbt
# or
brew tap SakshhamTheCoder/tap && brew install --cask adbt
```

### Scoop (Windows)
```bash
scoop bucket add SakshhamTheCoder https://github.com/SakshhamTheCoder/scoop-bucket
scoop install adbt
```

### AUR (Arch Linux)
```bash
yay -S adbt-bin
# or
paru -S adbt-bin
```

For manual binary downloads (`.deb`, `.rpm`, or generic archives), check the [Releases Page](https://github.com/SakshhamTheCoder/adbt/releases) or the [Installation Docs](https://adbt-tui.vercel.app/docs/installation).

---

## Usage

Once installed, plug in your device and run:
```bash
adbt
```

Check which build you're on with `adbt --version` (or `adbt -v`). `adbt` requires `adb` on your `PATH` and exits with install instructions if it is missing.

| Key | Action |
| :--- | :--- |
| `q` / `Ctrl+C` | Quit `adbt` |
| `Esc` / `Backspace` | Back / Go Up a directory |
| `↑` `↓` / `k` `j` | Navigate lists and menus |
| `Enter` | Select / Confirm / Launch |

For the full set of module-specific hotkeys (App Manager, File Explorer, Logcat, etc.), refer to the [Keyboard Shortcuts Cheat Sheet](https://adbt-tui.vercel.app/docs/shortcuts).

---

## Troubleshooting and support

If you run into connection issues, unauthorized device screens, or a missing `adb`:
- Read the [troubleshooting guide](https://adbt-tui.vercel.app/docs/troubleshooting).
- See [SUPPORT.md](SUPPORT.md) for where to ask questions and report bugs.
- Report security problems privately, as described in [SECURITY.md](SECURITY.md).

## Contributing

Contributions are welcome, from documentation fixes to new screens. Start with
[CONTRIBUTING.md](CONTRIBUTING.md) and look for issues labelled
`good first issue`. Please follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## License

[Apache License 2.0](LICENSE). See [NOTICE](NOTICE) and
[THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md) for third-party components.

Created by [Sakshham Bhagat](https://github.com/SakshhamTheCoder).
