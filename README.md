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

Download the macOS binary from the  
**[Releases page](https://github.com/tomitahisaki/annotrax/releases)**.

### 1. Make the binary executable

```bash
chmod +x annotrax
```

### 2. Move it to a directory in your PATH

#### Option A (recommended): `/usr/local/bin`

```bash
sudo mv annotrax /usr/local/bin/annotrax
```

#### Option B: Use `~/bin`

```bash
mkdir -p ~/bin
mv annotrax ~/bin/annotrax
```

Add `~/bin` to your PATH (only once):

```bash
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### 3. Verify installation

```bash
annotrax -h
```

---

## 🚀 Usage

### Scan the current directory

```bash
annotrax
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
annotrax -dir ./src
```

---

## 🔤 Supported annotation keywords

```
TODO
FIXME
NOTE
HACK
XXX
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
GOOS=darwin GOARCH=arm64 go build -o annotrax
```

---

## 📄 License

This project is licensed under the terms of the  
[MIT License](./LICENSE).
