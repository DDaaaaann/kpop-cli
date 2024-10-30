# Kpop CLI Tool

A command-line tool for killing processes running on a specified port. This tool is useful for developers who encounter errors stating that a port is already in use.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [License](#license)
- [Contributing](#contributing)
- [Acknowledgements](#acknowledgements)

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

If you run into an "address already in use" error while starting a server, use the autodetect feature:

```bash
$ kpop
Kill process using port 8080 (PID 12345)? (y/n)
Killed process 12345 on port 8080.
```


## License
This project is licensed under the GNU General Public License (GPL) v3. See the LICENSE file for details.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request if you have suggestions or improvements.