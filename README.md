[//]: # ( TODO: Add badges)
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/DDaaaaann/kpop-cli">
    <img src="https://github.com/user-attachments/assets/64e25e73-8450-4032-967d-06402a7b76b9" alt="Logo" width="400" height="400">
  </a>

<h3 align="center">KPOP-CLI</h3>
<em>Kill Process on Port</em>
  <p align="center">
    A CLI solution to kill any process on any port!
    <br />
    <a href="#installation"><strong>Installation »</strong></a>
    <br />
    <br />
    <a href="#usage">Usage</a>
    ·
    <a href="https://github.com/DDaaaaann/kpop-cli/issues/new?labels=Bug">Report Bug</a>
    ·
    <a href="https://github.com/DDaaaaann/kpop-cli/issues/new?labels=Improvement">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#example">Example</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contributing">Contributing</a></li>
  </ol>
</details>

## About The Project

A command-line tool for killing processes running on a specified port. This tool is useful for
developers who encounter errors stating that a port is already in use.

## Features

- Kill a process running on a specified port.
- Autodetect the port in use from error messages in the terminal.
- Cross-platform compatibility (Unix and Windows).
- Includes help command for usage instructions.
- Man page support for easy reference.
- Autocompletion for command-line inputs.
- Search command line history for previously used ports (nice to have).

## Installation

You can install `kpop` in different ways, depending on your OS and preference.

### Package Managers (Coming Soon)
Support for package managers is coming soon. You will be able to install `kpop` using:
- **Homebrew (macOS & Linux)**: `brew install DDaaaaann/kpop-cli/kpop`
- **APT (Debian/Ubuntu)**: `sudo apt install kpop`
- **Winget (Windows)**: `winget install kpop`


### Quick Install Using Shell Script
Run the following command to install `kpop` automatically:

```sh
curl -fsSL https://raw.githubusercontent.com/DDaaaaann/kpop-cli/main/install.sh | bash
```

This script will:
- Detect your OS and architecture.
- Download the latest version from GitHub.
- Install `kpop` in a local user directory (`~/.local/bin` on Linux/macOS, `~/.kpop/bin` on Windows).
- Add it to your `PATH` (or guide you on how to do so).

Once installed, run:

```sh
kpop --help
```

### Manual Installation
If you prefer to install manually:

1. Download the latest release from [GitHub Releases](https://github.com/DDaaaaann/kpop-cli/releases/latest).
2. Extract the binary:
   ```sh
   tar -xzf kpop_linux_amd64.tar.gz  # Replace with the correct OS/architecture
   ```
3. Move it to a directory in your `PATH`, e.g.:
   ```sh
   mv kpop ~/.local/bin/
   chmod +x ~/.local/bin/kpop
   ```

For Windows, download the `.zip`, extract `kpop.exe`, and place it in a folder like `C:\Users\YourName\.kpop\bin`.


## Updating `kpop`
To update to the latest version, simply re-run the installation script:

```sh
curl -fsSL https://raw.githubusercontent.com/DDaaaaann/kpop-cli/main/install.sh | bash
```

Or, if installed via Homebrew (once available):

```sh
brew upgrade kpop
```


### Uninstallation
To remove `kpop`:

```sh
rm -rf ~/.local/bin/kpop ~/.local/share/kpop-cli  # Linux/macOS
rm -rf ~/.kpop/bin/kpop.exe ~/.kpop/bin/share/kpop-cli  # Windows
```

If you manually added `kpop` to `PATH`, remove the entry from your shell profile.

## Usage

To kill a process on a specified port, run the following command:

```bash
kpop [-f] [-q] <port>
```

Replace <port> with the port number you want to free.

### Autodetecting Port

To automatically detect and kill the process on the last used port, just run:

```bash
kpop
```

The autodetect functionality scans previous temrinal output.

### Help command

To see usage instructions, run:

```bash
kpop -h or kpop --help
```

## Example

To kill a process running on port 8080:

```bash
$ kpop 8080
Kill process using port 8080 (PID 12345)? (y/n)
Killed process 12345 on port 8080.
```

## License

This project is licensed under the GNU General Public License (GPL) v3. See the LICENSE file for
details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have suggestions or
improvements.
