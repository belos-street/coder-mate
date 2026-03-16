# 核心功能详细设计与实现要求

## 核心约束

所有功能均遵循「**核心解析逻辑不变，上层能力灵活扩展**」的原则，基于 TextMate 规范实现，无 AST 语义解析，保证学习价值与性能。

---

## 4.1 核心基础功能 - 双核心 API

**必实现，优先级最高**

提供两个全局通用的核心 API，无平台差异，参数规范统一，是所有能力的入口，与原生 Shiki API 风格完全对齐。

### ✅ 1. `codeToTokens(code: string, lang: string = 'javascript'): Promise<Token[]>`

#### 入参
- `code`: 原始代码字符串
- `lang`: 语言标识（默认 JavaScript）

#### 返回值
标准 TextMate Token 数组（Promise 包裹，兼容异步 WASM 加载）

#### 核心要求

**1. Token 为纯数据对象**
- 无样式、无耦合、无冗余字段
- 是所有能力的核心数据载体

**2. Token 标准结构**
与 VSCode/Shiki 完全一致，不可修改：

```typescript
interface Token {
  type: string; // TextMate 语法作用域，如：js.keyword、js.doc.tag、text.plain
  value: string; // 纯文本内容，无任何标签/转义，如："const"、"@param"、" "
}
```

**3. Token 数组顺序**
- 与原始代码完全一致
- 拼接所有 Token.value 可还原原始代码

**4. 空格、换行符处理**
- 统一标记为 `type: 'text.plain'`

#### 使用示例

```typescript
import { codeToTokens } from './dist/ts/index.js';

const code = 'const x = 42;';
const tokens = await codeToTokens(code, 'javascript');

console.log(tokens);
// 输出：
// [
//   { type: 'keyword.control.js', value: 'const' },
//   { type: 'text.plain', value: ' ' },
//   { type: 'variable.other.readwrite.js', value: 'x' },
//   { type: 'text.plain', value: ' ' },
//   { type: 'keyword.operator.js', value: '=' },
//   { type: 'text.plain', value: ' ' },
//   { type: 'constant.numeric.js', value: '42' },
//   { type: 'punctuation.js', value: ';' }
// ]
```

#### 错误处理

```typescript
try {
    const tokens = await codeToTokens(code, 'javascript');
} catch (error) {
    if (error.message.includes('未注册')) {
        console.error('语言未注册:', error.message);
    } else {
        console.error('解析失败:', error.message);
    }
}
```

---

### ✅ 2. `codeToHtml(code: string, lang: string = 'javascript', theme: string = 'dark'): Promise<string>`

#### 入参
- `code`: 原始代码字符串
- `lang`: 语言标识（默认 JavaScript）
- `theme`: 主题名称（默认暗色主题）

#### 返回值
可直接渲染的高亮 HTML 字符串

#### 核心要求

**1. 底层实现**
- 基于 `codeToTokens` 生成的 Token 数组 + 主题样式映射实现
- 无重复解析

**2. HTML 结构**
标准结构：`<pre><code><span class="xxx" style="xxx">文本</span></code></pre>`

**3. 样式**
- 行内样式，开箱即用
- 无需引入额外 CSS 文件

**4. 主题切换**
- 支持动态切换主题
- 切换主题无需重新解析代码，仅需重新映射样式

#### 使用示例

