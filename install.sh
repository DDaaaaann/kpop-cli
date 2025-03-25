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

# Fetch latest version by following GitHub redirect
LATEST_VERSION=$(curl -Ls -o /dev/null -w %{url_effective} https://github.com/DDaaaaann/kpop-cli/releases/latest | grep -oE "[^/]+$")

if [[ -z "$LATEST_VERSION" ]]; then
    echo -e "${RED}Failed to fetch latest version. Exiting.${NC}"
    exit 1
fi

# Construct tarball name based on GoReleaser naming convention
TARBALL="${BIN_NAME}_${OS}_${ARCH}.tar.gz"

# Choose installation location
echo -e "${WHITE}Choose installation method:${NC}"
echo -e "1) ${GREEN}Install to user directory ($DEFAULT_INSTALL_DIR) [Recommended]${NC}"
echo -e "2) ${YELLOW}Install system-wide ($SYSTEM_INSTALL_DIR) [Requires sudo]${NC}"
read -p "Enter option (1 or 2): " OPTION

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

# Download the tar.gz file
echo -e "${YELLOW}⚡ Downloading $TARBALL...${NC}"
curl -L -o "/tmp/$TARBALL" "https://github.com/DDaaaaann/kpop-cli/releases/download/$LATEST_VERSION/$TARBALL"

# Extract the binary and README
echo -e "${YELLOW}⚡ Extracting $BIN_NAME and $README_FILE...${NC}"
tar -xzf "/tmp/$TARBALL" -C "/tmp/"

# Move the binary to the install directory
$SUDO mv "/tmp/$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"
$SUDO chmod +x "$INSTALL_DIR/$BIN_NAME"

# Move the README file to the appropriate directory
$SUDO mv "/tmp/$README_FILE" "$README_INSTALL_DIR/$README_FILE"

# Clean up
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
