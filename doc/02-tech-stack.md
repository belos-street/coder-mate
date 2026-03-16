# 技术栈与目录结构

## 一、技术栈与运行环境

### 1.1 技术栈（分层架构 + 职责明确，各司其职）

本项目采用**双层架构设计**，严格遵循「关注点分离」原则，核心解析与上层能力完全解耦，是现代高性能前端库的标准设计模式，与原生 Shiki 架构一致：

#### ✅ 底层核心解析层：Golang + WASM 编译

**核心职责：**
- 实现「TextMate 规范的有限状态机」
- 语法规则匹配、纯文本切分、标准 Token 数组生成
- 处理所有纯 CPU 密集型的解析逻辑

**技术选型原因：**
- Go 语法简洁易上手，学习成本远低于 Rust
- 编译为 WASM 后具备编译型语言的极致性能
- Go 对 WASM 原生内置支持，无需额外编译工具，编译命令极简
- 完美适配「逐行流式解析」，内存占用极低

**核心约束：**
- **仅做文本特征匹配+状态机流转，不做任何代码语义解析（无 AST）**
- 保证与 VSCode/Shiki 高亮逻辑一致性

#### ✅ 上层封装扩展层：TypeScript

**核心职责：**
- 提供统一的跨平台 API
- 多语言注册与管理
- 主题注册与切换
- 代码装饰器能力
- Token 转 HTML 渲染
- WASM 桥接与初始化
- 全平台兼容适配

**技术选型原因：**
- TypeScript 提供强类型约束，保证代码健壮性
- 天然兼容浏览器/Node.js/Bun 环境
- 上层逻辑无性能瓶颈，用 TS 开发效率更高
- 类型定义可完美对齐 Token/语言/主题的标准结构

### 1.2 运行平台（全平台无缝兼容，无差异化）

✅ **浏览器环境**（Chrome/Firefox/Safari/Edge）
✅ **Node.js 环境**（v16+）
✅ **Bun 环境**（最新稳定版）

> **核心要求**：三个平台的 API 调用方式**完全一致**，无平台专属代码，无额外适配成本。

---

## 二、标准化项目文件目录结构

### 核心设计原则

目录结构遵循「**模块化、可扩展、低耦合**」原则：
- 所有扩展能力（新增语言/主题/装饰器）均无需修改核心代码
- 仅需在对应模块新增文件即可
- 核心解析层与上层封装层完全隔离
- 便于独立编译与调试

### 完整目录结构（所有目录/文件均为必要，无冗余）

```plain
├── root/                  # 项目根目录
│
├── core/                 # Golang 底层解析源码（核心）
│   ├── src/
│   │   ├── highlighter.go # Go核心文件：TextMate状态机+Token生成+WASM导出
│   │   └── state_machine.go # 状态机实现
│   ├── wasm/
│   │   └── highlighter.wasm # 编译后的WASM二进制文件
│   └── go.mod             # Go依赖配置文件
│
├── lib/                  # TypeScript 上层封装源码（核心）
│   ├── src/
│   │   ├── types/         # 全局类型定义目录 - 所有核心类型统一管理
│   │   │   └── index.ts   # Token/语言/主题/装饰器 标准类型定义
│   │   ├── core/          # 核心API封装目录
│   │   │   └── index.ts   # codeToTokens/codeToHtml 核心方法封装
│   │   ├── language/      # 多语言管理目录 - 新增语言仅需在此目录新增文件
│   │   │   ├── index.ts   # 语言注册/切换/获取 核心方法
│   │   │   └── js.ts      # JS语言规则模板（示例）
│   │   ├── theme/         # 主题管理目录 - 新增主题仅需在此目录新增文件
│   │   │   ├── index.ts   # 主题注册/切换/获取 核心方法
│   │   │   ├── dark.ts    # 默认暗色主题（示例）
│   │   │   └── light.ts   # 默认亮色主题（示例）
│   │   ├── decorator/     # 装饰器核心目录
│   │   │   └── index.ts   # 装饰器注册/执行/管理 核心方法
│   │   ├── wasm/          # WASM桥接目录
│   │   │   ├── loader.ts  # WASM初始化/加载 方法
│   │   │   └── bridge.ts  # Go-WASM 方法调用桥接
│   │   ├── utils/         # 工具方法目录
│   │   │   └── index.ts   # 字符串转义、数组处理等通用工具
│   │   └── index.ts       # 库的全局入口 - 暴露所有对外API
│   └── dist/              # TS编译后的JS+类型声明文件
│
├── doc/                   # 文档目录
├── tsconfig.json          # TS编译配置
└── README.md              # 项目说明文档
```

