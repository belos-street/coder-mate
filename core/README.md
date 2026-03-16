# Core - Go + WASM Highlighter

## 编译 WASM

```bash
cd core
chmod +x build.sh
./build.sh
```

或者手动编译：

```bash
GOOS=js GOARCH=wasm go build -o wasm/highlighter.wasm highlighter.go
```

## 文件说明

- `highlighter.go` - 核心高亮逻辑，导出 `highlightCode` 函数到 JS
- `wasm/highlighter.wasm` - 编译后的 WASM 二进制文件
- `build.sh` - WASM 编译脚本

## JS 调用示例

```javascript
const highlightCode = (code, language, theme) => {
  return window.highlightCode(code, language, theme);
};
```