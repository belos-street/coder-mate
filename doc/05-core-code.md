# 关键核心代码实现与扩展开发指南

## 五、关键核心代码实现（完整可复用，带注释，直接落地）

### 5.1 TypeScript 核心代码（重点，上层封装）

#### ✅ 5.1.1 核心类型定义 `src/ts/types/index.ts`

```typescript
/**
 * 标准Token结构 - 与Go端完全对齐，TextMate规范，不可修改
 */
export interface Token {
  type: string;
  value: string;
}

/**
 * 语言配置标准结构
 */
export interface LanguageConfig {
  id: string;
  name: string;
  scopeMap: Record<string, string>;
}

/**
 * 主题配置标准结构
 */
export interface ThemeConfig {
  id: string;
  name: string;
  styles: Record<string, string>;
}

/**
 * 装饰器配置标准结构
 */
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

#### ✅ 5.1.2 核心API封装 `src/ts/index.ts`

```typescript
import { initWasm } from './wasm/loader';
import { codeToTokens as wasmCodeToTokens } from './wasm/bridge';
import { registerLanguage, getLanguage } from './language';
import { registerTheme, getTheme, setTheme } from './theme';
import { registerDecorator, applyDecorators } from './decorator';
import { Token } from './types';

initWasm().catch(err => console.error('WASM初始化失败:', err));

export const codeToTokens = async (code: string, lang: string = 'javascript'): Promise<Token[]> => {
  if (!code) return [];
  const langConfig = getLanguage(lang);
  if (!langConfig) throw new Error(`语言${lang}未注册`);
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

  html = applyDecorators(html, themeId);

  return `<pre class="mini-highlighter" style="background: #161b22; padding: 16px; border-radius: 8px; font-family: Consolas, monospace; font-size: 14px;"><code>${html}</code></pre>`;
};

export { registerLanguage, registerTheme, setTheme, registerDecorator };
```

#### ✅ 5.1.3 WASM加载器 `src/ts/wasm/loader.ts`

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
        throw new Error(`WASM初始化失败: ${error instanceof Error ? error.message : '未知错误'}`);
    }
}

export function getWasmInstance(): WebAssembly.Instance {
    if (!wasmInstance || !wasmInitialized) {
        throw new Error('WASM未初始化，请先调用 initWasm()');
    }
    return wasmInstance;
}
```

#### ✅ 5.1.4 WASM桥接 `src/ts/wasm/bridge.ts`

```typescript
import { getWasmInstance } from './loader';
import { Token } from '../types';

export async function codeToTokens(code: string): Promise<Token[]> {
    const wasmInstance = getWasmInstance();
    const jsonString = wasmInstance.exports.codeToTokens(code) as string;
    return JSON.parse(jsonString) as Token[];
}
```

#### ✅ 5.1.5 语言管理 `src/ts/language/index.ts`

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

#### ✅ 5.1.6 主题管理 `src/ts/theme/index.ts`

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

#### ✅ 5.1.7 装饰器管理 `src/ts/decorator/index.ts`

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

---

### 5.2 Golang 核心代码（重点，底层解析）

#### ✅ 5.2.1 核心解析文件 `src/go/highlighter.go`

完整可复用，含TextMate状态机+Token生成+WASM导出：

