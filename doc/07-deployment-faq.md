# 部署和使用示例与常见问题FAQ

## 十一、部署和使用示例（快速上手）

### 11.1 安装依赖

```bash
# 安装开发依赖
bun install --dev typescript @types/node

# 安装测试依赖
bun install --dev jest @types/jest
```

### 11.2 编译项目

```bash
# 编译TS代码
bun run build:ts

# 编译Go代码为WASM
bun run build:wasm

# 一键编译所有
bun run build
```

### 11.3 浏览器使用示例

```html
<!DOCTYPE html>
<html>
<head>
    <title>Mini Highlighter 示例</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            padding: 20px;
            background: #f5f5f5;
        }
        .code-container {
            max-width: 800px;
            margin: 0 auto;
        }
    </style>
</head>
<body>
    <div class="code-container">
        <h1>代码高亮示例</h1>
        <div id="code-container"></div>
    </div>
    
    <script type="module">
        import { codeToHtml } from './dist/ts/index.js';
        
        const code = `
const greet = (name) => {
    console.log(\`Hello, \${name}!\`);
    return true;
};

// 这是一个注释
const result = greet('World');
        `.trim();
        
        const html = await codeToHtml(code, 'javascript', 'dark');
        document.getElementById('code-container').innerHTML = html;
    </script>
</body>
</html>
```

### 11.4 Node.js使用示例

```typescript
import { codeToHtml, registerTheme } from './dist/ts/index.js';

const code = `
function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}
`.trim();

const html = await codeToHtml(code, 'javascript', 'dark');
console.log(html);
```

**Node.js 特殊处理**

如果使用 ES 模块，需要在 `package.json` 中添加：

```json
{
    "type": "module"
}
```

或者使用 CommonJS 格式：

```javascript
const { codeToHtml } = require('./dist/ts/index.js');
```

### 11.5 Bun使用示例

```typescript
import { codeToHtml } from './dist/ts/index.js';

const code = 'const x = 42;';
const html = await codeToHtml(code, 'javascript', 'light');
console.log(html);
```

Bun 直接运行：

```bash
bun run example.ts
```

### 11.6 自定义主题示例

```typescript
import { registerTheme, codeToHtml } from './dist/ts/index.js';

registerTheme({
    id: 'custom',
    name: '自定义主题',
    styles: {
        'keyword.control.js': 'color: #ff6b6b; font-weight: bold;',
        'string.quoted.double.js': 'color: #4ecdc4;',
        'string.quoted.single.js': 'color: #4ecdc4;',
        'comment.line.double-slash.js': 'color: #95a5a6; font-style: italic;',
        'comment.block.js': 'color: #95a5a6; font-style: italic;',
        'comment.block.documentation.js': 'color: #95a5a6; font-style: italic;',
        'constant.numeric.js': 'color: #f39c12;',
        'constant.language.js': 'color: #f39c12;',
        'variable.other.readwrite.js': 'color: #a29bfe;',
        'function.js': 'color: #a29bfe;',
        'keyword.operator.js': 'color: #ff6b6b;',
        'punctuation.js': 'color: #dfe6e9;',
        'text.plain': 'color: #ecf0f1;',
    }
});

const html = await codeToHtml('const x = "hello";', 'javascript', 'custom');
console.log(html);
```

### 11.7 装饰器使用示例

#### 全局装饰器示例

```typescript
import { registerDecorator, codeToHtml } from './dist/ts/index.js';

registerDecorator({
    id: 'copy-button',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<button class="copy-btn" onclick="copyCode(this)">复制</button>',
        wrapClass: 'code-with-copy'
    }
});

registerDecorator({
    id: 'code-title',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<div class="code-title">JavaScript 示例代码</div>'
    }
});

const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
```

#### 精准节点装饰器示例

```typescript
import { registerDecorator, codeToHtml } from './dist/ts/index.js';

registerDecorator({
    id: 'highlight-keywords',
    enabled: true,
    type: 'node',
    config: {
        matchRule: {
            tokenType: 'keyword.control.js',
        },
        wrapClass: 'highlighted-keyword',
        attr: {
            'data-type': 'keyword',
        }
    }
});