---

## 三、目录结构详解

### 3.1 源码目录

#### Golang 源码（core/）

**core/src/highlighter.go**
- Go 核心文件
- 实现 TextMate 有限状态机
- Token 数组生成
- WASM 方法导出

**core/src/state_machine.go**
- 状态机实现
- 状态转换逻辑
- 语法规则匹配

**core/go.mod**
- Go 依赖配置文件
- 定义模块名称和版本
- 无需额外依赖（Go 原生支持 WASM）

**core/wasm/highlighter.wasm**
- Go 编译后的 WASM 二进制文件
- 需要被正确加载和初始化

#### TypeScript 源码（lib/）

**types/index.ts**
- 定义所有核心类型：Token、LanguageConfig、ThemeConfig、DecoratorConfig
- 类型定义遵循 TextMate 规范
- 与 Go 端数据结构完全对齐

**core/index.ts**
- 封装核心 API：codeToTokens、codeToHtml
- 提供 Token 到 HTML 的转换逻辑
- 处理主题样式映射

**language/index.ts**
- 语言注册、切换、获取的核心方法
- 管理已注册的语言配置
- 支持多语言共存切换

**language/js.ts**
- JavaScript 语言规则模板
- 展示如何编写语言配置
- 可作为其他语言的参考模板

**theme/index.ts**
- 主题注册、切换、获取的核心方法
- 管理已注册的主题配置
- 支持主题动态切换

**theme/dark.ts**
- 默认暗色主题配置
- 包含常用语法作用域的样式映射
- 可作为自定义主题的参考模板

**theme/light.ts**
- 默认亮色主题配置
- 包含常用语法作用域的样式映射
- 可作为自定义主题的参考模板

**decorator/index.ts**
- 装饰器注册、执行、管理方法
- 支持全局装饰和精准节点装饰
- 装饰器执行错误隔离

**wasm/loader.ts**
- WASM 初始化和加载逻辑
- 处理 WASM 文件加载错误
- 管理 WASM 实例生命周期

**wasm/bridge.ts**
- Go-WASM 方法调用桥接
- JSON 序列化/反序列化
- 错误处理和异常捕获

**utils/index.ts**
- 通用工具方法
- 字符串转义、数组处理等
- 跨平台兼容性处理

**index.ts**
- 库的全局入口文件
- 暴露所有对外 API
- 初始化 WASM

### 3.2 编译输出目录