```typescript
import { codeToHtml } from './dist/ts/index.js';

const code = `
function greet(name) {
    console.log(\`Hello, \${name}!\`);
    return true;
}
`.trim();

const html = await codeToHtml(code, 'javascript', 'dark');
document.getElementById('code-container').innerHTML = html;
```

#### HTML 输出示例

```html
<pre class="mini-highlighter" style="background: #161b22; padding: 16px; border-radius: 8px; font-family: Consolas, monospace; font-size: 14px;">
<code>
<span style="color: #ff7b72; font-weight: bold;">function</span>
<span style="color: #c9d1d9;"> </span>
<span style="color: #d2a8ff;">greet</span>
<span style="color: #c9d1d9;">(</span>
<span style="color: #ffa657;">name</span>
<span style="color: #c9d1d9;">) {</span>
<span style="color: #c9d1d9;">    </span>
<span style="color: #d2a8ff;">console</span>
<span style="color: #c9d1d9;">.</span>
<span style="color: #d2a8ff;">log</span>
<span style="color: #c9d1d9;">(</span>
<span style="color: #a5d6ff;">`Hello, ${name}!`</span>
<span style="color: #c9d1d9;">);</span>
<span style="color: #c9d1d9;">    </span>
<span style="color: #ff7b72; font-weight: bold;">return</span>
<span style="color: #c9d1d9;"> </span>
<span style="color: #79c0ff;">true</span>
<span style="color: #c9d1d9;">;</span>
<span style="color: #c9d1d9;">}</span>
</code>
</pre>
```

#### 主题切换示例

```typescript
// 使用暗色主题
const darkHtml = await codeToHtml(code, 'javascript', 'dark');

// 切换到亮色主题（无需重新解析）
const lightHtml = await codeToHtml(code, 'javascript', 'light');

// 使用自定义主题
const customHtml = await codeToHtml(code, 'javascript', 'custom');
```

---

## 4.2 核心扩展功能 - 标准化多语言接口

**必实现，易扩展**

### 设计原则

1. **无入侵式扩展**：新增任意编程语言，**无需修改 Go 核心解析逻辑与 TS 上层核心代码**
2. **规则解耦**：每种语言的语法规则独立维护，遵循 TextMate 状态机规范
3. **按需加载**：支持注册/注销语言，支持多语言共存切换

### 核心结构

```typescript
interface LanguageConfig {
  id: string; // 语言唯一标识，如：javascript、typescript、html
  name: string; // 语言名称，如：JavaScript
  scopeMap: Record<string, string>; // 语法特征 → TextMate语法作用域映射
}
```

### 核心方法

```typescript
// 注册语言（新增语言仅需调用此方法）
function registerLanguage(config: LanguageConfig): void;

// 切换默认语言
function setDefaultLanguage(langId: string): void;

// 获取已注册的语言
function getLanguage(langId: string): LanguageConfig | null;

// 获取所有已注册的语言列表
function getRegisteredLanguages(): string[];
```

### 扩展要求

新增语言仅需 3 步：
1. 新增语言规则文件
2. 编写语法规则映射
3. 调用 `registerLanguage` 注册，即可无缝使用

### 使用示例

#### 注册新语言

```typescript
import { registerLanguage } from './dist/ts/index.js';

// 注册 Python 语言
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
    }
});
```

#### 使用新语言

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

#### 切换默认语言

```typescript
import { setDefaultLanguage } from './dist/ts/index.js';

setDefaultLanguage('python');
```

#### 获取已注册的语言

```typescript
import { getRegisteredLanguages } from './dist/ts/index.js';

const languages = getRegisteredLanguages();
console.log(languages); // ['javascript', 'python', 'typescript']
```

---

## 4.3 核心扩展功能 - 标准化主题接口

**必实现，易扩展，样式解耦**

### 设计原则

1. **完全解耦**：**Token 的语法作用域（type） 与 CSS 样式 彻底分离**，高亮样式不硬编码，全部由主题配置管理
2. **灵活切换**：解析一次代码生成 Token 数组后，可基于不同主题生成不同样式的 HTML，无需重新解析
3. **易扩展**：新增主题仅需编写主题配置文件，调用注册方法即可

### 核心结构

```typescript
interface ThemeConfig {
  id: string; // 主题唯一标识，如：dark、light、custom
  name: string; // 主题名称，如：默认暗色主题
  styles: Record<string, string>; // 语法作用域 → CSS行内样式映射，如："js.keyword": "color: #79c0ff; font-weight: bold;"
}
```

### 核心方法

```typescript
// 注册主题（新增主题仅需调用此方法）
function registerTheme(config: ThemeConfig): void;

// 切换当前主题
function setTheme(themeId: string): void;

// 获取已注册的主题
function getTheme(themeId: string): ThemeConfig | null;

// 获取当前主题
function getCurrentTheme(): string;

