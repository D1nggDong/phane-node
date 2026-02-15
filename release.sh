#!/bin/bash
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

echo "🛡️ Fetching PHANE Sovereign Node for $OS/$ARCH..."

# Determine binary name based on architecture
if [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
    BINARY="phane_arm64"
else
    BINARY="phane_amd64"
fi

# In production, this would wget from your GitHub releases
# For now, we build it locally to test the "One-File" portability
go build -o $BINARY main.go

echo "✅ Success. Run ./$BINARY to launch your agent."
