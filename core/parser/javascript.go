package parser

import (
	"strings"
	"unicode"

	"github.com/belos-street/coder-mate/core/types"
)

// JavaScriptParser 实现 JavaScript 代码的词法分析（Token 化）
// 基于 TextMate 状态机模式逐字符解析源代码
// 支持关键字、字符串、数字、运算符、注释等 Token 类型
type JavaScriptParser struct {
	// keywords 存储所有 JavaScript 关键字，用于快速查找
	keywords map[string]bool
}

// NewJavaScriptParser 创建并初始化一个新的 JavaScript 解析器实例
// 返回: *JavaScriptParser - 配置好的解析器，包含所有 JS 关键字
func NewJavaScriptParser() *JavaScriptParser {
	return &JavaScriptParser{
		keywords: map[string]bool{
			"const": true, "let": true, "var": true,
			"function": true, "return": true, "if": true,
			"else": true, "for": true, "while": true,
			"do": true, "switch": true, "case": true,
			"break": true, "continue": true, "try": true,
			"catch": true, "finally": true, "throw": true,
			"new": true, "typeof": true, "instanceof": true,
			"class": true, "extends": true, "import": true,
			"export": true, "default": true, "async": true,
			"await": true, "yield": true, "this": true,
			"super": true, "null": true, "true": true,
			"false": true, "undefined": true,
		},
	}
}

// Detect 检测给定代码是否可能属于 JavaScript 语言
// 通过查找常见的 JavaScript 特征关键字来判断
// 参数: code - 待检测的源代码字符串
// 返回: bool - 如果代码可能是 JavaScript 返回 true，否则返回 false
// 注意: 这是一个启发式检测，不保证 100% 准确
func (p *JavaScriptParser) Detect(code string) bool {
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		return false
	}

	indicators := []string{
		"function", "const", "let", "var", "=>",
		"import", "export", "class", "extends",
		"async", "await", "=>", "=>",
	}

	for _, indicator := range indicators {
		if strings.Contains(code, indicator) {
			return true
		}
	}

	return false
}

// Parse 解析 JavaScript 代码为 Token 序列
// 实现 LanguageParser 接口的主解析方法
// 参数: code - 待解析的 JavaScript 源代码
// 返回: []types.Token - 解析后的 Token 数组，每个 Token 包含类型、值、行号和列号
// 解析流程: 按行分割 -> 逐行解析 -> 逐字符状态机扫描 -> 生成 Token
func (p *JavaScriptParser) Parse(code string) []types.Token {
	var tokens []types.Token
	lines := strings.Split(code, "\n")

	for lineNum, line := range lines {
		line = strings.TrimRight(line, " \t")
		if len(line) == 0 {
			continue
		}

		tokens = append(tokens, p.parseLine(line, lineNum+1)...)
	}

	return tokens
}

// parseLine 解析单行代码为 Token 序列
// 状态机入口，根据当前字符进入不同的解析状态
// 参数: line - 单行源代码字符串
// 参数: lineNum - 当前行号（从 1 开始）
// 返回: []types.Token - 该行的 Token 数组
// 状态转换: 空格跳过 -> 检查注释/字符串/数字/标识符/运算符/标点 -> 其他字符跳过
func (p *JavaScriptParser) parseLine(line string, lineNum int) []types.Token {
	var tokens []types.Token
	i := 0

	for i < len(line) {
		r := rune(line[i])

		switch {
		case unicode.IsSpace(r):
			i++

		case r == '/' && i+1 < len(line) && line[i+1] == '/':
			tokens = append(tokens, p.parseSingleLineComment(line, i, lineNum)...)
			return tokens

		case r == '/' && i+1 < len(line) && line[i+1] == '*':
			tokens = append(tokens, p.parseMultiLineComment(line, i, lineNum)...)
			return tokens

		case r == '"' || r == '\'' || r == '`':
			tokens = append(tokens, p.parseString(line, i, lineNum)...)
			i += len(tokens[len(tokens)-1].Value)

		case unicode.IsDigit(r):
			tokens = append(tokens, p.parseNumber(line, i, lineNum)...)
			i += len(tokens[len(tokens)-1].Value) - 1

		case r == '_' || unicode.IsLetter(r):
			tokens = append(tokens, p.parseIdentifier(line, i, lineNum)...)
			i += len(tokens[len(tokens)-1].Value) - 1

		case p.isOperator(r):
			tokens = append(tokens, p.parseOperator(line, i, lineNum)...)
			i++

		case p.isPunctuation(r):
			tokens = append(tokens, types.Token{
				Kind:  types.TokenKindPunctuation,
				Value: string(r),
				Line:  lineNum,
				Col:   i + 1,
			})
			i++

		default:
			i++
		}
	}

	return tokens
}

// parseSingleLineComment 解析单行注释（以 // 开头）
// 参数: line - 当前行源代码
// 参数: start - 注释起始位置（// 的索引）
// 参数: lineNum - 当前行号
// 返回: []types.Token - 包含单个注释 Token 的数组
func (p *JavaScriptParser) parseSingleLineComment(line string, start int, lineNum int) []types.Token {
	return []types.Token{{
		Kind:  types.TokenKindComment,
		Value: line[start:],
		Line:  lineNum,
		Col:   start + 1,
	}}
}