// 获取所有已注册的主题列表
function getRegisteredThemes(): string[];
```

### 扩展要求

新增主题仅需 2 步：
1. 新增主题配置文件
2. 调用 `registerTheme` 注册，即可无缝切换

### 使用示例

#### 注册新主题

```typescript
import { registerTheme } from './dist/ts/index.js';

// 注册自定义主题
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
        'variable.other.readwrite.js': 'color: #d2a8ff;',
        'text.plain': 'color: #c9d1d9;',
    }
});
```

#### 使用新主题

```typescript
import { codeToHtml } from './dist/ts/index.js';

const code = 'const x = 42;';
const html = await codeToHtml(code, 'javascript', 'github-dark');
document.getElementById('code-container').innerHTML = html;
```

#### 切换主题

```typescript
import { setTheme, codeToHtml } from './dist/ts/index.js';

const code = 'const x = 42;';

// 使用暗色主题
setTheme('dark');
const darkHtml = await codeToHtml(code, 'javascript');

// 切换到亮色主题（无需重新解析）
setTheme('light');
const lightHtml = await codeToHtml(code, 'javascript');

// 切换到自定义主题
setTheme('github-dark');
const customHtml = await codeToHtml(code, 'javascript');
```

#### 获取已注册的主题

```typescript
import { getRegisteredThemes } from './dist/ts/index.js';

const themes = getRegisteredThemes();
console.log(themes); // ['dark', 'light', 'github-dark']
```

### 主题样式映射示例

#### 暗色主题（Dark Theme）

```typescript
{
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
```

#### 亮色主题（Light Theme）

```typescript
{
    'keyword.control.js': 'color: #cf222e; font-weight: bold;',
    'string.quoted.double.js': 'color: #0a3069;',
    'string.quoted.single.js': 'color: #0a3069;',
    'comment.line.double-slash.js': 'color: #6e7781; font-style: italic;',
    'comment.block.js': 'color: #6e7781; font-style: italic;',
    'comment.block.documentation.js': 'color: #6e7781; font-style: italic;',
    'constant.numeric.js': 'color: #0550ae;',
    'constant.language.js': 'color: #0550ae;',
    'variable.other.readwrite.js': 'color: #953800;',
    'function.js': 'color: #953800;',
    'keyword.operator.js': 'color: #cf222e;',
    'punctuation.js': 'color: #24292f;',
    'text.plain': 'color: #24292f;',
}
```

---

## 4.4 核心扩展功能 - 灵活的代码装饰器能力

**必实现，个性化增强**

### 设计原则

1. **解耦性**：装饰逻辑与解析、渲染逻辑完全解耦，不修改核心 Token 数据与高亮结构
2. **灵活性**：支持全局装饰与精准节点装饰，覆盖所有主流个性化需求
3. **可组合性**：支持注册多个装饰器，按注册顺序执行，支持启用/禁用单个装饰器

### 核心装饰能力（两类全覆盖）

#### ✅ 类型一：全局装饰能力

支持在渲染后的高亮代码块进行全局装饰：

**1. 顶部/底部追加内容**
- 支持在渲染后的高亮代码块「顶部/底部」追加自定义文本/HTML 内容
- 例如：代码标题、复制按钮、说明文字

**2. 外层包裹容器**
- 支持对整个 `<pre>` 代码块外层包裹自定义 class 或自定义容器标签
- 例如：添加自定义样式类、包裹在自定义容器中

**3. 添加自定义属性**
- 支持为整个代码块添加自定义属性
- 例如：data-lang、data-theme

#### ✅ 类型二：精准节点装饰能力

支持基于 **3种匹配条件** 对指定高亮节点进行精准装饰，满足精细化需求：

**1. 按 Token 语法作用域匹配**
- 例如：所有 `js.doc.tag` 类型的节点、所有 `js.keyword` 类型的节点

**2. 按代码位置匹配**
- 例如：第 5 行、第 3-8 行、第 2 行第 5 列到第 10 列

**3. 按文本内容匹配**
- 例如：包含 `username` 的节点、等于 `const` 的节点

**装饰效果支持：**
- 将指定节点包裹在自定义 class 中
- 为节点添加自定义前缀/后缀标签
- 为节点追加自定义属性（如：data-token-type）

### 核心方法

```typescript
// 注册装饰器（新增装饰器仅需调用此方法）
function registerDecorator(decorator: DecoratorConfig): void;

// 启用/禁用装饰器
function toggleDecorator(decoratorId: string, enabled: boolean): void;

// 获取已启用的装饰器
function getEnabledDecorators(): DecoratorConfig[];
```

### 使用示例

#### 全局装饰示例

**1. 添加代码标题**

```typescript
import { registerDecorator } from './dist/ts/index.js';

registerDecorator({
    id: 'code-title',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<div class="code-title">JavaScript 示例代码</div>',
    }
});
```

**2. 添加复制按钮**

```typescript
registerDecorator({
    id: 'copy-button',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<button class="copy-btn" onclick="copyCode()">复制</button>',
    }
});
```

**3. 外层包裹容器**

```typescript
registerDecorator({
    id: 'custom-wrapper',
    enabled: true,
    type: 'global',
    config: {
        wrapClass: 'custom-code-wrapper',
    }
});
```

**4. 添加自定义属性**

```typescript
registerDecorator({
    id: 'custom-attributes',
    enabled: true,
    type: 'global',
    config: {
        attr: {
            'data-lang': 'javascript',
            'data-theme': 'dark',
            'data-version': '1.0.0',
        }
    }
});
```

#### 精准节点装饰示例

**1. 按类型装饰所有关键字**

```typescript
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
```

**2. 按内容装饰特定变量**

```typescript
registerDecorator({
    id: 'highlight-username',
    enabled: true,
    type: 'node',
    config: {
        matchRule: {
            content: 'username',
        },
        wrapClass: 'important-variable',
    }
});
```

**3. 按位置装饰特定行**

```typescript
registerDecorator({
    id: 'highlight-line-5',
    enabled: true,
    type: 'node',
    config: {
        matchRule: {
            lineRange: [5, 5], // 仅第5行
        },
        wrapClass: 'highlighted-line',
    }
});
```

**4. 按位置装饰多行**

```typescript
registerDecorator({
    id: 'highlight-lines-3-8',
    enabled: true,
    type: 'node',
    config: {
        matchRule: {
            lineRange: [3, 8], // 第3行到第8行
        },
        wrapClass: 'highlighted-range',
    }
});
```

#### 启用/禁用装饰器

```typescript
import { toggleDecorator } from './dist/ts/index.js';

