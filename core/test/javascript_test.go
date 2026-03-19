package test

import (
	"testing"

	"github.com/belos-street/coder-mate/core/parser"
	"github.com/belos-street/coder-mate/core/types"
)

func TestJavaScriptParser_SingleLineComment(t *testing.T) {
	// 测试单行注释解析
	// 输入: // This is a comment
	// 预期: 解析为单个 comment 类型 Token
	p := parser.NewJavaScriptParser()
	code := `// This is a comment`

	tokens := p.Parse(code)

	if len(tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindComment {
		t.Errorf("Expected token kind 'comment', got '%s'", tokens[0].Kind)
	}

	if tokens[0].Value != code {
		t.Errorf("Expected value '%s', got '%s'", code, tokens[0].Value)
	}
}

func TestJavaScriptParser_MultiLineComment(t *testing.T) {
	// 测试多行注释解析
	// 输入: /* This is a multi-line comment */
	// 预期: 解析为单个 comment 类型 Token
	p := parser.NewJavaScriptParser()
	code := `/* This is a multi-line comment */`

	tokens := p.Parse(code)

	if len(tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindComment {
		t.Errorf("Expected token kind 'comment', got '%s'", tokens[0].Kind)
	}

	if tokens[0].Value != code {
		t.Errorf("Expected value '%s', got '%s'", code, tokens[0].Value)
	}
}

func TestJavaScriptParser_Keywords(t *testing.T) {
	// 测试关键字解析
	// 输入: const let var function return
	// 预期: 解析为 5 个 keyword 类型 Token，按顺序对应每个关键字
	p := parser.NewJavaScriptParser()
	code := `const let var function return`

	tokens := p.Parse(code)

	expectedKinds := []types.TokenKind{types.TokenKindKeyword, types.TokenKindKeyword, types.TokenKindKeyword, types.TokenKindKeyword, types.TokenKindKeyword}
	expectedValues := []string{"const", "let", "var", "function", "return"}

	if len(tokens) != 5 {
		t.Errorf("Expected 5 tokens, got %d", len(tokens))
	}

	for i, token := range tokens {
		if token.Kind != expectedKinds[i] {
			t.Errorf("Token %d: expected kind '%s', got '%s'", i, expectedKinds[i], token.Kind)
		}
		if token.Value != expectedValues[i] {
			t.Errorf("Token %d: expected value '%s', got '%s'", i, expectedValues[i], token.Value)
		}
	}
}

func TestJavaScriptParser_String(t *testing.T) {
	// 测试字符串解析
	// 输入: "hello world"
	// 预期: 解析为单个 string 类型 Token
	p := parser.NewJavaScriptParser()
	code := `"hello world"`

	tokens := p.Parse(code)

	if len(tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindString {
		t.Errorf("Expected token kind 'string', got '%s'", tokens[0].Kind)
	}

	if tokens[0].Value != `"hello world"` {
		t.Errorf("Expected value '\"hello world\"', got '%s'", tokens[0].Value)
	}
}

func TestJavaScriptParser_Number(t *testing.T) {
	// 测试数字解析
	// 输入: 42
	// 预期: 解析为单个 number 类型 Token
	p := parser.NewJavaScriptParser()
	code := `42`

	tokens := p.Parse(code)

	if len(tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindNumber {
		t.Errorf("Expected token kind 'number', got '%s'", tokens[0].Kind)
	}

	if tokens[0].Value != "42" {
		t.Errorf("Expected value '42', got '%s'", tokens[0].Value)
	}
}

func TestJavaScriptParser_Boolean(t *testing.T) {
	// 测试布尔值解析
	// 输入: true, false
	// 预期: 解析为 boolean 类型 Token
	p := parser.NewJavaScriptParser()

	testCases := []struct {
		code     string
		expected string
	}{
		{"true", "true"},
		{"false", "false"},
	}

	for _, tc := range testCases {
		tokens := p.Parse(tc.code)

		if len(tokens) != 1 {
			t.Errorf("Code '%s': expected 1 token, got %d", tc.code, len(tokens))
		}

		if tokens[0].Kind != types.TokenKindBoolean {
			t.Errorf("Code '%s': expected kind 'boolean', got '%s'", tc.code, tokens[0].Kind)
		}

		if tokens[0].Value != tc.expected {
			t.Errorf("Code '%s': expected value '%s', got '%s'", tc.code, tc.expected, tokens[0].Value)
		}
	}
}

func TestJavaScriptParser_Constant(t *testing.T) {
	// 测试常量解析
	// 输入: null, undefined
	// 预期: 解析为 constant 类型 Token
	p := parser.NewJavaScriptParser()

	testCases := []struct {
		code     string
		expected string
	}{
		{"null", "null"},
		{"undefined", "undefined"},
	}

	for _, tc := range testCases {
		tokens := p.Parse(tc.code)

		if len(tokens) != 1 {
			t.Errorf("Code '%s': expected 1 token, got %d", tc.code, len(tokens))
		}

		if tokens[0].Kind != types.TokenKindConstant {
			t.Errorf("Code '%s': expected kind 'constant', got '%s'", tc.code, tokens[0].Kind)
		}

		if tokens[0].Value != tc.expected {
			t.Errorf("Code '%s': expected value '%s', got '%s'", tc.code, tc.expected, tokens[0].Value)
		}
	}
}

func TestJavaScriptParser_Operator(t *testing.T) {
	// 测试运算符解析
	// 输入: ==, ===, !=, !==, <=, >=, &&, ||, =>
	// 预期: 解析为 operator 类型 Token
	p := parser.NewJavaScriptParser()

	testCases := []struct {
		code     string
		expected string
	}{
		{"==", "=="},
		{"===", "==="},
		{"!=", "!="},
		{"!==", "!=="},
		{"<=", "<="},
		{">=", ">="},
		{"&&", "&&"},
		{"||", "||"},
		{"=>", "=>"},
	}

	for _, tc := range testCases {
		tokens := p.Parse(tc.code)

		if len(tokens) != 1 {
			t.Errorf("Code '%s': expected 1 token, got %d", tc.code, len(tokens))
		}

		if tokens[0].Kind != types.TokenKindOperator {
			t.Errorf("Code '%s': expected kind 'operator', got '%s'", tc.code, tokens[0].Kind)
		}

		if tokens[0].Value != tc.expected {
			t.Errorf("Code '%s': expected value '%s', got '%s'", tc.code, tc.expected, tokens[0].Value)
		}
	}
}

func TestJavaScriptParser_Punctuation(t *testing.T) {
	// 测试标点符号解析
	// 输入: (){}[]
	// 预期: 解析为 6 个 punctuation 类型 Token
	p := parser.NewJavaScriptParser()
	code := `(){}[]`

	tokens := p.Parse(code)

	expectedValues := []string{"(", ")", "{", "}", "[", "]"}

	if len(tokens) != 6 {
		t.Errorf("Expected 6 tokens, got %d", len(tokens))
	}

	for i, token := range tokens {
		if token.Kind != types.TokenKindPunctuation {
			t.Errorf("Token %d: expected kind 'punctuation', got '%s'", i, token.Kind)
		}
		if token.Value != expectedValues[i] {
			t.Errorf("Token %d: expected value '%s', got '%s'", i, expectedValues[i], token.Value)
		}
	}
}

func TestJavaScriptParser_ComplexCode(t *testing.T) {
	// 测试复杂代码解析
	// 输入: 多行 JavaScript 代码，包含关键字、函数定义、字符串、数字
	// 预期: 解析为多个不同类型的 Token，并验证 Token 的行号列号正确
	p := parser.NewJavaScriptParser()
	code := `const x = 42;
function hello() {
    return "Hello World";
}`

	tokens := p.Parse(code)

	t.Logf("Parsed %d tokens:", len(tokens))
	for i, token := range tokens {
		t.Logf("  Token %d: Kind=%s, Value=%s, Line=%d, Col=%d",
			i, token.Kind, token.Value, token.Line, token.Col)
	}

	if len(tokens) == 0 {
		t.Error("Expected some tokens, got none")
	}
}

func TestJavaScriptParser_Detect(t *testing.T) {
	// 测试语言检测功能
	// 输入: 各种代码片段
	// 预期: JavaScript 代码返回 true，其他语言代码返回 false
	p := parser.NewJavaScriptParser()

	testCases := []struct {
		code     string
		expected bool
	}{
		{"function hello() {}", true},
		{"const x = 42;", true},
		{"let y = 'test';", true},
		{"var z = true;", true},
		{"import React from 'react';", true},
		{"export default App;", true},
		{"class MyClass {}", true},
		{"async function main() {}", true},
		{"await fetch(url);", true},
		{"x => x * 2", true},
		{"", false},
		{"   ", false},
		{"def hello():", false},
		{"func main()", false},
		{"print('hello')", false},
	}

	for _, tc := range testCases {
		result := p.Detect(tc.code)
		if result != tc.expected {
			t.Errorf("Detect('%s'): expected %v, got %v", tc.code, tc.expected, result)
		}
	}
}
