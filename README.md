# annotrax

Annotrax is a lightweight **macOS-only** CLI tool that scans your project for
annotation-style comments such as `TODO`, `FIXME`, `NOTE`, `HACK`, and `XXX`.

This tool helps developers quickly locate pending tasks, warnings, or temporary
notes left inside source code.

---

## ✨ Features

- **macOS-only** single binary
- Recursively scans directories
- Detects common annotation keywords:
  - `TODO`
  - `FIXME`
  - `NOTE`
  - `HACK`
  - `XXX`
- Outputs results in the format:

  ```
  path:line: [KEYWORD] message
  ```

- Skips common binary file types automatically

---

## 📦 Installation

annotrax is distributed as a standalone macOS binary.
You do not need Go installed to use it.

### 1. Download the binary

Download the latest release from the
**[Releases page](https://github.com/tomitahisaki/annotrax/releases)**.

### 2. Make it executable

```bash
chmod +x annotrax-macos
```

### 3. (Option A — recommended) Install to /usr/local/bin

```bash
sudo mv annotrax-macos /usr/local/bin/annotrax
```

---

## 🚀 Usage

### Scan the current directory

```bash
annotrax-macos
```

Example output:

```
app/models/user.rb:42: [TODO] # TODO: improve validation
frontend/src/App.vue:88: [FIXME] // FIXME: type handling
backend/main.go:12: [NOTE] // NOTE: temporary workaround
```

---

### Scan a specific directory

```bash
annotrax-macos -dir ./src
```

---

## 🔤 Supported annotation keywords

```
TODO
FIXME
NOTE
```

Future versions will allow custom keywords via configuration.

---

## 🗺 Roadmap / Ideas

```
- Add default ignore directories (.git, node_modules, vendor, etc.)
- Add file extension filtering (-ext .go,.rb,.ts,.vue)
- Add Markdown and JSON output formats
- Add annotrax.yaml configuration support
- macOS-specific integrations (QuickLook plugin)
- VSCode / Neovim plugin integration
```

---

## 🛠 Development (for contributors)

Annotrax is written in Go.  
To build or modify the tool:

### Clone the repository

```bash
git clone https://github.com/<your-name>/annotrax.git
cd annotrax
go mod tidy
```

### Run in development mode

```bash
go run .
```

## 🧪 Running Tests

This project includes basic unit tests to ensure the core logic works correctly.

To run all tests:

```bash
go test ./...
```

You should see output similar to:

```
ok   github.com/yourname/annotrax 0.15s
```

If you plan to contribute or modify the tool, please make sure all tests pass before submitting changes.

### Build macOS binary

```bash
GOOS=darwin GOARCH=arm64 go build -o annotrax-macos
```

---

## 📄 License

This project is licensed under the terms of the  
[MIT License](./LICENSE).
