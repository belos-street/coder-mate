# 分阶段完整实现步骤

## 整体开发原则

按「**基础层 → 核心层 → 扩展层 → 优化层**」分阶段开发，每完成一个阶段即可测试对应能力，无依赖阻塞；所有阶段的核心逻辑均基于 TextMate 规范，无黑盒代码。

---

## 阶段 1：环境初始化与基础结构搭建

**优先级：★★★★★（前置必做）**

### 任务清单

#### 1. 初始化 TS 环境
创建 `tsconfig.json`，配置编译目标为 ES6，输出目录为 `lib/dist`，开启类型声明。

**tsconfig.json 配置示例：**
```json
{
  "compilerOptions": {
    "target": "ES6",
    "module": "ESNext",
    "outDir": "./lib/dist",
    "declaration": true,
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true
  },
  "include": ["lib/src/**/*"],
  "exclude": ["node_modules", "lib/dist", "core"]
}
```

#### 2. 初始化 Go 环境
在 `core/src` 目录执行 `go mod init highlighter`，生成 `go.mod`，无需额外依赖（Go 原生支持 WASM）。

```bash
cd core/src
go mod init highlighter
```

**生成的 go.mod 文件：**
```go
module highlighter

go 1.21
```

#### 3. 创建完整的目录结构
按「标准化目录结构」创建所有目录与空文件，保证后续开发的规范性。

```bash
# 创建 TypeScript 源码目录
mkdir -p lib/src/types
mkdir -p lib/src/core
mkdir -p lib/src/language
mkdir -p lib/src/theme
mkdir -p lib/src/decorator
mkdir -p lib/src/wasm
mkdir -p lib/src/utils

# 创建 Go 源码目录
mkdir -p core/src

# 创建编译输出目录
mkdir -p lib/dist
mkdir -p core/wasm

# 创建空文件
touch lib/src/types/index.ts
touch lib/src/core/index.ts
touch lib/src/language/index.ts
touch lib/src/language/js.ts
touch lib/src/theme/index.ts
touch lib/src/theme/dark.ts
touch lib/src/theme/light.ts
touch lib/src/decorator/index.ts
touch lib/src/wasm/loader.ts
touch lib/src/wasm/bridge.ts
touch lib/src/utils/index.ts
touch lib/src/index.ts
touch core/src/highlighter.go
```

#### 4. 编写 Go-WASM 编译脚本
新增 `build:wasm` 命令，编译 Go 代码为 WASM 二进制文件到 `dist/wasm`。

**在 package.json 中添加：**
```json
{
  "scripts": {
    "build:wasm": "cd src/go && GOOS=js GOARCH=wasm go build -o ../../dist/wasm/highlighter.wasm highlighter.go"
  }
}
```

#### 5. 编写 TS 编译脚本
新增 `build:ts` 命令，编译 TS 代码到 `dist/ts`。

**在 package.json 中添加：**
```json
{
  "scripts": {
    "build:ts": "tsc"
  }
}
```

### 验收标准

- [ ] tsconfig.json 配置正确
- [ ] go.mod 文件已生成
- [ ] 所有目录和空文件已创建
- [ ] `bun run build:ts` 可以成功编译（即使文件为空）
- [ ] `bun run build:wasm` 可以成功编译（即使文件为空）

---

## 阶段 2：核心基础层开发

**优先级：★★★★★（核心必做）**

> 本阶段完成「核心类型定义 + Go 状态机解析 + WASM 桥接」，是整个库的基石，所有上层能力均依赖此阶段的输出。

### 任务清单

#### 1. TS 核心类型定义
在 `src/ts/types/index.ts` 中定义「Token、语言配置、主题配置、装饰器配置」的标准 TypeScript 类型，所有类型均遵循行业规范。

