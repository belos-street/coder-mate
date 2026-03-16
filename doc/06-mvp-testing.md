# MVP定义与测试策略

## 九、MVP定义（最小可行产品，快速验证）

### 9.1 MVP范围

为确保快速验证核心假设，MVP阶段仅实现以下功能：

#### ✅ 核心功能（必做）

1. **JavaScript语言支持**：完整的JS语法高亮（关键字、字符串、注释、数字、运算符）
   - 关键字：const、let、var、function、return、if、else、for、while、class、import、export等
   - 字符串：双引号字符串、单引号字符串、模板字符串
   - 注释：单行注释（//）、多行注释（/* */）、JSDoc注释（/** */）
   - 数字：整数、浮点数、十六进制、科学计数法
   - 运算符：=、==、===、!=、!==、+、-、*、/、%、&&、||、!等
   - 标点符号：(、)、{、}、[、]、;、:、,、.

2. **基础主题**：提供dark和light两个默认主题
   - Dark主题：暗色背景，适合夜间使用
   - Light主题：亮色背景，适合日间使用
   - 包含常用语法作用域的样式映射

3. **核心API**：`codeToTokens` 和 `codeToHtml` 两个基础API
   - `codeToTokens`：返回标准Token数组
   - `codeToHtml`：返回可直接渲染的HTML字符串

4. **WASM桥接**：完整的Go-WASM通信机制
   - WASM初始化和加载
   - Go方法调用桥接
   - JSON序列化/反序列化

#### ✅ 基础扩展（推荐做）

1. **装饰器系统**：支持全局装饰（代码块标题、复制按钮）
   - 全局装饰：代码块前后追加内容、外层包裹自定义容器
   - 装饰器注册、启用/禁用

2. **多语言接口**：语言注册/切换/获取的完整接口
   - `registerLanguage`：注册新语言
   - `getLanguage`：获取已注册的语言
   - `setDefaultLanguage`：切换默认语言
   - `getRegisteredLanguages`：获取所有已注册的语言列表

3. **主题接口**：主题注册/切换/获取的完整接口
   - `registerTheme`：注册新主题
   - `getTheme`：获取已注册的主题
   - `setTheme`：切换当前主题
   - `getCurrentTheme`：获取当前主题
   - `getRegisteredThemes`：获取所有已注册的主题列表

#### ❌ 暂不实现（后续迭代）

1. 其他编程语言（TypeScript、Python、HTML等）
2. 复杂装饰器（精准节点装饰、行号显示等）
3. 性能优化（TinyGo编译、代码压缩）
4. 高级主题（自定义主题编辑器）

### 9.2 MVP验收标准

- [ ] 能够正确高亮JavaScript代码（关键字、字符串、注释、数字）
- [ ] 支持dark/light主题切换
- [ ] 能够在浏览器、Node.js、Bun三个平台运行
- [ ] API调用方式与文档描述一致
- [ ] 处理空代码、语法错误代码时不崩溃
- [ ] 解析1000行代码耗时<100ms

### 9.3 MVP开发时间估算

| 阶段 | 任务 | 预计时间 |
|------|------|----------|
| 阶段1 | 环境初始化与基础结构搭建 | 1-2小时 |
| 阶段2 | 核心基础层开发 | 4-6小时 |
| 阶段3 | 核心API封装 | 2-3小时 |
| 阶段4 | 扩展能力开发 - 多语言+主题接口 | 2-3小时 |
| 阶段5 | 扩展能力开发 - 代码装饰器核心能力 | 2-3小时 |
| 阶段6 | 全平台兼容性测试与性能优化 | 2-3小时 |
| **总计** | | **13-20小时** |

### 9.4 MVP成功标准

**功能完整性：**
- JavaScript语言高亮覆盖率 > 90%
- 主题切换功能正常
- 装饰器系统可用

**性能指标：**
- 解析100行代码 < 10ms
- 解析1000行代码 < 100ms
- WASM文件大小 < 500KB

**稳定性：**
- 空代码不崩溃
- 语法错误代码不崩溃
- 内存占用稳定

**兼容性：**
- 浏览器（Chrome/Firefox/Safari/Edge）测试通过
- Node.js v16+ 测试通过
- Bun 最新稳定版测试通过

---

## 十、测试策略（质量保障，持续迭代）

### 10.1 单元测试

使用Jest测试框架，覆盖核心逻辑：

#### 测试文件结构

```
tests/
├── unit/
│   ├── token.test.ts          # Token生成测试
│   ├── language.test.ts       # 语言管理测试
│   ├── theme.test.ts          # 主题管理测试
│   ├── decorator.test.ts       # 装饰器测试
│   └── wasm.test.ts          # WASM桥接测试
├── integration/
│   ├── api.test.ts           # API集成测试
│   └── platform.test.ts      # 跨平台测试
└── benchmark/
    └── performance.test.ts   # 性能测试
```

