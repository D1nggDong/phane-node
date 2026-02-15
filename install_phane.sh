#!/bin/bash
echo "🐙 INITIALIZING PHANE NODE AGENT..."
# 1. Install dependencies
sudo apt-get update && sudo apt-get install -y curl cloudflared
# 2. Ask for the Pairing Token (provided by your website)
read -p "ENTER YOUR 6-DIGIT PAIRING TOKEN: " TOKEN
# 3. Create the local config linking this device to the Phane Hub
echo "pairing_token: $TOKEN" > ~/.phane_config.yml
# 4. Start the tunnel and the brain
echo "✅ DEVICE LINKED SUCCESSFULLY. REBOOTING OCTOPUS AGENT..."