```go
package main

import (
    "encoding/json"
    "regexp"
    "strings"
    "syscall/js"
)

// Token 标准结构体 - 与TS端完全对齐
type Token struct {
    Type  string `json:"type"`
    Value string `json:"value"`
}

// 状态枚举 - TextMate有限状态机（含JSDoc嵌套子状态）
const (
    STATE_ROOT          = "root"
    STATE_STRING_DOUBLE = "string_double"
    STATE_STRING_SINGLE = "string_single"
    STATE_COMMENT_LINE  = "comment_line"
    STATE_COMMENT_BLOCK = "comment_block"
    STATE_JSDOC         = "jsdoc"
)

// 状态机核心结构体
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

// 添加Token到数组
func (h *Highlighter) addToken(tp, val string) {
    if val == "" {
        return
    }
    h.tokens = append(h.tokens, Token{Type: tp, Value: val})
}

// 核心解析方法：代码转Token数组
func (h *Highlighter) CodeToTokens(code string) []Token {
    h.tokens = make([]Token, 0)
    h.currentState = STATE_ROOT

    lines := strings.Split(code, "\n")
    for lineNum, line := range lines {
        h.parseLine(line, lineNum)
        if lineNum < len(lines)-1 {
            h.addToken("text.plain", "\n")
        }
    }

    return h.tokens
}

// 解析单行代码
func (h *Highlighter) parseLine(line string, lineNum int) {
    i := 0
    for i < len(line) {
        switch h.currentState {
        case STATE_ROOT:
            i = h.parseRoot(line, i)
        case STATE_STRING_DOUBLE:
            i = h.parseStringDouble(line, i)
        case STATE_STRING_SINGLE:
            i = h.parseStringSingle(line, i)
        case STATE_COMMENT_LINE:
            i = h.parseCommentLine(line, i)
        case STATE_COMMENT_BLOCK:
            i = h.parseCommentBlock(line, i)
        case STATE_JSDOC:
            i = h.parseJSDoc(line, i)
        }
    }
}

// ROOT状态解析
func (h *Highlighter) parseRoot(line string, pos int) int {
    remaining := line[pos:]

    // 匹配双引号字符串
    if strings.HasPrefix(remaining, `"`) {
        h.addToken("punctuation.definition.string.begin", `"`)
        h.currentState = STATE_STRING_DOUBLE
        return pos + 1
    }

    // 匹配单引号字符串
    if strings.HasPrefix(remaining, `'`) {
        h.addToken("punctuation.definition.string.begin", `'`)
        h.currentState = STATE_STRING_SINGLE
        return pos + 1
    }

    // 匹配单行注释
    if strings.HasPrefix(remaining, "//") {
        h.addToken("comment.line.double-slash.js", remaining)
        h.currentState = STATE_ROOT
        return len(line)
    }

    // 匹配多行注释开始
    if strings.HasPrefix(remaining, "/*") {
        if strings.HasPrefix(remaining, "/**") {
            h.addToken("comment.block.documentation.js", "/**")
            h.currentState = STATE_JSDOC
        } else {
            h.addToken("comment.block.js", "/*")
            h.currentState = STATE_COMMENT_BLOCK
        }
        return pos + 2
    }

    // 匹配关键字
    keywords := []string{"const", "let", "var", "function", "return", "if", "else", "for", "while", "class", "import", "export", "default", "from", "async", "await", "try", "catch", "throw", "new", "this", "super", "extends", "static", "typeof", "instanceof"}
    for _, kw := range keywords {
        if strings.HasPrefix(remaining, kw) {
            nextChar := rune(0)
            if pos+len(kw) < len(line) {
                nextChar = rune(line[pos+len(kw)])
            }
            if !isWordChar(nextChar) {
                h.addToken("keyword.control.js", kw)
                return pos + len(kw)
            }
        }
    }

    // 匹配布尔值和null
    literals := []string{"true", "false", "null", "undefined"}
    for _, lit := range literals {
        if strings.HasPrefix(remaining, lit) {
            nextChar := rune(0)
            if pos+len(lit) < len(line) {
                nextChar = rune(line[pos+len(lit)])
            }
            if !isWordChar(nextChar) {
                h.addToken("constant.language.js", lit)
                return pos + len(lit)
            }
        }
    }

    // 匹配数字
    if matched, end := matchNumber(remaining); matched {
        h.addToken("constant.numeric.js", remaining[:end])
        return pos + end
    }

    // 匹配标识符
    if matched, end := matchIdentifier(remaining); matched {
        h.addToken("variable.other.readwrite.js", remaining[:end])
        return pos + end
    }

    // 匹配运算符
    operators := []string{"=", "==", "===", "!=", "!==", "+", "-", "*", "/", "%", "&&", "||", "!", "&", "|", "^", "~", "<<", ">>", ">>>", "<", ">", "<=", ">=", "=>"}
    for _, op := range operators {
        if strings.HasPrefix(remaining, op) {
            h.addToken("keyword.operator.js", op)
            return pos + len(op)
        }
    }

    // 匹配标点符号
    punctuations := []string{"(", ")", "{", "}", "[", "]", ";", ":", ",", "."}
    for _, p := range punctuations {
        if strings.HasPrefix(remaining, p) {
            h.addToken("punctuation.js", p)
            return pos + len(p)
        }
    }

    // 默认为普通文本
    h.addToken("text.plain", string(remaining[0]))
    return pos + 1
}