#### Token生成测试

```typescript
// tests/unit/token.test.ts
import { codeToTokens } from '../../src/ts';

describe('codeToTokens', () => {
    test('正确解析关键字', async () => {
        const tokens = await codeToTokens('const x = 1;');
        expect(tokens).toContainEqual({ type: 'keyword.control.js', value: 'const' });
    });

    test('正确解析字符串', async () => {
        const tokens = await codeToTokens('const s = "hello";');
        expect(tokens).toContainEqual({ type: 'string.quoted.double.js', value: 'hello' });
    });

    test('正确解析注释', async () => {
        const tokens = await codeToTokens('// comment');
        expect(tokens).toContainEqual({ type: 'comment.line.double-slash.js', value: '// comment' });
    });

    test('正确解析数字', async () => {
        const tokens = await codeToTokens('const x = 42;');
        expect(tokens).toContainEqual({ type: 'constant.numeric.js', value: '42' });
    });

    test('正确解析布尔值', async () => {
        const tokens = await codeToTokens('const x = true;');
        expect(tokens).toContainEqual({ type: 'constant.language.js', value: 'true' });
    });

    test('空代码返回空数组', async () => {
        const tokens = await codeToTokens('');
        expect(tokens).toEqual([]);
    });

    test('语法错误不崩溃', async () => {
        const tokens = await codeToTokens('const x = ;');
        expect(tokens).toBeDefined();
        expect(tokens.length).toBeGreaterThan(0);
    });

    test('Token顺序与代码一致', async () => {
        const code = 'const x = 1;';
        const tokens = await codeToTokens(code);
        const reconstructed = tokens.map(t => t.value).join('');
        expect(reconstructed).toBe(code);
    });
});
```

#### 语言管理测试

```typescript
// tests/unit/language.test.ts
import { registerLanguage, getLanguage, getRegisteredLanguages } from '../../src/ts';

describe('语言管理', () => {
    beforeEach(() => {
        // 清理已注册的语言
        // ...
    });

    test('注册新语言', () => {
        registerLanguage({
            id: 'test-lang',
            name: 'Test Language',
            scopeMap: {}
        });

        const lang = getLanguage('test-lang');
        expect(lang).toBeDefined();
        expect(lang?.id).toBe('test-lang');
    });

    test('获取不存在的语言返回null', () => {
        const lang = getLanguage('non-existent');
        expect(lang).toBeNull();
    });

    test('获取所有已注册的语言', () => {
        registerLanguage({ id: 'lang1', name: 'Lang1', scopeMap: {} });
        registerLanguage({ id: 'lang2', name: 'Lang2', scopeMap: {} });

        const languages = getRegisteredLanguages();
        expect(languages).toContain('lang1');
        expect(languages).toContain('lang2');
    });
});
```

#### 主题管理测试

```typescript
// tests/unit/theme.test.ts
import { registerTheme, getTheme, setTheme, getRegisteredThemes } from '../../src/ts';

describe('主题管理', () => {
    beforeEach(() => {
        // 清理已注册的主题
        // ...
    });

    test('注册新主题', () => {
        registerTheme({
            id: 'test-theme',
            name: 'Test Theme',
            styles: {}
        });

        const theme = getTheme('test-theme');
        expect(theme).toBeDefined();
        expect(theme?.id).toBe('test-theme');
    });

    test('切换主题', () => {
        registerTheme({ id: 'theme1', name: 'Theme1', styles: {} });
        registerTheme({ id: 'theme2', name: 'Theme2', styles: {} });

        setTheme('theme1');
        expect(getCurrentTheme()).toBe('theme1');

        setTheme('theme2');
        expect(getCurrentTheme()).toBe('theme2');
    });

    test('获取不存在的主题返回null', () => {
        const theme = getTheme('non-existent');
        expect(theme).toBeNull();
    });
});
```

#### 装饰器测试

