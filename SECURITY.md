# Security Policy

## Supported versions

Only the latest release of adbt receives security fixes. Please update before reporting.

## Reporting a vulnerability

Please do not open a public issue for security problems.

Report privately through the **Security** tab of this repository by choosing **Report a vulnerability**, or email sakshhamtg@gmail.com.

Include the adbt version (`adbt --version`), your operating system, and steps to reproduce. You can expect a reply within a week. Once a fix is released, the report is credited in the release notes unless you ask otherwise.

adbt runs `adb` commands against devices you connect, and `adbt serve` opens a local web UI. Issues that let another machine, page, or app reach a device through adbt without the user's consent are in scope.