registerDecorator({
    id: 'highlight-strings',
    enabled: true,
    type: 'node',
    config: {
        matchRule: {
            tokenType: 'string.quoted.double.js',
        },
        wrapClass: 'highlighted-string',
        attr: {
            'data-type': 'string',
        }
    }
});

const html = await codeToHtml('const x = "hello";', 'javascript', 'dark');
```

### 11.8 完整示例：React中使用

```tsx
import React, { useState, useEffect } from 'react';
import { codeToHtml } from './dist/ts/index.js';

interface CodeHighlightProps {
    code: string;
    language?: string;
    theme?: string;
}

export const CodeHighlight: React.FC<CodeHighlightProps> = ({
    code,
    language = 'javascript',
    theme = 'dark'
}) => {
    const [html, setHtml] = useState('');
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const highlight = async () => {
            setLoading(true);
            try {
                const result = await codeToHtml(code, language, theme);
                setHtml(result);
            } catch (error) {
                console.error('代码高亮失败:', error);
                setHtml(`<pre><code>${code}</code></pre>`);
            } finally {
                setLoading(false);
            }
        };

        highlight();
    }, [code, language, theme]);

    if (loading) {
        return <div className="loading">加载中...</div>;
    }

    return (
        <div 
            className="code-highlight"
            dangerouslySetInnerHTML={{ __html: html }}
        />
    );
};
```

### 11.9 完整示例：Vue 3中使用

```vue
<template>
    <div class="code-highlight" v-html="html"></div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';
import { codeToHtml } from './dist/ts/index.js';

const props = defineProps<{
    code: string;
    language?: string;
    theme?: string;
}>();

const html = ref('');

const highlight = async () => {
    try {
        html.value = await codeToHtml(
            props.code, 
            props.language || 'javascript', 
            props.theme || 'dark'
        );
    } catch (error) {
        console.error('代码高亮失败:', error);
        html.value = `<pre><code>${props.code}</code></pre>`;
    }
};

onMounted(highlight);
watch(() => [props.code, props.language, props.theme], highlight);
</script>
```

### 11.10 服务端渲染示例

```typescript
import { codeToHtml } from './dist/ts/index.js';

async function renderServerSide(code: string, language: string, theme: string) {
    const html = await codeToHtml(code, language, theme);
    return html;
}

const html = await renderServerSide('const x = 42;', 'javascript', 'dark');
console.log(html);
```

---

## 十二、常见问题FAQ（快速排查）

### 12.1 WASM相关问题

#### Q: WASM文件加载失败，提示404错误？

A: 检查以下几点：
1. 确认 `dist/wasm/highlighter.wasm` 文件存在
2. 检查文件路径是否正确（相对路径/绝对路径）
3. 确认Web服务器正确配置了MIME类型（.wasm → application/wasm）

**解决方案：**

```typescript
// 检查文件是否存在
import { readFile } from 'fs/promises';

async function checkWasmFile() {
    try {
        const wasmPath = './dist/wasm/highlighter.wasm';
        const buffer = await readFile(wasmPath);
        console.log('WASM文件存在，大小:', buffer.length);
    } catch (error) {
        console.error('WASM文件不存在:', error);
    }
}
```

**服务器配置：**

```javascript
// Express 服务器配置
const express = require('express');
const app = express();

app.get('*.wasm', (req, res) => {
    res.set('Content-Type', 'application/wasm');
    res.sendFile(__dirname + req.path);
});
```

---

#### Q: WASM初始化超时？

A: 可能原因：
1. WASM文件过大，网络加载慢 → 考虑使用TinyGo编译
2. 浏览器不支持WASM → 检查浏览器兼容性
3. 内存不足 → 增加WebAssembly.Memory的initial值

**解决方案：**

```typescript
// 增加超时时间和内存
export async function initWasm(): Promise<void> {
    const wasmPath = '/dist/wasm/highlighter.wasm';
    const response = await fetch(wasmPath);
    
    const wasmBuffer = await response.arrayBuffer();
    const wasmModule = await WebAssembly.compile(wasmBuffer);
    
    // 增加内存大小
    const wasmImports = {
        env: {
            memory: new WebAssembly.Memory({ initial: 512, maximum: 1024 }),
        },
    };
    
    wasmInstance = await WebAssembly.instantiate(wasmModule, wasmImports);
}
```

---

#### Q: Go编译WASM失败？

A: 检查Go版本（1.11+支持WASM）：

```bash
go version
# 确保版本 >= 1.11
```

**解决方案：**

```bash
# 升级Go版本
# macOS
brew upgrade go

