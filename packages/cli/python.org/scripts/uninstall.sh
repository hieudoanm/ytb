#!/usr/bin/env bash
set -e

BINARY_NAME="ytts"

TARGET_DIR="$HOME/.local/bin"
if [ "$1" = "--global" ]; then
  TARGET_DIR="/usr/local/bin"
  echo "Uninstalling $BINARY_NAME globally from $TARGET_DIR..."
  sudo rm -f "$TARGET_DIR/$BINARY_NAME"
else
  echo "Uninstalling $BINARY_NAME locally from $TARGET_DIR..."
  rm -f "$TARGET_DIR/$BINARY_NAME"
fi

echo "Uninstallation complete: $BINARY_NAME"