**src/ts/types/index.ts：**
```typescript
export interface Token {
  type: string;
  value: string;
}

export interface LanguageConfig {
  id: string;
  name: string;
  scopeMap: Record<string, string>;
}

export interface ThemeConfig {
  id: string;
  name: string;
  styles: Record<string, string>;
}

export interface DecoratorConfig {
  id: string;
  enabled: boolean;
  type: 'global' | 'node';
  config: {
    wrapClass?: string;
    prependHtml?: string;
    appendHtml?: string;
    matchRule?: {
      tokenType?: string;
      lineRange?: [number, number];
      content?: string;
    };
    attr?: Record<string, string>;
  };
}
```

#### 2. Go 核心解析逻辑开发
在 `src/go/highlighter.go` 中实现：
- 定义标准 Token 结构体，与 TS 的 Token 类型完全对齐
- 实现 TextMate 有限状态机（基础状态+嵌套子状态，如 JSDoc 嵌套在多行注释中）
- 实现纯文本切分、语法规则匹配、Token 数组收集
- 导出 WASM 方法 `codeToTokens`，将 Token 数组序列化为 JSON 字符串返回给 TS

**核心代码结构：**
```go
package main

import (
    "encoding/json"
    "regexp"
    "strings"
    "syscall/js"
)

type Token struct {
    Type  string `json:"type"`
    Value string `json:"value"`
}

const (
    STATE_ROOT          = "root"
    STATE_STRING_DOUBLE = "string_double"
    STATE_STRING_SINGLE = "string_single"
    STATE_COMMENT_LINE  = "comment_line"
    STATE_COMMENT_BLOCK = "comment_block"
    STATE_JSDOC         = "jsdoc"
)

type Highlighter struct {
    currentState string
    tokens       []Token
}

func NewHighlighter() *Highlighter {
    return &Highlighter{
        currentState: STATE_ROOT,
        tokens:       make([]Token, 0),
    }
}

func (h *Highlighter) addToken(tp, val string) {
    if val == "" {
        return
    }
    h.tokens = append(h.tokens, Token{Type: tp, Value: val})
}

func (h *Highlighter) CodeToTokens(code string) []Token {
    // 状态机核心逻辑
    // 详见完整实现
}

func codeToTokensWrapper(this js.Value, args []js.Value) interface{} {
    if len(args) < 1 {
        return "[]"
    }
    code := args[0].String()
    h := NewHighlighter()
    tokens := h.CodeToTokens(code)
    jsonBytes, _ := json.Marshal(tokens)
    return string(jsonBytes)
}

func main() {
    js.Global().Set("codeToTokens", js.FuncOf(codeToTokensWrapper))
    <-make(chan bool)
}
```

#### 3. Go 编译为 WASM
执行编译命令，生成 `dist/wasm/highlighter.wasm` 二进制文件。

```bash
bun run build:wasm
```

#### 4. TS WASM 桥接开发
在 `src/ts/wasm/loader.ts` 实现 WASM 初始化加载，在 `src/ts/wasm/bridge.ts` 实现调用 Go 的 WASM 方法，解析 JSON 为标准 Token 数组。

**src/ts/wasm/loader.ts：**
```typescript
let wasmInstance: WebAssembly.Instance | null = null;
let wasmInitialized = false;

export async function initWasm(): Promise<void> {
    if (wasmInitialized) return;

    try {
        const wasmPath = '/dist/wasm/highlighter.wasm';
        const response = await fetch(wasmPath);
        
        if (!response.ok) {
            throw new Error(`WASM文件加载失败: ${response.status} ${response.statusText}`);
        }

        const wasmBuffer = await response.arrayBuffer();
        const wasmModule = await WebAssembly.compile(wasmBuffer);
        const wasmImports = {
            env: {
                memory: new WebAssembly.Memory({ initial: 256 }),
            },
        };
        
        wasmInstance = await WebAssembly.instantiate(wasmModule, wasmImports);
        wasmInitialized = true;
    } catch (error) {
        console.error('WASM初始化失败:', error);
        throw error;
    }
}

export function getWasmInstance(): WebAssembly.Instance {
    if (!wasmInstance || !wasmInitialized) {
        throw new Error('WASM未初始化，请先调用 initWasm()');
    }
    return wasmInstance;
}
```