# Linux
sudo apt-get update
sudo apt-get upgrade golang

# Windows
# 下载安装包：https://golang.org/dl/
```

---

### 12.2 解析相关问题

#### Q: 某些语法没有被正确高亮？

A: 可能原因：
1. 语法规则未覆盖 → 检查语言配置的scopeMap
2. 状态机逻辑有误 → 查看Go端parseLine方法
3. 主题样式未定义 → 检查主题配置的styles

**解决方案：**

```typescript
// 检查Token类型
import { codeToTokens } from './dist/ts/index.js';

const code = 'const x = 1;';
const tokens = await codeToTokens(code);

console.log('Token列表:', tokens.map(t => `${t.type}: ${t.value}`));
// 输出：keyword.control.js: const, text.plain:  , variable.other.readwrite.js: x, ...
```

---

#### Q: 解析速度很慢？

A: 优化建议：
1. 使用TinyGo编译减小WASM体积
2. 减少正则表达式使用，改用字符串匹配
3. 避免在循环中创建临时对象

**使用TinyGo：**

```bash
# 安装TinyGo
brew install tinygo

# 编译
cd src/go
tinygo build -o ../../dist/wasm/highlighter.wasm -target wasm highlighter.go
```

---

#### Q: 语法错误导致解析崩溃？

A: 检查Go端是否有panic捕获：

```go
// 在parseRoot方法中增加容错逻辑
func (h *Highlighter) parseRoot(line string, pos int) int {
    defer func() {
        if r := recover(); r != nil {
            // 捕获panic，恢复到ROOT状态
            h.currentState = STATE_ROOT
            fmt.Printf("解析错误已恢复: %v\n", r)
        }
    }()
    
    // 原有解析逻辑...
}
```

---

### 12.3 平台兼容性问题

#### Q: Node.js环境下WASM加载失败？

A: Node.js需要特殊处理：

```typescript
import { readFile } from 'fs/promises';
import { dirname, resolve } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const wasmPath = resolve(__dirname, '../wasm/highlighter.wasm');
const wasmBuffer = await readFile(wasmPath);

// 使用 buffer 创建 WebAssembly.Module
const wasmModule = await WebAssembly.compile(wasmBuffer);
const wasmImports = {
    env: {
        memory: new WebAssembly.Memory({ initial: 256 }),
    },
};
const wasmInstance = await WebAssembly.instantiate(wasmModule, wasmImports);
```

---

#### Q: Bun环境下API调用报错？

A: Bun与Node.js的模块系统略有不同，确保：
1. 使用ES模块（.mjs或package.json设置type: module）
2. 正确处理相对路径

**Bun 配置文件：**

```json
{
    "type": "module"
}
```

---

### 12.4 扩展相关问题

#### Q: 新增语言后无法使用？

A: 检查步骤：
1. 确认调用了 `registerLanguage(config)`
2. 确认语言ID正确（与config.id一致）
3. 确认scopeMap映射完整

**调试代码：**

```typescript
import { getRegisteredLanguages, getLanguage } from './dist/ts/index.js';

// 检查已注册的语言
console.log('已注册的语言:', getRegisteredLanguages());

// 检查特定语言
const lang = getLanguage('python');
console.log('Python语言配置:', lang);
```

---

#### Q: 主题切换不生效？

A: 检查步骤：
1. 确认调用了 `setTheme(themeId)`
2. 确认主题ID正确（与config.id一致）
3. 确认styles映射包含所有需要的token type

**调试代码：**

```typescript
import { getRegisteredThemes, getCurrentTheme } from './dist/ts/index.js';

// 检查已注册的主题
console.log('已注册的主题:', getRegisteredThemes());

