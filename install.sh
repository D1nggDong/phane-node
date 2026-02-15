#!/bin/bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

echo "🔍 Detecting System: $OS / $ARCH"

if [ "$OS" == "darwin" ]; then
    if [ "$ARCH" == "arm64" ]; then
        FILE="phane-mac-m1"
    else
        FILE="phane-mac-intel"
    fi
elif [ "$OS" == "linux" ]; then
    FILE="phane-linux-x64"
else
    FILE="phane-windows.exe"
fi

echo "🚀 Downloading $FILE for your system..."
# In production, this would be: curl -L https://phane.ai/bin/$FILE -o phane
# For now, we'll just simulate the link:
echo "✅ Ready to launch: ./$FILE"