**src/ts/wasm/bridge.ts：**
```typescript
import { getWasmInstance } from './loader';
import { Token } from '../types';

export async function codeToTokens(code: string): Promise<Token[]> {
    const wasmInstance = getWasmInstance();
    const jsonString = wasmInstance.exports.codeToTokens(code) as string;
    return JSON.parse(jsonString) as Token[];
}
```

### 验收标准

- [ ] TS 类型定义完整且正确
- [ ] Go 状态机能够正确解析 JavaScript 代码
- [ ] WASM 文件编译成功
- [ ] WASM 加载和初始化成功
- [ ] 能够调用 Go 方法并返回正确的 Token 数组

---

## 阶段 3：核心 API 封装

**优先级：★★★★★（核心必做）**

> 本阶段完成「codeToTokens、codeToHtml」两个核心基础 API 的封装，是库的对外核心能力，全平台通用。

### 任务清单

#### 1. 在 `src/ts/core/index.ts` 中封装核心方法
- `codeToTokens`: 调用 WASM 桥接方法，返回标准 Token 数组
- `codeToHtml`: 基于 Token 数组 + 主题配置，映射生成可直接渲染的高亮 HTML 字符串

**src/ts/core/index.ts：**
```typescript
import { codeToTokens as wasmCodeToTokens } from '../wasm/bridge';
import { getTheme } from '../theme';
import { Token } from '../types';

export const codeToTokens = async (code: string, lang: string = 'javascript'): Promise<Token[]> => {
    if (!code) return [];
    return await wasmCodeToTokens(code);
};

export const codeToHtml = async (code: string, lang: string = 'javascript', themeId: string = 'dark'): Promise<string> => {
    const tokens = await codeToTokens(code, lang);
    const theme = getTheme(themeId);
    if (!theme) throw new Error(`主题${themeId}未注册`);
    
    let html = tokens.map(token => {
        const style = theme.styles[token.type] || theme.styles['text.plain'] || '';
        return `<span style="${style}">${token.value}</span>`;
    }).join('');

    return `<pre class="mini-highlighter" style="background: #161b22; padding: 16px; border-radius: 8px; font-family: Consolas, monospace; font-size: 14px;"><code>${html}</code></pre>`;
};
```

#### 2. 在 `src/ts/index.ts` 中全局暴露这两个 API
保证调用简洁。

**src/ts/index.ts：**
```typescript
import { initWasm } from './wasm/loader';
import { codeToTokens, codeToHtml } from './core';
import { registerLanguage, getLanguage } from './language';
import { registerTheme, getTheme, setTheme } from './theme';
import { registerDecorator } from './decorator';

initWasm().catch(err => console.error('WASM初始化失败:', err));

export { codeToTokens, codeToHtml };
export { registerLanguage, getLanguage };
export { registerTheme, getTheme, setTheme };
export { registerDecorator };
```

#### 3. 测试核心能力
编写测试代码，验证 Token 数组生成的正确性、HTML 渲染的高亮效果。

**tests/core.test.ts：**
```typescript
import { codeToTokens, codeToHtml } from '../src/ts';

describe('核心API测试', () => {
    test('codeToTokens - 正确解析关键字', async () => {
        const tokens = await codeToTokens('const x = 1;');
        expect(tokens).toContainEqual({ type: 'keyword.control.js', value: 'const' });
    });

    test('codeToHtml - 生成完整HTML', async () => {
        const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
        expect(html).toContain('<pre');
        expect(html).toContain('<code');
        expect(html).toContain('</code>');
        expect(html).toContain('</pre>');
    });
});
```

### 验收标准

- [ ] `codeToTokens` 能够正确返回 Token 数组
- [ ] `codeToHtml` 能够生成完整的 HTML
- [ ] API 调用方式简洁易用
- [ ] 测试用例全部通过

