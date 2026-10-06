# Security policy

adbt runs `adb` commands against your Android devices and moves files between
them and your computer, so security reports are taken seriously and handled
privately.

## Reporting a vulnerability

**Do not open a public issue, discussion or pull request for a security problem.**

Report it privately through GitHub: open the repository's **Security** tab and
click **Report a vulnerability** (GitHub private vulnerability reporting). Only
the maintainers can see the report.

Please include:

- what is affected (screen, file, function, operating system, version or commit);
- steps to reproduce, or a proof of concept;
- the impact you expect, and any idea for a fix.

What to expect:

- an acknowledgement within 7 days;
- an assessment and a plan within 30 days;
- credit in the advisory and the changelog when the fix is released, unless you
  prefer to stay anonymous.

This is a volunteer project, so these are goals, not guarantees.

## In scope

- Command injection: text entered in adbt (file names, package names, intent
  extras, input text) running as a different command on the device or computer.
- File transfers writing outside the folder you chose, on the device or the computer.
- Release archives or packages (Homebrew, Scoop, AUR, deb, rpm) that do not match
  the published source or checksums.

## Out of scope

- Bugs in `adb`, `scrcpy` or Android itself; report those upstream.
- Attacks that need an already authorized USB debugging connection to your
  device or a shell on your computer. adbt has the same access as `adb`.
- Vulnerabilities in dependencies that do not affect adbt; report those to the
  dependency.

## Supported versions

| Version | Supported |
|---------|-----------|
| 0.2.x | Yes |
| 0.1.x | No |

Before 1.0, only the latest minor version gets security fixes.
