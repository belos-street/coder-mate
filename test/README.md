# WASM Hello World Demo

## 文件说明

- `index.html` - 主页面
- `app.js` - WASM 加载和调用逻辑
- `wasm_exec.js` - Go WASM 运行时支持
- `../core/wasm/highlighter.wasm` - 编译后的 WASM 文件

## 使用方法

1. 使用 IDE 的 Live Server 打开 `index.html`
2. 等待 WASM 模块加载完成
3. 点击 "Call WASM Function" 按钮
4. 查看结果

## 重新编译 WASM

```bash
cd ../core
chmod +x build.sh
./build.sh
```