---

## 阶段 4：扩展能力开发 - 多语言 + 主题接口

**优先级：★★★★（核心扩展）**

> 本阶段完成「多语言注册接口、主题注册与切换接口」，实现「开闭原则」：对扩展开放，对修改关闭，核心代码永不改动。

### 任务清单

#### 1. 多语言接口开发
在 `src/ts/language/index.ts` 实现语言注册、切换、获取方法，新增语言仅需在此目录新增规则文件，无需修改核心代码。

**src/ts/language/index.ts：**
```typescript
import { LanguageConfig } from '../types';

const languages = new Map<string, LanguageConfig>();
let defaultLanguage = 'javascript';

export function registerLanguage(config: LanguageConfig): void {
    languages.set(config.id, config);
}

export function getLanguage(langId: string): LanguageConfig | null {
    return languages.get(langId) || null;
}

export function setDefaultLanguage(langId: string): void {
    if (!languages.has(langId)) {
        throw new Error(`语言${langId}未注册`);
    }
    defaultLanguage = langId;
}

export function getRegisteredLanguages(): string[] {
    return Array.from(languages.keys());
}
```

**src/ts/language/js.ts：**
```typescript
import { registerLanguage } from './index';

registerLanguage({
    id: 'javascript',
    name: 'JavaScript',
    scopeMap: {
        'const': 'keyword.control.js',
        'let': 'keyword.control.js',
        'var': 'keyword.control.js',
        'function': 'keyword.control.js',
        '"': 'string.quoted.double.js',
        "'": 'string.quoted.single.js',
        '//': 'comment.line.double-slash.js',
        '/*': 'comment.block.js',
        '/**': 'comment.block.documentation.js',
    }
});
```

#### 2. 主题接口开发
在 `src/ts/theme/index.ts` 实现主题注册、切换、获取方法，新增主题仅需在此目录新增主题配置文件，实现「Token 语法作用域 与 样式 完全解耦」。

**src/ts/theme/index.ts：**
```typescript
import { ThemeConfig } from '../types';

const themes = new Map<string, ThemeConfig>();
let currentTheme = 'dark';

export function registerTheme(config: ThemeConfig): void {
    themes.set(config.id, config);
}

export function getTheme(themeId: string): ThemeConfig | null {
    return themes.get(themeId) || null;
}

export function setTheme(themeId: string): void {
    if (!themes.has(themeId)) {
        throw new Error(`主题${themeId}未注册`);
    }
    currentTheme = themeId;
}

export function getCurrentTheme(): string {
    return currentTheme;
}

export function getRegisteredThemes(): string[] {
    return Array.from(themes.keys());
}
```

**src/ts/theme/dark.ts：**
```typescript
import { registerTheme } from './index';

registerTheme({
    id: 'dark',
    name: '默认暗色主题',
    styles: {
        'keyword.control.js': 'color: #ff7b72; font-weight: bold;',
        'string.quoted.double.js': 'color: #a5d6ff;',
        'string.quoted.single.js': 'color: #a5d6ff;',
        'comment.line.double-slash.js': 'color: #8b949e; font-style: italic;',
        'comment.block.js': 'color: #8b949e; font-style: italic;',
        'comment.block.documentation.js': 'color: #8b949e; font-style: italic;',
        'constant.numeric.js': 'color: #79c0ff;',
        'variable.other.readwrite.js': 'color: #d2a8ff;',
        'text.plain': 'color: #c9d1d9;',
    }
});
```

**src/ts/theme/light.ts：**
```typescript
import { registerTheme } from './index';

registerTheme({
    id: 'light',
    name: '默认亮色主题',
    styles: {
        'keyword.control.js': 'color: #cf222e; font-weight: bold;',
        'string.quoted.double.js': 'color: #0a3069;',
        'string.quoted.single.js': 'color: #0a3069;',
        'comment.line.double-slash.js': 'color: #6e7781; font-style: italic;',
        'comment.block.js': 'color: #6e7781; font-style: italic;',
        'comment.block.documentation.js': 'color: #6e7781; font-style: italic;',
        'constant.numeric.js': 'color: #0550ae;',
        'variable.other.readwrite.js': 'color: #953800;',
        'text.plain': 'color: #24292f;',
    }
});
```