```typescript
// tests/unit/decorator.test.ts
import { registerDecorator, toggleDecorator, getEnabledDecorators } from '../../src/ts';

describe('装饰器', () => {
    beforeEach(() => {
        // 清理已注册的装饰器
        // ...
    });

    test('注册新装饰器', () => {
        const decorator = {
            id: 'test-decorator',
            enabled: true,
            type: 'global' as const,
            config: {}
        };

        registerDecorator(decorator);
        const enabledDecorators = getEnabledDecorators();
        expect(enabledDecorators).toContainEqual(decorator);
    });

    test('启用/禁用装饰器', () => {
        const decorator = {
            id: 'test-decorator',
            enabled: true,
            type: 'global' as const,
            config: {}
        };

        registerDecorator(decorator);
        toggleDecorator('test-decorator', false);

        const enabledDecorators = getEnabledDecorators();
        expect(enabledDecorators).not.toContainEqual(decorator);

        toggleDecorator('test-decorator', true);
        const enabledDecorators2 = getEnabledDecorators();
        expect(enabledDecorators2).toContainEqual(decorator);
    });
});
```

---

### 10.2 集成测试

测试端到端API调用：

#### API集成测试

```typescript
// tests/integration/api.test.ts
import { codeToHtml, registerTheme } from '../../src/ts';

describe('codeToHtml 集成测试', () => {
    test('生成完整HTML', async () => {
        const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
        expect(html).toContain('<pre');
        expect(html).toContain('<code');
        expect(html).toContain('</code>');
        expect(html).toContain('</pre>');
    });

    test('HTML包含正确的样式', async () => {
        const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
        expect(html).toContain('style=');
        expect(html).toContain('color:');
    });

    test('主题切换生效', async () => {
        const darkHtml = await codeToHtml('const x = 1;', 'javascript', 'dark');
        const lightHtml = await codeToHtml('const x = 1;', 'javascript', 'light');
        expect(darkHtml).not.toEqual(lightHtml);
    });

    test('自定义主题生效', async () => {
        registerTheme({
            id: 'custom',
            name: 'Custom Theme',
            styles: {
                'keyword.control.js': 'color: #ff0000;',
            }
        });

        const html = await codeToHtml('const x = 1;', 'javascript', 'custom');
        expect(html).toContain('#ff0000');
    });
});
```

#### 跨平台兼容性测试

```typescript
// tests/integration/platform.test.ts
import { codeToHtml } from '../../src/ts';

describe('跨平台兼容性', () => {
    test('浏览器环境', async () => {
        if (typeof window !== 'undefined') {
            const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
            expect(html).toBeDefined();
            expect(html).toContain('<pre');
        }
    });

    test('Node.js环境', async () => {
        if (typeof process !== 'undefined' && typeof window === 'undefined') {
            const html = await codeToHtml('const x = 1;', 'javascript', 'dark');
            expect(html).toBeDefined();
            expect(html).toContain('<pre');
        }
    });
});
```

---

### 10.3 性能基准测试

使用benchmark.js测试解析性能：

#### 性能测试代码

```typescript
// tests/benchmark/performance.test.ts
import { codeToTokens } from '../../src/ts';
import Benchmark from 'benchmark';

const suite = new Benchmark.Suite();

const sampleCode = `
const x = 1;
const y = 2;
function add(a, b) {
    return a + b;
}
`.repeat(100); // 重复100次，约400行代码

suite
    .add('codeToTokens 400行代码', async () => {
        await codeToTokens(sampleCode);
    })
    .add('codeToTokens 1000行代码', async () => {
        const largeCode = sampleCode.repeat(2.5);
        await codeToTokens(largeCode);
    })
    .on('cycle', (event) => {
        console.log(String(event.target));
    })
    .on('complete', function() {
        console.log('Fastest is ' + this.filter('fastest').map('name'));
    })
    .run();
```

#### 性能指标

| 代码行数 | 目标耗时 | 实际耗时 | 状态 |
|----------|----------|----------|------|
| 100行 | <10ms | - | 待测试 |
| 400行 | <40ms | - | 待测试 |
| 1000行 | <100ms | - | 待测试 |
| 10000行 | <1000ms | - | 待测试 |

---

### 10.4 跨平台兼容性测试

在三个平台分别运行测试：

#### 浏览器测试

```bash
# 浏览器测试
bun run test:browser
```

**package.json 配置：**
```json
{
  "scripts": {
    "test:browser": "jest --config jest.browser.config.js"
  }
}
```

**jest.browser.config.js：**
```javascript
module.exports = {
    testEnvironment: 'jsdom',
    testMatch: ['**/tests/**/*.test.ts'],
    preset: 'ts-jest',
};
```

#### Node.js测试

```bash
# Node.js测试
bun run test:node
```

**package.json 配置：**
```json
{
  "scripts": {
    "test:node": "jest --config jest.node.config.js"
  }
}
```

**jest.node.config.js：**
```javascript
module.exports = {
    testEnvironment: 'node',
    testMatch: ['**/tests/**/*.test.ts'],
    preset: 'ts-jest',
};
```

#### Bun测试

```bash
# Bun测试
bun run test:bun
```

