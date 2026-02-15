# PHANE Sovereign OS 🛰️
> **Official Site:** [askphane.com](https://askphane.com) | **Status:** v1.0.0-beta

PHANE is a lightweight, secure, and hardware-agnostic agentic hub. Designed to run blazingly fast on **Raspberry Pi 4B (1GB)**, it provides a unified interface for local AI and cloud-based productivity tools.

[![Build Status](https://github.com/D1nggDong/phane-node/actions/workflows/build.yml/badge.svg)](https://github.com/D1nggDong/phane-node/actions)

## ✨ Sovereign Features
* **Zero-Cloud Auth:** Secure local login and registration. Your credentials never leave your hardware.
* **Ollama Integration:** Plug into local LLMs for private, uncensored intelligence.
* **AES-256 Vault:** Industry-standard encryption for sensitive keys (Notion, etc.).
* **Universal Binary:** Built in Go for effortless execution on Mac (M1/Intel), Windows, and Linux.

## 🚀 One-Line Installation
Deploy your node instantly with:

\`\`\`bash
curl -sSL https://askphane.com/install.sh | bash
\`\`\`

## 🏗️ Architecture
PHANE operates on a "Local-First" principle. While the landing page at **askphane.com** provides the gateway, the actual registration and login logic are handled by your individual node.



## 🔧 Active Roadmap
- [x] **Core OS:** Lightweight Go-based kernel.
- [x] **Cross-Platform:** GitHub Actions automated builds.
- [ ] **Notion Bridge:** Encrypted sync for long-term memory.
- [ ] **Web Dashboard:** Local UI for node management.
- [ ] **Subscription Model:** Tiered feature access for advanced users.

## 📄 License
Distributed under the MIT License. See `LICENSE` for more information.