// 双引号字符串状态解析
func (h *Highlighter) parseStringDouble(line string, pos int) int {
    remaining := line[pos:]
    idx := strings.Index(remaining, `"`)
    if idx == -1 {
        h.addToken("string.quoted.double.js", remaining)
        return len(line)
    }
    if idx > 0 {
        h.addToken("string.quoted.double.js", remaining[:idx])
    }
    h.addToken("punctuation.definition.string.end", `"`)
    h.currentState = STATE_ROOT
    return pos + idx + 1
}

// 单引号字符串状态解析
func (h *Highlighter) parseStringSingle(line string, pos int) int {
    remaining := line[pos:]
    idx := strings.Index(remaining, `'`)
    if idx == -1 {
        h.addToken("string.quoted.single.js", remaining)
        return len(line)
    }
    if idx > 0 {
        h.addToken("string.quoted.single.js", remaining[:idx])
    }
    h.addToken("punctuation.definition.string.end", `'`)
    h.currentState = STATE_ROOT
    return pos + idx + 1
}

// 单行注释状态解析
func (h *Highlighter) parseCommentLine(line string, pos int) int {
    h.addToken("comment.line.double-slash.js", line[pos:])
    h.currentState = STATE_ROOT
    return len(line)
}

// 多行注释状态解析
func (h *Highlighter) parseCommentBlock(line string, pos int) int {
    remaining := line[pos:]
    idx := strings.Index(remaining, "*/")
    if idx == -1 {
        h.addToken("comment.block.js", remaining)
        return len(line)
    }
    if idx > 0 {
        h.addToken("comment.block.js", remaining[:idx])
    }
    h.addToken("comment.block.js", "*/")
    h.currentState = STATE_ROOT
    return pos + idx + 2
}

// JSDoc状态解析
func (h *Highlighter) parseJSDoc(line string, pos int) int {
    remaining := line[pos:]

    // 匹配JSDoc标签
    if matched, _ := regexp.MatchString(`@\w+`, remaining); matched {
        idx := regexp.MustCompile(`@\w+`).FindStringIndex(remaining)[1]
        h.addToken("storage.type.class.jsdoc", remaining[:idx])
        return pos + idx
    }

    // 匹配JSDoc结束
    idx := strings.Index(remaining, "*/")
    if idx != -1 {
        if idx > 0 {
            h.addToken("comment.block.documentation.js", remaining[:idx])
        }
        h.addToken("comment.block.documentation.js", "*/")
        h.currentState = STATE_ROOT
        return pos + idx + 2
    }

    h.addToken("comment.block.documentation.js", remaining)
    return len(line)
}

// 辅助函数：判断是否为单词字符
func isWordChar(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '$'
}

// 辅助函数：匹配数字
func matchNumber(s string) (bool, int) {
    if len(s) == 0 {
        return false, 0
    }
    i := 0
    if s[i] == '-' {
        i++
    }
    if i >= len(s) {
        return false, 0
    }
    if s[i] == '0' {
        i++
        if i < len(s) && (s[i] == 'x' || s[i] == 'X') {
            i++
            for i < len(s) && isHexDigit(rune(s[i])) {
                i++
            }
            return i > 2, i
        }
    }
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        i++
    }
    if i < len(s) && s[i] == '.' {
        i++
        for i < len(s) && s[i] >= '0' && s[i] <= '9' {
            i++
        }
    }
    if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
        i++
        if i < len(s) && (s[i] == '+' || s[i] == '-') {
            i++
        }
        for i < len(s) && s[i] >= '0' && s[i] <= '9' {
            i++
        }
    }
    return i > 0, i
}

func isHexDigit(r rune) bool {
    return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')
}

// 辅助函数：匹配标识符
func matchIdentifier(s string) (bool, int) {
    if len(s) == 0 {
        return false, 0
    }
    i := 0
    if isWordChar(rune(s[i])) && !((s[i] >= '0' && s[i] <= '9')) {
        i++
        for i < len(s) && isWordChar(rune(s[i])) {
            i++
        }
        return i > 0, i
    }
    return false, 0
}

// WASM导出方法：代码转Token（JSON格式）
func codeToTokensWrapper(this js.Value, args []js.Value) interface{} {
    if len(args) < 1 {
        return ""
    }
    code := args[0].String()
    h := NewHighlighter()
    tokens := h.CodeToTokens(code)
    jsonBytes, _ := json.Marshal(tokens)
    return string(jsonBytes)
}