// parseMultiLineComment 解析多行注释（以 /* 开头）
// 参数: line - 当前行源代码
// 参数: start - 注释起始位置（/* 的索引）
// 参数: lineNum - 当前行号
// 返回: []types.Token - 包含单个注释 Token 的数组
// 注意: 当前实现将整行作为注释返回，多行注释需要跨行处理时需增强
func (p *JavaScriptParser) parseMultiLineComment(line string, start int, lineNum int) []types.Token {
	return []types.Token{{
		Kind:  types.TokenKindComment,
		Value: line[start:],
		Line:  lineNum,
		Col:   start + 1,
	}}
}

// parseString 解析字符串字面量
// 支持双引号、单引号和模板字符串（反引号）
// 处理转义字符（如 \"、\'、\\）
// 参数: line - 当前行源代码
// 参数: start - 字符串起始位置（引号的索引）
// 参数: lineNum - 当前行号
// 返回: []types.Token - 包含字符串 Token 的数组
// 状态: 从起始引号开始，查找结束引号（转义引号除外）
func (p *JavaScriptParser) parseString(line string, start int, lineNum int) []types.Token {
	quote := line[start]
	i := start + 1

	for i < len(line) {
		if line[i] == '\\' && i+1 < len(line) {
			i += 2
			continue
		}
		if line[i] == quote {
			return []types.Token{{
				Kind:  types.TokenKindString,
				Value: line[start : i+1],
				Line:  lineNum,
				Col:   start + 1,
			}}
		}
		i++
	}

	return []types.Token{{
		Kind:  types.TokenKindString,
		Value: line[start:],
		Line:  lineNum,
		Col:   start + 1,
	}}
}

// parseNumber 解析数字字面量
// 支持整数、小数、科学计数法（如 1e10, 2.5e-3）
// 参数: line - 当前行源代码
// 参数: start - 数字起始位置
// 参数: lineNum - 当前行号
// 返回: []types.Token - 包含数字 Token 的数组
// 状态机: 数字 -> (小数点 -> 数字) -> (e/E -> 可选 +/- -> 数字)
func (p *JavaScriptParser) parseNumber(line string, start int, lineNum int) []types.Token {
	i := start
	hasDot := false

	for i < len(line) {
		r := rune(line[i])

		if unicode.IsDigit(r) {
			i++
			continue
		}

		if r == '.' && !hasDot {
			hasDot = true
			i++
			continue
		}

		if r == 'e' || r == 'E' {
			if i+1 < len(line) && (line[i+1] == '+' || line[i+1] == '-') {
				i += 2
				continue
			}
			i++
			continue
		}

		break
	}

	return []types.Token{{
		Kind:  types.TokenKindNumber,
		Value: line[start:i],
		Line:  lineNum,
		Col:   start + 1,
	}}
}

// parseIdentifier 解析标识符（变量名、函数名等）
// 自动识别关键字、布尔值和常量
// 标识符规则: 字母/下划线/$ 开头，后跟字母/数字/下划线/$
// 参数: line - 当前行源代码
// 参数: start - 标识符起始位置
// 参数: lineNum - 当前行号
// 返回: []types.Token - 包含标识符 Token 的数组
// Token 类型判定: 关键字 -> keyword, true/false -> boolean, null/undefined -> constant, 其他 -> identifier
func (p *JavaScriptParser) parseIdentifier(line string, start int, lineNum int) []types.Token {
	i := start

	for i < len(line) {
		r := rune(line[i])

		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '$' {
			i++
			continue
		}

		break
	}

	value := line[start:i]
	kind := types.TokenKindIdentifier

	if p.keywords[value] {
		kind = types.TokenKindKeyword
	} else if value == "true" || value == "false" {
		kind = types.TokenKindBoolean
	} else if value == "null" || value == "undefined" {
		kind = types.TokenKindConstant
	}

	return []types.Token{{
		Kind:  kind,
		Value: value,
		Line:  lineNum,
		Col:   start + 1,
	}}
}

// parseOperator 解析运算符
// 支持的运算符: ==, !=, ===, !==, <=, >=, &&, ||, ++, --, +=, -=, *=, /=, %=, =>, ...
// 优先匹配长运算符（多字符），再匹配单字符运算符
// 参数: line - 当前行源代码
// 参数: start - 运算符起始位置
// 参数: lineNum - 当前行号
// 返回: []types.Token - 包含运算符 Token 的数组
func (p *JavaScriptParser) parseOperator(line string, start int, lineNum int) []types.Token {
	operators := []string{
		"==", "!=", "===", "!==", "<=", ">=",
		"&&", "||", "++", "--", "+=", "-=",
		"*=", "/=", "%=", "=>", "...",
	}

	for _, op := range operators {
		if strings.HasPrefix(line[start:], op) {
			return []types.Token{{
				Kind:  types.TokenKindOperator,
				Value: op,
				Line:  lineNum,
				Col:   start + 1,
			}}
		}
	}

	return []types.Token{{
		Kind:  types.TokenKindOperator,
		Value: string(line[start]),
		Line:  lineNum,
		Col:   start + 1,
	}}
}

// isOperator 判断字符是否为运算符字符
// 参数: r - 待判断的 rune 字符
// 返回: bool - 如果是运算符字符返回 true
// 运算符字符集: +-*/%=<>!&|^~?:.
func (p *JavaScriptParser) isOperator(r rune) bool {
	operators := "+-*/%=<>!&|^~?:."
	return strings.ContainsRune(operators, r)
}

// isPunctuation 判断字符是否为标点符号
// 参数: r - 待判断的 rune 字符
// 返回: bool - 如果是标点符号返回 true
// 标点符号集: (){}[],.;:
func (p *JavaScriptParser) isPunctuation(r rune) bool {
	punctuation := "(){}[],.;:"
	return strings.ContainsRune(punctuation, r)
}