#### 3. 适配核心 API
修改 `codeToTokens/codeToHtml`，支持传入语言标识、主题名称，实现多语言+多主题的灵活切换。

### 验收标准

- [ ] 能够注册新语言
- [ ] 能够注册新主题
- [ ] 能够切换默认语言和主题
- [ ] 核心 API 支持语言和主题参数

---

## 阶段 5：扩展能力开发 - 代码装饰器核心能力

**优先级：★★★★（核心扩展）**

> 本阶段完成「全局装饰 + 精准节点装饰」两类核心装饰能力，装饰逻辑与解析/渲染逻辑完全解耦，不破坏原有高亮结构。

### 任务清单

#### 1. 在 `src/ts/decorator/index.ts` 实现装饰器注册、执行、管理方法

**src/ts/decorator/index.ts：**
```typescript
import { DecoratorConfig } from '../types';

const decorators = new Map<string, DecoratorConfig>();

export function registerDecorator(decorator: DecoratorConfig): void {
    decorators.set(decorator.id, decorator);
}

export function toggleDecorator(decoratorId: string, enabled: boolean): void {
    const decorator = decorators.get(decoratorId);
    if (!decorator) {
        throw new Error(`装饰器${decoratorId}未注册`);
    }
    decorator.enabled = enabled;
}

export function getEnabledDecorators(): DecoratorConfig[] {
    return Array.from(decorators.values()).filter(d => d.enabled);
}

export function applyDecorators(html: string, themeId: string): string {
    const enabledDecorators = getEnabledDecorators();
    let result = html;
    
    for (const decorator of enabledDecorators) {
        try {
            result = applySingleDecorator(result, decorator);
        } catch (error) {
            console.warn(`装饰器"${decorator.id}"执行失败，已跳过:`, error);
        }
    }
    
    return result;
}

function applySingleDecorator(html: string, decorator: DecoratorConfig): string {
    if (decorator.type === 'global') {
        return applyGlobalDecorator(html, decorator);
    } else if (decorator.type === 'node') {
        return applyNodeDecorator(html, decorator);
    }
    return html;
}

function applyGlobalDecorator(html: string, decorator: DecoratorConfig): string {
    const { prependHtml, appendHtml, wrapClass } = decorator.config;
    
    let result = html;
    
    if (prependHtml) {
        result = prependHtml + result;
    }
    
    if (appendHtml) {
        result = result + appendHtml;
    }
    
    if (wrapClass) {
        result = `<div class="${wrapClass}">${result}</div>`;
    }
    
    return result;
}

function applyNodeDecorator(html: string, decorator: DecoratorConfig): string {
    const { matchRule, wrapClass, attr } = decorator.config;
    
    if (!matchRule) {
        return html;
    }
    
    let result = html;
    
    if (matchRule.tokenType) {
        const regex = new RegExp(`<span style="[^"]*">([^<]+)</span>`, 'g');
        result = result.replace(regex, (match, content) => {
            if (content.match(matchRule.tokenType)) {
                return `<span class="${wrapClass}" ${formatAttributes(attr)}>${content}</span>`;
            }
            return match;
        });
    }
    
    return result;
}