// 禁用装饰器
toggleDecorator('copy-button', false);

// 启用装饰器
toggleDecorator('copy-button', true);
```

### 装饰器组合示例

```typescript
import { registerDecorator, codeToHtml } from './dist/ts/index.js';

// 注册多个装饰器
registerDecorator({
    id: 'code-title',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<div class="code-title">JavaScript 示例</div>',
    }
});

registerDecorator({
    id: 'copy-button',
    enabled: true,
    type: 'global',
    config: {
        prependHtml: '<button class="copy-btn">复制</button>',
    }
});

registerDecorator({
    id: 'highlight-keywords',
    enabled: true,
    type: 'node',
    config: {
        matchRule: {
            tokenType: 'keyword.control.js',
        },
        wrapClass: 'highlighted-keyword',
    }
});

// 所有装饰器会按注册顺序依次应用
const html = await codeToHtml(code, 'javascript', 'dark');
```

---

## 总结

本章节详细介绍了核心功能的四个主要部分：

1. **双核心 API**：`codeToTokens` 和 `codeToHtml`，是所有能力的入口
2. **多语言接口**：支持注册新语言，无需修改核心代码
3. **主题接口**：支持注册新主题，样式与 Token 完全解耦
4. **装饰器系统**：支持全局装饰和精准节点装饰，灵活扩展

所有功能都遵循「核心解析逻辑不变，上层能力灵活扩展」的原则，基于 TextMate 规范实现，无 AST 语义解析。

**相关文档：**
- [产品说明](./01-product-overview.md)
- [技术栈与目录结构](./02-tech-stack.md)
- [分阶段实现步骤](./03-implementation-steps.md)
- [核心代码与扩展指南](./05-core-code.md)