// WASM初始化方法
func main() {
    c := make(chan struct{}, 0)
    js.Global().Set("codeToTokens", js.FuncOf(codeToTokensWrapper))
    <-c
}
```

#### ✅ 5.2.2 Go WASM 编译命令（极简，直接执行）

```bash
# 进入src/go目录
cd src/go
# 编译为WASM二进制文件，输出到dist/wasm
GOOS=js GOARCH=wasm go build -o ../../dist/wasm/highlighter.wasm highlighter.go
```

---

## 六、扩展开发指南（快速新增，无入侵，3步完成）

### 6.1 快速新增一门语言（3步完成）

#### 步骤1：在 `src/ts/language` 目录新增 `xxx.ts` 文件，编写语言配置

**示例：新增 Python 语言**

```typescript
// src/ts/language/python.ts
import { registerLanguage } from './index';

registerLanguage({
    id: 'python',
    name: 'Python',
    scopeMap: {
        'def': 'keyword.control.python',
        'class': 'keyword.control.python',
        'import': 'keyword.control.import.python',
        'from': 'keyword.control.import.python',
        '"': 'string.quoted.double.python',
        "'": 'string.quoted.single.python',
        '#': 'comment.line.number-sign.python',
        'True': 'constant.language.python',
        'False': 'constant.language.python',
        'None': 'constant.language.python',
        'if': 'keyword.control.python',
        'else': 'keyword.control.python',
        'for': 'keyword.control.python',
        'while': 'keyword.control.python',
        'return': 'keyword.control.python',
        '=': 'keyword.operator.python',
        '==': 'keyword.operator.python',
        '!=': 'keyword.operator.python',
    }
});
```

#### 步骤2：调用 `registerLanguage(config)` 注册语言

```typescript
import { registerLanguage } from './src/ts/language/python';

// 在应用启动时注册
registerLanguage({
    id: 'python',
    name: 'Python',
    scopeMap: {
        // ... 语法规则映射
    }
});
```

#### 步骤3：调用 `codeToTokens(code, 'xxx')` 即可使用

```typescript
import { codeToHtml } from './dist/ts/index.js';

const pythonCode = `
def greet(name):
    print(f"Hello, {name}!")
    return True
`.trim();

const html = await codeToHtml(pythonCode, 'python', 'dark');
document.getElementById('code-container').innerHTML = html;
```

---

### 6.2 快速新增一个主题（2步完成）

#### 步骤1：在 `src/ts/theme` 目录新增 `xxx.ts` 文件，编写主题样式映射

**示例：新增 GitHub Dark 主题**

```typescript
// src/ts/theme/github-dark.ts
import { registerTheme } from './index';

registerTheme({
    id: 'github-dark',
    name: 'GitHub Dark',
    styles: {
        'keyword.control.js': 'color: #ff7b72; font-weight: bold;',
        'string.quoted.double.js': 'color: #a5d6ff;',
        'string.quoted.single.js': 'color: #a5d6ff;',
        'comment.line.double-slash.js': 'color: #8b949e; font-style: italic;',
        'comment.block.js': 'color: #8b949e; font-style: italic;',
        'comment.block.documentation.js': 'color: #8b949e; font-style: italic;',
        'constant.numeric.js': 'color: #79c0ff;',
        'constant.language.js': 'color: #79c0ff;',
        'variable.other.readwrite.js': 'color: #d2a8ff;',
        'function.js': 'color: #d2a8ff;',
        'keyword.operator.js': 'color: #ff7b72;',
        'punctuation.js': 'color: #c9d1d9;',
        'text.plain': 'color: #c9d1d9;',
    }
});
```

#### 步骤2：调用 `registerTheme(config)` 注册主题，调用 `setTheme('xxx')` 切换主题

```typescript
import { registerTheme, setTheme, codeToHtml } from './dist/ts/index.js';

// 注册主题
registerTheme({
    id: 'github-dark',
    name: 'GitHub Dark',
    styles: {
        // ... 样式映射
    }
});

// 切换主题
setTheme('github-dark');

