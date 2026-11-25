<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:111827,100:2563eb&height=160&section=header&text=Go%20Context%20Cancellation%20Lab&fontSize=30&fontColor=ffffff" alt="Go Context Cancellation Lab header"/></p>
<p align="center"><img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=18&pause=1000&color=60A5FA&center=true&vCenter=true&width=760&lines=See%20cancellation%2C%20deadlines%2C%20and%20timeouts%20propagate%20from%20a%20parent%20context%20to%20worker%20goroutines." alt="Animated tagline"/></p>
<p align="center"><a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-2563eb.svg" alt="MIT"/></a> <img src="https://img.shields.io/badge/language-Go-0ea5e9.svg" alt="Go"/> <img src="https://img.shields.io/github/stars/CyberLoop-dev/go-context-cancel-lab?style=flat" alt="Stars"/> <img src="https://img.shields.io/github/last-commit/CyberLoop-dev/go-context-cancel-lab?style=flat" alt="Last commit"/></p>
<p align="center"><img src="assets/mark.svg" width="110" alt="Custom project mark"/></p>

## 🧭 Overview
See cancellation, deadlines, and timeouts propagate from a parent context to worker goroutines. The examples use the Go standard library and short, bounded demonstrations.

## ✨ Features
- 🛑 Compare explicit cancellation, absolute deadlines, and relative timeouts
- 🧵 Watch multiple workers stop when their context is done
- 🧭 Inspect context errors and parent-to-child cancellation

## 🎬 Demo
Run the CLI and compare the printed event order and context/channel state. No hosted demo is claimed.

## 🚀 Quick Start
```bash
go run .
```
Use `go run . -scenario cancel`, `-scenario deadline`, `-scenario timeout`, or `-scenario all`.

## 🧰 Tech Stack
- 🧩 Go standard library
- 🔒 No network access or external services
- 📦 No third-party dependencies

## 🗂️ Folder Structure
```text
go-context-cancel-lab/
├── assets/mark.svg
├── main.go
├── main_test.go
├── go.mod
├── .gitignore
├── LICENSE
└── README.md
```

## 🛣️ Roadmap
- [ ] Add a JSON output mode
- [ ] Add more guided scenarios

## 🤝 Contributing
Issues and pull requests are welcome. Keep examples deterministic where possible and explain any timing-dependent behavior.

## 📄 License
Released under the [MIT License](LICENSE).

<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:2563eb,100:111827&height=105&section=footer" alt="Animated wave footer"/></p>