function formatAttributes(attrs?: Record<string, string>): string {
    if (!attrs) return '';
    return Object.entries(attrs)
        .map(([key, value]) => `${key}="${value}"`)
        .join(' ');
}
```

#### 2. 实现两类装饰器的核心逻辑
- 全局装饰（代码块前后追加内容、外层包裹自定义容器）
- 精准节点装饰（按 Token 类型/位置/内容匹配，包裹自定义类/追加属性）

#### 3. 适配 `codeToHtml` 方法
支持传入装饰器配置，实现装饰后的 HTML 生成。

### 验收标准

- [ ] 能够注册新装饰器
- [ ] 能够启用/禁用装饰器
- [ ] 全局装饰能够正确应用
- [ ] 精准节点装饰能够正确应用
- [ ] 装饰器执行错误不影响其他装饰器

---

## 阶段 6：全平台兼容性测试与性能优化

**优先级：★★★（收尾优化）**

### 任务清单

#### 1. 兼容性测试
在浏览器、Node.js、Bun 三个平台分别测试所有 API，保证调用方式一致、结果正确。

**浏览器测试：**
```html
<!DOCTYPE html>
<html>
<head>
    <title>浏览器测试</title>
</head>
<body>
    <div id="output"></div>
    <script type="module">
        import { codeToHtml } from './dist/ts/index.js';
        const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
        document.getElementById('output').innerHTML = html;
    </script>
</body>
</html>
```

**Node.js 测试：**
```typescript
import { codeToHtml } from './dist/ts/index.js';
const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
console.log(html);
```

**Bun 测试：**
```typescript
import { codeToHtml } from './dist/ts/index.js';
const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
console.log(html);
```

#### 2. 性能测试
使用几万行超大 JS 文件测试解析性能，验证流式解析的内存占用、解析耗时。

**性能测试代码：**
```typescript
import { codeToTokens } from './dist/ts/index.js';
import Benchmark from 'benchmark';

const sampleCode = `
const x = 1;
const y = 2;
function add(a, b) {
    return a + b;
}
`.repeat(1000); // 约4000行代码

const suite = new Benchmark.Suite();

suite
    .add('codeToTokens 4000行代码', async () => {
        await codeToTokens(sampleCode);
    })
    .on('cycle', (event) => {
        console.log(String(event.target));
    })
    .run();
```

#### 3. 容错性测试
测试有语法错误的代码，验证状态机的容错能力（语法错误不影响高亮）。

**容错性测试代码：**
```typescript
import { codeToTokens } from './dist/ts/index.js';

const errorCodes = [
    'const x = ;',
    'function test( {',
    'const y = "unclosed string',
    '// unclosed comment',
];

for (const code of errorCodes) {
    try {
        const tokens = await codeToTokens(code);
        console.log('✓ 容错测试通过:', code);
    } catch (error) {
        console.error('✗ 容错测试失败:', code, error);
    }
}
```

#### 4. 体积优化
使用 TinyGo 编译 Go 代码，减小 WASM 文件体积；TS 代码开启压缩编译。

**TinyGo 编译：**
```bash
# 安装 TinyGo
brew install tinygo

# 使用 TinyGo 编译
cd src/go
tinygo build -o ../../dist/wasm/highlighter.wasm -target wasm highlighter.go
```

**TS 代码压缩：**
```json
{
  "scripts": {
    "build:ts:prod": "tsc && terser dist/ts/**/*.js --compress --mangle -o dist/ts/**/*.min.js"
  }
}
```

### 验收标准

- [ ] 浏览器、Node.js、Bun 三个平台测试通过
- [ ] 性能测试达到预期（解析1000行代码<100ms）
- [ ] 容错性测试通过（语法错误不崩溃）
- [ ] WASM 文件体积优化（<500KB）
- [ ] TS 代码压缩成功

---

## 总结

通过以上6个阶段的开发，你将完成一个完整的、生产级可用的高性能代码高亮库。每个阶段都有明确的验收标准，确保开发质量和进度。

**关键要点：**
- 分阶段开发，每阶段可独立测试
- 遵循 TextMate 规范，无 AST 语义解析
- 开闭原则，扩展无需修改核心代码
- 全平台兼容，API 调用方式一致
- 性能优化，流式解析内存占用低

**相关文档：**
- [产品说明](./01-product-overview.md)
- [技术栈与目录结构](./02-tech-stack.md)
- [核心功能设计](./04-core-features.md)
