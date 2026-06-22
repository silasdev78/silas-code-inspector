content = r"""# 🔍 Silas Code Inspector

[![Go Version](https://img.shields.io/github/go-mod/go-version/silasdev78/silas-code-inspector)](https://go.dev)
[![Build](https://img.shields.io/github/actions/workflow/status/silasdev78/silas-code-inspector/silas.yml?branch=main)](https://github.com/silasdev78/silas-code-inspector/actions)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/silasdev78/silas-code-inspector)](https://github.com/silasdev78/silas-code-inspector/releases)

**Silas Code Inspector** is a multi-language static security scanner designed to find vulnerabilities in **TON blockchain smart contracts (Tact & FunC)**, as well as Go, Docker, and Web codebases. Over **90 security patterns**, adaptive learning, JSON/SARIF export, and ready for CI/CD.

## ✨ Features

- ⚡ **90+ patterns for TON** – 45 for Tact, 45 for FunC (replay, race condition, cell depth, impure, etc.)
- 🧠 **Adaptive learning** – pattern weights adjust via feedback (`.silas-state.json`)
- 📦 **Go dependency CVE scanner** – Check your `go.mod` against known CVEs
- 🐳 **Docker & Web scanners** – hardcoded secrets, XSS, CSRF, root user, CVE check
- 🌈 **Colored terminal output** + JSON / SARIF export (GitHub Code Scanning)
- 🧹 **Comment filtering** and **only meaningful `receive` analysed** – low noise
- 📊 **Summary mode** (`--summary`) for quick overview
- 🔌 **CLI-first**, ready for CI/CD and future APIs
- ⚙️ **GitHub Actions** – automatically scans your code on every push and PR

---

## 📋 Supported Languages & Extensions

| Language / Target | Extensions | Patterns | Notes |
|-------------------|------------|----------|-------|
| **Tact (TON)** | `.tact` | 45 | Smart contract security (Jetton, NFT, DeFi) |
| **FunC (TON)** | `.fc`, `.fif` | 45 | Legacy TON contracts |
| **Go** | `.go` | 8 | Hardcoded secrets, TLS, SQL injection, crypto |
| **Go modules** | `go.mod` | CVE feed | Checks dependencies against known vulnerabilities |
| **Docker** | `Dockerfile` | 4 | Root user, HEALTHCHECK, latest tag |
| **Web** | `.html`, `.js`, `.ts` | 4 | XSS, CSRF, mixed content, innerHTML |

---

## 🚀 Quick Start

### Prerequisites

- Go 1.22+ installed
- Git

### Install from source

    go install github.com/silasdev78/silas-code-inspector/cmd/silas@latest

### Or build locally

    git clone https://github.com/silasdev78/silas-code-inspector.git
    cd silas-code-inspector
    make build
    ./silas --help

---

## 📖 Usage

### Scan a single file

    silas contract.tact

### Scan an entire directory (concurrently)

    silas ./my-ton-project/

### Force a language

    silas --lang func wallet.fc

### Output formats

    # JSON for further processing
    silas --output json ./src > report.json

    # SARIF for GitHub Code Scanning
    silas --output sarif ./src > results.sarif

    # Summary (only counts per severity and pattern)
    silas --summary ./contracts/

### Enable adaptive learning

    silas --learner ./contracts/

Weights are stored in `.silas-state.json` and updated based on your manual feedback.

---

## 🔎 Example Vulnerability Report (Text Mode)

    ✗ contract.tact: 9 issues found.
      • Missing seqno check
        Severity: CRITICAL | Line: 4
        Code: receive(msg: Message) {
        Fix: Add require(self.seqno == msg.seqno) inside receive().

      • Unprotected selfdestruct
        Severity: CRITICAL | Line: 7
        Code: selfdestruct();
        Fix: Wrap selfdestruct in a require with owner check.
    ...

    ---

## 🛡️ Pattern Categories (TON)

### Tact (45 patterns)

| Category | Examples | Severity |
|----------|----------|----------|
| Message & Network Attacks | Missing seqno, unchecked sender, signature replay, bounced handling, race condition | CRITICAL, HIGH |
| Financial & Logic | Insufficient balance, division before multiplication, overflow, double-spend | CRITICAL, HIGH |
| Access Control | Unprotected selfdestruct, missing modifier, insecure proxy, trait issues | CRITICAL, HIGH, MEDIUM |
| Gas & Storage | Gas check / pre‑check, unbounded loop, large messages, storage rent, excessive maps | CRITICAL, HIGH, MEDIUM |
| Standards & Types | Unsafe serialization, `any` type, `Address?` mismatch, fake jetton, sensitive data | HIGH, MEDIUM, LOW |
| Code Quality | Hardcoded address, raw assembler, deprecated compiler, missing invariants | INFO, HIGH |

### FunC (45 patterns)

| Category | Examples | Severity |
|----------|----------|----------|
| External Messages | Missing seqno, unchecked sender, signature, timeout, replay | CRITICAL, HIGH |
| Gas & State | Missing `impure`, raw_reserve before validation, unvalidated `set_data` | CRITICAL |
| Cell & Slice | Raw manipulation, cell depth, mutable parameters | HIGH, MEDIUM |
| Inter‑Contract | Message race, silent send failure, unhandled bounced, carry‑value | CRITICAL, HIGH |
| Jetton / NFT | Fake notifications, incorrect bounceable flag, code hash validation | CRITICAL |
| Math & Logic | Int as boolean, division before multiplication, insufficient balance | HIGH, CRITICAL |

---

## 📄 Output Formats

| Format | Flag | Use Case |
|--------|------|----------|
| **Text (colored)** | `--output text` (default) | Human‑readable terminal output |
| **JSON** | `--output json` | Integration with other tools |
| **SARIF 2.1.0** | `--output sarif` | Upload to GitHub Code Scanning |
| **Summary** | `--summary` | Quick overview of findings |

---

## 🔧 CI/CD Integration (GitHub Actions)

This repository includes a workflow (`.github/workflows/silas.yml`) that:
- Builds the scanner
- Runs it on the entire codebase (SARIF output)
- Uploads the results to **GitHub Code Scanning**

To enable it, simply fork the repo and push to `main`. The action runs automatically on every push and pull request.

---

## 📁 Project Structure

    silas-code-inspector/
    ├── cmd/silas/               # Cobra CLI entry point
    ├── internal/
    │   ├── domain/              # Issue, Pattern, Severity models
    │   ├── engine/              # Scanner factory + per‑language scanners
    │   │   ├── tact/
    │   │   ├── func/
    │   │   ├── golang/
    │   │   ├── docker/
    │   │   ├── web/
    │   │   └── gomod/
    │   ├── knowledge/           # Pattern databases (45 Tact + 45 FunC + Go + Docker + Web + CVE)
    │   ├── report/              # JSON, SARIF, text formatters
    │   └── learner/             # Adaptive learning engine
    ├── .github/workflows/       # GitHub Actions
    ├── Makefile
    └── go.mod

    ---

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feat/my-feature`).
3. Commit your changes (`git commit -m 'feat: add something'`).
4. Push to your fork (`git push origin feat/my-feature`).
5. Open a Pull Request.

Please write tests for new scanners and keep the code idiomatic.  
Check the `good first issue` tags for easy tasks.

---

## 📃 License

This project is licensed under the **MIT License** – see the [LICENSE](LICENSE) file for details.  
You are free to use, modify, and distribute this software, even commercially.

---

## 🗺️ Roadmap

- [x] Phase 1: TON scanner + CLI
- [x] Phase 2: Go, Docker, Web scanners
- [x] Phase 3: CVE feed, SARIF, adaptive learner, CI/CD
- [x] Phase 3.5: FunC support (45 patterns), comment filtering, summary mode
- [ ] Phase 4: Telegram bot + TON payments (Pro)
- [ ] Phase 5: AST‑based analysis for race conditions & reentrancy
- [ ] Phase 6: Solidity / Rust (Solana) support
- [ ] Phase 7: LSP server for IDE integration

---

**Built with ❤️ by [silasdev78](https://t.me/Silas_78)**