// 检查当前主题
console.log('当前主题:', getCurrentTheme());

// 重新设置主题
setTheme('dark');
```

---

#### Q: 装饰器没有生效？

A: 检查步骤：
1. 确认装饰器的enabled为true
2. 确认装饰器已注册（`registerDecorator`）
3. 检查装饰器配置是否正确（matchRule、wrapClass等）

**调试代码：**

```typescript
import { getEnabledDecorators } from './dist/ts/index.js';

// 检查已启用的装饰器
const decorators = getEnabledDecorators();
console.log('已启用的装饰器:', decorators.map(d => d.id));
```

---

### 12.5 性能相关问题

#### Q: 解析大文件时内存占用过高？

A: 优化建议：
1. 实现流式解析（逐行处理）
2. 及时释放临时对象
3. 减少Token数组的拷贝

**优化代码：**

```go
// 使用流式解析，减少内存占用
func (h *Highlighter) CodeToTokens(code string) []Token {
    h.tokens = make([]Token, 0, len(code)/10) // 预分配内存
    h.currentState = STATE_ROOT
    
    // 逐行解析
    lines := strings.Split(code, "\n")
    for lineNum, line := range lines {
        h.parseLine(line, lineNum)
        if lineNum < len(lines)-1 {
            h.addToken("text.plain", "\n")
        }
    }
    
    return h.tokens
}
```

---

#### Q: 首次加载WASM文件很慢？

A: 优化建议：
1. 预加载WASM文件
2. 使用CDN加速
3. 考虑使用Service Worker缓存

**预加载：**

```html
<!-- 在HTML head中添加 -->
<link rel="preload" href="/dist/wasm/highlighter.wasm" as="fetch" type="application/wasm">
```

**CDN配置：**

```typescript
const WASM_CDN = 'https://cdn.example.com/wasm/';

export async function initWasm(): Promise<void> {
    const response = await fetch(WASM_CDN + 'highlighter.wasm');
    // ...
}
```

---

### 12.6 错误处理相关

#### Q: 如何处理所有可能的错误？

A: 使用try-catch包裹所有API调用：

```typescript
import { codeToHtml, codeToTokens } from './dist/ts/index.js';

async function safeHighlight(code: string, language: string, theme: string) {
    try {
        // 初始化WASM
        await initWasm();
        
        // 调用API
        const html = await codeToHtml(code, language, theme);
        return html;
    } catch (error) {
        if (error instanceof Error) {
            if (error.message.includes('WASM')) {
                console.error('WASM初始化失败:', error.message);
                return `<pre><code>${code}</code></pre>`;
            }
            if (error.message.includes('语言')) {
                console.error('语言未注册:', error.message);
                return `<pre><code>${code}</code></pre>`;
            }
            if (error.message.includes('主题')) {
                console.error('主题未注册:', error.message);
                return `<pre><code>${code}</code></pre>`;
            }
        }
        console.error('未知错误:', error);
        return `<pre><code>${code}</code></pre>`;
    }
}
```

---

## 总结

本章节提供了完整的部署和使用指南以及常见问题的解决方案：

### 部署和使用
- 安装依赖和编译项目的命令
- 浏览器、Node.js、Bun 三个平台的使用示例
- 自定义主题和装饰器的使用示例
- React 和 Vue 框架中的集成示例
- 服务端渲染示例

### 常见问题
- WASM相关问题（加载失败、超时、编译失败）
- 解析相关问题（语法高亮、解析速度、语法错误）
- 平台兼容性问题（Node.js、Bun）
- 扩展相关问题（语言、主题、装饰器）
- 性能相关问题（内存占用、加载速度）
- 错误处理相关

通过本章节的内容，可以快速上手项目开发，并解决开发过程中遇到的常见问题。

**相关文档：**
- [产品说明](./01-product-overview.md)
- [技术栈与目录结构](./02-tech-stack.md)
- [分阶段实现步骤](./03-implementation-steps.md)
- [核心功能设计](./04-core-features.md)
- [核心代码与扩展指南](./05-core-code.md)
- [MVP与测试策略](./06-mvp-testing.md)
