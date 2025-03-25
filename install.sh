#!/bin/bash

set -e  # Exit on any error

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
WHITE='\033[1;37m'
NC='\033[0m' # No Color

BIN_NAME="kpop"
README_FILE="README.md"
DEFAULT_INSTALL_DIR="$HOME/.local/bin"
README_INSTALL_DIR="$HOME/.local/share/kpop-cli"
SYSTEM_INSTALL_DIR="/usr/local/bin"
SYSTEM_README_INSTALL_DIR="/usr/local/share/kpop-cli"

# Detect OS & Architecture
ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

# Normalize architecture names for GoReleaser
if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "aarch64" ]]; then
    ARCH="arm64"
fi

# Set up Windows environment
if [[ "$OS" == "darwin" || "$OS" == "linux" ]]; then
    TARBALL="${BIN_NAME}_${OS}_${ARCH}.tar.gz"
    INSTALL_METHOD="tarball"
else
    # Windows: We'll need a .zip for Windows
    OS="windows"
    TARBALL="${BIN_NAME}_${OS}_${ARCH}.zip"
    INSTALL_METHOD="zip"
fi

# Fetch latest version by following GitHub redirect
LATEST_VERSION=$(curl -Ls -o /dev/null -w %{url_effective} https://github.com/DDaaaaann/kpop-cli/releases/latest | grep -oE "[^/]+$")

if [[ -z "$LATEST_VERSION" ]]; then
    echo -e "${RED}Failed to fetch latest version. Exiting.${NC}"
    exit 1
fi

# Choose installation location
echo -e "${WHITE}Choose installation method:${NC}"
echo -e "1) ${GREEN}Install to user directory ($DEFAULT_INSTALL_DIR) [Recommended]${NC}"
echo -e "2) ${YELLOW}Install system-wide ($SYSTEM_INSTALL_DIR) [Requires sudo]${NC}"

# Check if we are in interactive mode (for piping)
if [ -t 0 ]; then
    read -p "Enter option (1 or 2): " OPTION
else
    # If not interactive, continue to ask for input
    while true; do
        echo -e "${RED}You must provide a valid option (1 or 2). Please try again.${NC}"
        read -p "Enter option (1 or 2): " OPTION
        if [[ "$OPTION" == "1" || "$OPTION" == "2" ]]; then
            break
        fi
    done
fi

if [[ "$OPTION" == "2" ]]; then
    INSTALL_DIR="$SYSTEM_INSTALL_DIR"
    README_INSTALL_DIR="$SYSTEM_README_INSTALL_DIR"
    SUDO="sudo"
else
    INSTALL_DIR="$DEFAULT_INSTALL_DIR"
    SUDO=""
fi

# Ensure install directory exists
$SUDO mkdir -p "$INSTALL_DIR"
$SUDO mkdir -p "$README_INSTALL_DIR"

# Download the appropriate tarball or zip based on OS
echo -e "${YELLOW}⚡ Downloading $TARBALL...${NC}"
curl -L -o "/tmp/$TARBALL" "https://github.com/DDaaaaann/kpop-cli/releases/download/$LATEST_VERSION/$TARBALL"

# Extract or unzip the file based on the OS
if [[ "$INSTALL_METHOD" == "tarball" ]]; then
    echo -e "${YELLOW}⚡ Extracting $BIN_NAME and $README_FILE...${NC}"
    tar -xzf "/tmp/$TARBALL" -C "/tmp/"
    # Move the binary and README file to the appropriate location
    $SUDO mv "/tmp/$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"
    $SUDO chmod +x "$INSTALL_DIR/$BIN_NAME"
    $SUDO mv "/tmp/$README_FILE" "$README_INSTALL_DIR/$README_FILE"
elif [[ "$INSTALL_METHOD" == "zip" ]]; then
    echo -e "${YELLOW}⚡ Unzipping $BIN_NAME and $README_FILE...${NC}"
    unzip -q "/tmp/$TARBALL" -d "/tmp/"
    # Move the binary and README file to the appropriate location
    $SUDO mv "/tmp/$BIN_NAME.exe" "$INSTALL_DIR/$BIN_NAME.exe"
    $SUDO mv "/tmp/$README_FILE" "$README_INSTALL_DIR/$README_FILE"
    # Windows-specific: No need for chmod on .exe files
fi

# Clean up the temporary downloaded file
rm "/tmp/$TARBALL"

# Add to PATH if needed
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${YELLOW}️  $INSTALL_DIR is not in your PATH.${NC}"
    echo -e "${WHITE}Add this to your shell profile:${NC}"
    echo -e "  echo 'export PATH=\"$INSTALL_DIR:\$PATH\"' >> ~/.bashrc"
    echo -e "  source ~/.bashrc"
fi

echo -e "${WHITE}The README has been installed to: $README_INSTALL_DIR/$README_FILE${NC}"
echo -e "${GREEN}Installation complete!${NC}"
echo -e "${GREEN}Run '${BIN_NAME} --help' to start.${NC}"
