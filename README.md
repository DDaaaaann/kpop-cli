<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/DDaaaaann/kpop-cli">
    <img src="assets/logo.jpg" alt="Logo" width="400" height="400">
  </a>

<h3 align="center">KPOP-CLI</h3>
  <em>Kill Process on Port</em>
  <p align="center">
    A CLI solution to kill any process on any port!
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template"><strong>Installation »</strong></a>
    <br />
    <br />
    <a href="#usage">Usage</a>
    ·
    <a href="https://github.com/DDaaaaann/kpop-cli/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/DDaaaaann/kpop-cli/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
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

### Using APT (Debian/Ubuntu)

```bash
sudo apt install kpop-cli
```

### Using Homebrew (macOS)

```bash
brew install DDaaaaann/kpop-cli/kpop-cli
```

### Using Winget (Windows)

```bash
winget install DDaaaaann.kpop-cli
```

### From source

To install the tool from source, clone the repository and build it:

```bash
git clone https://github.com/DDaaaaann/kpop-cli.git
cd kpop-cli
go build
```

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

If you run into an "address already in use" error while starting a server, use the autodetect
feature:

```bash
$ kpop
Kill process using port 8080 (PID 12345)? (y/n)
Killed process 12345 on port 8080.
```

## License

This project is licensed under the GNU General Public License (GPL) v3. See the LICENSE file for
details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have suggestions or
improvements.