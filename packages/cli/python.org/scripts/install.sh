#!/usr/bin/env bash
set -e

# -------- CONFIGURE --------
RAW_URL="https://raw.githubusercontent.com/hieudoanm/micro/packages/cli/python.org/ytts/bin/ytts" # <-- Direct raw link to binary
BINARY_NAME="ytts"
# ----------------------------

# Default install dir (local)
TARGET_DIR="$HOME/.local/bin"

# Check for --global flag
if [ "$1" = "--global" ]; then
  TARGET_DIR="/usr/local/bin"
  echo "Installing $BINARY_NAME globally to $TARGET_DIR..."
  sudo curl -L "$RAW_URL" -o "$TARGET_DIR/$BINARY_NAME"
  sudo chmod +x "$TARGET_DIR/$BINARY_NAME"
else
  echo "Installing $BINARY_NAME locally to $TARGET_DIR..."
  mkdir -p "$TARGET_DIR"
  curl -L "$RAW_URL" -o "$TARGET_DIR/$BINARY_NAME"
  chmod +x "$TARGET_DIR/$BINARY_NAME"
  echo "Ensure $TARGET_DIR is in your PATH:"
  echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
fi

echo "Installation complete: $BINARY_NAME"
