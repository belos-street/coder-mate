package types

type TokenKind string

const (
	TokenKindKeyword       TokenKind = "keyword"       // 关键字: if, for, function, const, class, return, import
	TokenKindString        TokenKind = "string"        // 字符串: "hello", 'world', `template`
	TokenKindNumber        TokenKind = "number"        // 数字: 42, 3.14, 0xFF
	TokenKindIdentifier    TokenKind = "identifier"    // 标识符: 变量名、函数名
	TokenKindOperator      TokenKind = "operator"      // 运算符: =, +, -, *, /, ==, =>, <>
	TokenKindPunctuation   TokenKind = "punctuation"   // 标点: (), {}, [], ,, ., ;, :
	TokenKindComment       TokenKind = "comment"       // 注释: // xxx, /* xxx */, # xxx, <!-- xxx -->
	TokenKindFunction      TokenKind = "function"      // 函数名: functionName(
	TokenKindVariable      TokenKind = "variable"      // 变量: let x, var y, int z
	TokenKindType          TokenKind = "type"          // 类型声明: int, string, void, class, interface
	TokenKindConstant      TokenKind = "constant"      // 常量: const PI = 3.14, enum
	TokenKindProperty      TokenKind = "property"      // 对象属性: obj.property
	TokenKindTag           TokenKind = "tag"           // HTML/XML 标签: <div>, </span>
	TokenKindAttribute     TokenKind = "attribute"     // HTML 属性: href, class, src
	TokenKindDecorator     TokenKind = "decorator"     // 装饰器: @decorator (Python, TS)
	TokenKindNamespace     TokenKind = "namespace"     // 命名空间: ::, .
	TokenKindInterpolation TokenKind = "interpolation" // 字符串插值: ${x}, {x}
	TokenKindRegex         TokenKind = "regex"         // 正则表达式: /pattern/g
	TokenKindBoolean       TokenKind = "boolean"       // 布尔值: true, false, null, undefined
	TokenKindText          TokenKind = "text"          // 普通文本
	TokenKindUnknown       TokenKind = "unknown"       // 未知类型
)

type Token struct {
	Kind  TokenKind `json:"type"`  // Token 类型
	Value string    `json:"value"` // Token 值
	Line  int       `json:"line"`  // 行号 (从 1 开始)
	Col   int       `json:"col"`   // 列号 (从 1 开始)
}

type TokenList []Token

type TokenLines []TokenList