**lib/dist/**
- TypeScript 编译后的 JavaScript 文件
- 类型声明文件（.d.ts）
- 可直接被浏览器/Node.js/Bun 使用

**core/wasm/**
- Go 编译后的 WASM 二进制文件
- highlighter.wasm
- 需要被正确加载和初始化

### 3.3 配置文件

**tsconfig.json**
- TypeScript 编译配置
- 编译目标：ES6
- 输出目录：dist/ts
- 开启类型声明

**go.mod**
- Go 项目根依赖配置
- 模块初始化配置
- 无需额外依赖

**README.md**
- 项目说明文档
- 快速开始指南
- API 使用示例

---

## 四、技术架构图

```
┌─────────────────────────────────────────────────────────────┐
│                    用户应用层                              │
│  (浏览器 / Node.js / Bun)                                │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│              TypeScript 上层封装层                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │  API 封装     │  │  扩展管理     │  │  渲染逻辑     │ │
│  │  codeToHtml  │  │  多语言/主题   │  │  Token→HTML  │ │
│  └──────────────┘  └──────────────┘  └──────────────┘ │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │  装饰器系统   │  │  WASM 桥接    │  │  工具方法     │ │
│  │  全局/节点    │  │  loader/bridge│  │  转义/处理    │ │
│  └──────────────┘  └──────────────┘  └──────────────┘ │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│              Golang 底层解析层 (WASM)                       │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │  状态机引擎   │  │  语法匹配     │  │  Token 生成   │ │
│  │  TextMate FSM │  │  正则/字符串   │  │  JSON 序列化   │ │
│  └──────────────┘  └──────────────┘  └──────────────┘ │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                   纯文本代码输入                            │
└─────────────────────────────────────────────────────────────┘
```

---

## 五、开发环境要求

### 5.1 必需工具

**Node.js 环境**
- Node.js v16+
- Bun（推荐）或 npm 或 yarn 或 pnpm

**Go 环境**
- Go 1.11+（支持 WASM 编译）
- Go 原生支持 WASM，无需额外工具

**TypeScript**
- TypeScript 4.0+
- tsconfig.json 配置

### 5.2 可选工具

**测试工具**
- Jest（单元测试）
- benchmark.js（性能测试）

**代码质量工具**
- ESLint（代码规范）
- Prettier（代码格式化）

**版本控制**
- Git

---

## 六、编译与构建

### 6.1 TypeScript 编译

```bash
# 编译 TS 代码
bun run build:ts

# 或直接使用 tsc
bunx tsc
```

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

### 6.2 Go 编译为 WASM

```bash
# 进入 core/src 目录
cd core/src

# 编译为 WASM
GOOS=js GOARCH=wasm go build -o ../wasm/highlighter.wasm highlighter.go
```

**编译说明：**
- `GOOS=js`：目标操作系统为 JavaScript
- `GOARCH=wasm`：目标架构为 WebAssembly
- 输出文件：`core/wasm/highlighter.wasm`

### 6.3 一键构建

在 `package.json` 中配置脚本：

```json
{
  "scripts": {
    "build:ts": "tsc",
    "build:wasm": "cd core/src && GOOS=js GOARCH=wasm go build -o ../wasm/highlighter.wasm highlighter.go",
    "build": "bun run build:ts && bun run build:wasm"
  }
}
```

---

## 七、开发工作流

### 7.1 开发流程

1. **修改 TypeScript 代码**
   - 编辑 `lib/src/` 下的文件
   - 运行 `bun run build:ts` 编译
   - 测试编译后的代码

2. **修改 Go 代码**
   - 编辑 `core/src/highlighter.go`
   - 运行 `bun run build:wasm` 编译
   - 测试 WASM 功能

3. **测试**
   - 运行单元测试：`bun test`
   - 运行集成测试：`bun run test:integration`
   - 运行性能测试：`bun run test:benchmark`

### 7.2 调试技巧

**TypeScript 调试**
- 使用 VS Code 调试器
- 设置断点调试
- 查看编译后的 JS 代码

**Go 调试**
- 在 Go 代码中添加日志
- 使用 `fmt.Printf` 输出调试信息
- 检查 WASM 导出的方法

**WASM 调试**
- 使用浏览器开发者工具
- 查看 WebAssembly 控制台
- 检查内存使用情况

---

## 八、性能优化建议

### 8.1 Go 端优化

- 使用 TinyGo 编译减小 WASM 体积
- 减少正则表达式使用，改用字符串匹配
- 避免在循环中创建临时对象
- 使用预编译的正则表达式

### 8.2 TypeScript 端优化

- 开启代码压缩
- 使用 Tree Shaking 移除未使用代码
- 懒加载非核心功能
- 缓存解析结果

### 8.3 WASM 优化

- 预加载 WASM 文件
- 使用 CDN 加速
- 考虑使用 Service Worker 缓存
- 增加内存池大小

---

**相关文档：**
- [产品说明](./01-product-overview.md)
- [分阶段实现步骤](./03-implementation-steps.md)
- [核心功能设计](./04-core-features.md)
