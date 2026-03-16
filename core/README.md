# Core

Code syntax highlighting WASM module.

## Directory Structure

```
core/
├── main.go          # Entry point, exports highlightCode function
├── parser/          # Code parsing logic
│   └── parser.go
├── renderer/        # Code rendering logic
│   └── renderer.go
├── go.mod           # Go module definition
└── .gitignore      # Git ignore rules
```

## Build

```bash
./scripts/build-wasm
```

Output: `bin/highlighter.wasm`

## Usage

The WASM module exposes a `highlightCode(code, language, theme)` function to JavaScript.

```javascript
const wasm = await WebAssembly.instantiateStreaming(fetch('bin/highlighter.wasm'));
const result = wasm.instance.exports.highlightCode('const x = 1', 'javascript', 'default');
```

## Development

```bash
cd core
go mod tidy
```
