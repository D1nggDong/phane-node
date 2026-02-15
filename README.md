# PHANE Sovereign OS 🛰️
> **Status:** Public Beta | **Website:** [phane.ai](https://D1nggDong.github.io/phane-node/)

PHANE is a lightweight, secure, and hardware-agnostic agentic hub designed to run on anything from a **Raspberry Pi 4B (1GB)** to a high-end Mac or Windows workstation.

[![Build Status](https://github.com/D1nggDong/phane-node/actions/workflows/build.yml/badge.svg)](https://github.com/D1nggDong/phane-node/actions)

## ✨ Key Features
* **Sovereign AI:** Connect to local Ollama instances. Your data, your hardware.
* **1GB RAM Optimized:** Built in Go to ensure even the smallest nodes stay fast.
* **Encrypted Vault:** Sensitive API keys (Notion, etc.) are encrypted at rest with AES-256 GCM.
* **Cross-Platform:** Native binaries for Linux (ARM/x64), macOS (Intel/M1), and Windows.

## 🚀 Quick Install (Any System)
Run the following command in your terminal to download the latest version for your specific hardware:

\`\`\`bash
curl -sSL https://raw.githubusercontent.com/D1nggDong/phane-node/main/install.sh | bash
\`\`\`

## 🔧 Skills & Integrations
### 📝 Notion (Active Development)
We are currently building the Notion bridge. This will allow PHANE to:
* Archive conversations directly to your Notion workspace.
* Use Notion pages as a "Long-term Memory" for your local LLM.
* Securely store your Notion Token in the PHANE Vault.

## 📄 License
Distributed under the MIT License. See \`LICENSE\` for more information.