**package.json 配置：**
```json
{
  "scripts": {
    "test:bun": "bun test"
  }
}
```

---

### 10.5 测试覆盖率

使用 Jest 的覆盖率功能：

```bash
# 生成覆盖率报告
bun run test:coverage
```

**package.json 配置：**
```json
{
  "scripts": {
    "test:coverage": "jest --coverage"
  }
}
```

**覆盖率目标：**

| 指标 | 目标值 | 说明 |
|------|--------|------|
| 语句覆盖率 | >80% | 所有语句被执行的比例 |
| 分支覆盖率 | >70% | 所有分支被执行的比例 |
| 函数覆盖率 | >80% | 所有函数被调用的比例 |
| 行覆盖率 | >80% | 所有代码行被执行的比例 |

---

### 10.6 持续集成（CI）

使用 GitHub Actions 进行持续集成：

#### .github/workflows/test.yml

```yaml
name: Test

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    strategy:
      matrix:
        node-version: [16.x, 18.x, 20.x]
    
    steps:
    - uses: actions/checkout@v2
    
    - name: Use Node.js ${{ matrix.node-version }}
      uses: actions/setup-node@v2
      with:
        node-version: ${{ matrix.node-version }}
    
    - name: Install dependencies
      run: bun install
    
    - name: Build
      run: bun run build
    
    - name: Run tests
      run: bun test
    
    - name: Generate coverage
      run: bun run test:coverage
    
    - name: Upload coverage
      uses: codecov/codecov-action@v2
```

---

### 10.7 测试最佳实践

#### 1. 测试命名规范

```typescript
// ✓ 好的命名
test('正确解析关键字', async () => {});
test('空代码返回空数组', async () => {});

// ✗ 不好的命名
test('test1', async () => {});
test('test keyword', async () => {});
```

#### 2. 测试独立性

```typescript
// ✓ 好的做法：每个测试独立运行
describe('Token生成', () => {
    test('测试1', async () => {});
    test('测试2', async () => {});
});

// ✗ 不好的做法：测试之间有依赖
let tokens;
beforeAll(async () => {
    tokens = await codeToTokens('const x = 1;');
});
test('测试1', () => {
    expect(tokens).toBeDefined();
});
```

#### 3. 测试覆盖率

```typescript
// ✓ 好的做法：覆盖正常和异常情况
test('正确解析关键字', async () => {
    const tokens = await codeToTokens('const x = 1;');
    expect(tokens).toContainEqual({ type: 'keyword.control.js', value: 'const' });
});

test('语法错误不崩溃', async () => {
    const tokens = await codeToTokens('const x = ;');
    expect(tokens).toBeDefined();
});

// ✗ 不好的做法：只覆盖正常情况
test('正确解析关键字', async () => {
    const tokens = await codeToTokens('const x = 1;');
    expect(tokens).toContainEqual({ type: 'keyword.control.js', value: 'const' });
});
```

#### 4. 测试可维护性

```typescript
// ✓ 好的做法：使用辅助函数
function expectToken(tokens: Token[], type: string, value: string) {
    expect(tokens).toContainEqual({ type, value });
}

test('正确解析多个关键字', async () => {
    const tokens = await codeToTokens('const x = 1; let y = 2;');
    expectToken(tokens, 'keyword.control.js', 'const');
    expectToken(tokens, 'keyword.control.js', 'let');
});

// ✗ 不好的做法：重复代码
test('正确解析多个关键字', async () => {
    const tokens = await codeToTokens('const x = 1; let y = 2;');
    expect(tokens).toContainEqual({ type: 'keyword.control.js', value: 'const' });
    expect(tokens).toContainEqual({ type: 'keyword.control.js', value: 'let' });
});
```

---

## 总结

本章节详细介绍了MVP定义和测试策略：

### MVP定义
- 明确了MVP的范围和验收标准
- 提供了开发时间估算
- 定义了成功标准

### 测试策略
- 单元测试：覆盖核心逻辑
- 集成测试：端到端API调用
- 性能测试：解析性能基准
- 跨平台测试：浏览器/Node.js/Bun
- 测试覆盖率：代码覆盖率目标
- 持续集成：自动化测试流程
- 测试最佳实践：命名规范、独立性、覆盖率、可维护性

通过完善的测试策略，确保代码质量和项目稳定性。

**相关文档：**
- [产品说明](./01-product-overview.md)
- [技术栈与目录结构](./02-tech-stack.md)
- [分阶段实现步骤](./03-implementation-steps.md)
- [核心功能设计](./04-core-features.md)
- [核心代码与扩展指南](./05-core-code.md)
- [部署使用与FAQ](./07-deployment-faq.md)
