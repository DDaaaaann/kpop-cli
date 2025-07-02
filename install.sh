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
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Normalize architecture names for GoReleaser
if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "aarch64" ]]; then
    ARCH="arm64"
fi

# Windows handling
if [[ "$OS" == "mingw"* || "$OS" == "cygwin" || "$OS" == "msys" ]]; then
    OS="windows"
    EXT=".exe"
    TARBALL="${BIN_NAME}_${OS}_${ARCH}.zip"
    INSTALL_DIR=$(powershell.exe -Command "echo /$env:USERPROFILE/.kpop/bin" | tr -d '\r')
else
    EXT=""
    TARBALL="${BIN_NAME}_${OS}_${ARCH}.tar.gz"
    INSTALL_DIR="$HOME/.local/bin"
fi

README_INSTALL_DIR="$INSTALL_DIR/share/kpop-cli"

# Fetch latest version
LATEST_VERSION=$(curl -Ls -o /dev/null -w %{url_effective} https://github.com/DDaaaaann/kpop-cli/releases/latest | grep -oE "[^/]+$")

if [[ -z "$LATEST_VERSION" ]]; then
    echo -e "${RED}Failed to fetch latest version. Exiting.${NC}"
    exit 1
fi

# Ensure install directories exist
mkdir -p "$INSTALL_DIR"
mkdir -p "$README_INSTALL_DIR"

# Download the package
echo -e "${YELLOW}⚡ Downloading $TARBALL...${NC}"
curl -L -o "/tmp/$TARBALL" "https://github.com/DDaaaaann/kpop-cli/releases/download/$LATEST_VERSION/$TARBALL"

# Extract and install
echo -e "${YELLOW}⚡ Extracting...${NC}"
if [[ "$OS" == "windows" ]]; then
    unzip -o "/tmp/$TARBALL" -d "/tmp/"
    mv "/tmp/$BIN_NAME$EXT" "$INSTALL_DIR/$BIN_NAME$EXT"
else
    tar -xzf "/tmp/$TARBALL" -C "/tmp/"
    mv "/tmp/$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"
    chmod +x "$INSTALL_DIR/$BIN_NAME"
fi

# Move README
mv "/tmp/$README_FILE" "$README_INSTALL_DIR/$README_FILE"

# Clean up
rm "/tmp/$TARBALL"

# Add to PATH
if [[ "$OS" == "windows" ]]; then
    if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
        WINDOWS_PATH=$(echo "$INSTALL_DIR" | sed 's/\//\\/g')
        echo -e "${YELLOW}⚡ Adding $WINDOWS_PATH to your PATH...${NC}"
        powershell.exe -Command "[System.Environment]::SetEnvironmentVariable('Path', \$env:Path + ';$WINDOWS_PATH', [System.EnvironmentVariableTarget]::User)"
        echo -e "${GREEN}Please restart your terminal or run 'refreshenv' for changes to take effect.${NC}"
    fi
else
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo -e "${YELLOW}️  $INSTALL_DIR is not in your PATH.${NC}"
        echo -e "${WHITE}Add this to your shell profile:${NC}"
        echo -e "  echo 'export PATH=\"$INSTALL_DIR:\$PATH\"' >> ~/.bashrc"
        echo -e "  source ~/.bashrc"
    fi
fi

echo -e "${WHITE}Kpop has been installed to: $INSTALL_DIR/$BIN_NAME$EXT${NC}"
echo -e "${WHITE}The README has been installed to: $README_INSTALL_DIR/$README_FILE${NC}"
echo -e "${GREEN}Installation complete!${NC}"
echo -e "${GREEN}Run '${BIN_NAME} --help' to start.${NC}"