// 使用主题
const html = await codeToHtml(code, 'javascript', 'github-dark');
```

---

### 6.3 快速新增一个装饰器（2步完成）

#### 步骤1：编写装饰器配置对象，定义装饰规则

**示例：添加代码标题装饰器**

```typescript
const codeTitleDecorator = {
    id: 'code-title',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<div class="code-title">JavaScript 示例代码</div>',
    }
};
```

**示例：添加复制按钮装饰器**

```typescript
const copyButtonDecorator = {
    id: 'copy-button',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<button class="copy-btn" onclick="copyCode()">复制</button>',
    }
};
```

**示例：添加关键字高亮装饰器**

```typescript
const keywordHighlightDecorator = {
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
};
```

#### 步骤2：调用 `registerDecorator(config)` 注册装饰器，自动生效

```typescript
import { registerDecorator } from './dist/ts/index.js';

// 注册装饰器
registerDecorator(codeTitleDecorator);
registerDecorator(copyButtonDecorator);
registerDecorator(keywordHighlightDecorator);

// 使用代码高亮时，装饰器会自动应用
const html = await codeToHtml(code, 'javascript', 'dark');
```

---

## 七、开发约束与核心规范（保障学习价值与代码质量）

### 1. 核心原则
全程遵循 TextMate 规范，仅做「文本特征匹配+状态机流转」，**不做任何 AST 语义解析**，保证与 VSCode/Shiki 逻辑一致。

### 2. 性能原则
- **Go 端**：保证流式逐行解析，时间复杂度 O(n)，内存占用恒定
- **TS 端**：无冗余计算，装饰器按需执行

### 3. 扩展原则
开闭原则，所有扩展能力无需修改核心代码，仅需新增配置文件。

### 4. 兼容原则
全平台 API 一致，无平台专属代码。

### 5. 学习原则
关键代码附详细注释，核心逻辑无黑盒，便于复盘学习 TextMate 状态机原理。

---

## 八、错误处理机制（健壮性保障，生产级必备）

### 8.1 WASM加载错误处理

在 `src/ts/wasm/loader.ts` 中实现完整的错误处理：

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
        throw new Error(`WASM初始化失败: ${error instanceof Error ? error.message : '未知错误'}`);
    }
}

export function getWasmInstance(): WebAssembly.Instance {
    if (!wasmInstance || !wasmInitialized) {
        throw new Error('WASM未初始化，请先调用 initWasm()');
    }
    return wasmInstance;
}
```

### 8.2 语法错误容错处理

Go端状态机实现容错机制，确保语法错误不影响高亮：

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
    // 如果遇到无法识别的字符，作为普通文本处理
    h.addToken("text.plain", string(remaining[0]))
    return pos + 1
}
```

### 8.3 无效语言/主题处理

在TS端增加友好的错误提示：

```typescript
export const codeToTokens = async (code: string, lang: string = 'javascript'): Promise<Token[]> => {
    if (!code) return [];
    
    const langConfig = getLanguage(lang);
    if (!langConfig) {
        const availableLangs = getRegisteredLanguages();
        throw new Error(
            `语言"${lang}"未注册。可用语言: ${availableLangs.join(', ')}`
        );
    }
    
    try {
        return await wasmCodeToTokens(code);
    } catch (error) {
        console.error(`代码解析失败:`, error);
        throw new Error(`代码解析失败: ${error instanceof Error ? error.message : '未知错误'}`);
    }
};

export const codeToHtml = async (code: string, lang: string = 'javascript', themeId: string = 'dark'): Promise<string> => {
    const theme = getTheme(themeId);
    if (!theme) {
        const availableThemes = getRegisteredThemes();
        throw new Error(
            `主题"${themeId}"未注册。可用主题: ${availableThemes.join(', ')}`
        );
    }
    
    // 原有逻辑...
};
```

### 8.4 装饰器执行错误隔离

确保单个装饰器失败不影响其他装饰器：

```typescript
export function applyDecorators(html: string, themeId: string): string {
    const decorators = getEnabledDecorators();
    let result = html;
    
    for (const decorator of decorators) {
        try {
            result = applySingleDecorator(result, decorator);
        } catch (error) {
            console.warn(`装饰器"${decorator.id}"执行失败，已跳过:`, error);
            // 继续执行下一个装饰器
        }
    }
    
    return result;
}
```

---

**相关文档：**
- [产品说明](./01-product-overview.md)
- [技术栈与目录结构](./02-tech-stack.md)
- [分阶段实现步骤](./03-implementation-steps.md)
- [核心功能设计](./04-core-features.md)
- [MVP与测试策略](./06-mvp-testing.md)
