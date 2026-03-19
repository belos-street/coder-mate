package javascript

// JavaScript tokenizer/lexer

import (
	"strings"
	"unicode"

	"github.com/belos-street/coder-mate/core/types"
)

var keywords = map[string]bool{
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
}

func Parse(code string) types.TokenLines {
	// Parse JavaScript code and return tokens by line
	var tokenLines types.TokenLines
	lines := strings.Split(code, "\n")

	for lineNum, line := range lines {
		line = strings.TrimRight(line, " \t")
		if len(line) == 0 {
			tokenLines = append(tokenLines, []types.Token{})
			continue
		}
		tokenLines = append(tokenLines, parseLine(line, lineNum+1))
	}

	return tokenLines
}

func parseLine(line string, lineNum int) []types.Token {
	// Parse a single line of JavaScript code
	var tokens []types.Token
	i := 0

	for i < len(line) {
		r := rune(line[i])

		switch {
		case unicode.IsSpace(r):
			i++

		case r == '/' && i+1 < len(line) && line[i+1] == '/':
			tokens = append(tokens, parseSingleLineComment(line, i, lineNum)...)
			return tokens

		case r == '/' && i+1 < len(line) && line[i+1] == '*':
			tokens = append(tokens, parseMultiLineComment(line, i, lineNum)...)
			return tokens

		case r == '"' || r == '\'' || r == '`':
			tokens = append(tokens, parseString(line, i, lineNum)...)
			valueLen := len(tokens[len(tokens)-1].Value)
			if valueLen > 0 && tokens[len(tokens)-1].Value[valueLen-1] == byte(r) {
				i += valueLen
			} else {
				i = len(line)
			}

		case unicode.IsDigit(r):
			tokens = append(tokens, parseNumber(line, i, lineNum)...)
			valueLen := len(tokens[len(tokens)-1].Value)
			i += valueLen
			if i > len(line) {
				i = len(line)
			}

		case r == '_' || unicode.IsLetter(r):
			tokens = append(tokens, parseIdentifier(line, i, lineNum)...)
			valueLen := len(tokens[len(tokens)-1].Value)
			i += valueLen
			if i > len(line) {
				i = len(line)
			}

		case isOperator(r):
			tokens = append(tokens, parseOperator(line, i, lineNum)...)
			i++

		case isPunctuation(r):
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

func parseSingleLineComment(line string, start int, lineNum int) []types.Token {
	// Parse single-line comment (// ...)
	return []types.Token{{
		Kind:  types.TokenKindComment,
		Value: line[start:],
		Line:  lineNum,
		Col:   start + 1,
	}}
}

func parseMultiLineComment(line string, start int, lineNum int) []types.Token {
	// Parse multi-line comment (/* ... */)
	return []types.Token{{
		Kind:  types.TokenKindComment,
		Value: line[start:],
		Line:  lineNum,
		Col:   start + 1,
	}}
}

func parseString(line string, start int, lineNum int) []types.Token {
	// Parse string literal (single/double/template literal)
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

func parseNumber(line string, start int, lineNum int) []types.Token {
	// Parse number literal (integer, float, scientific notation)
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

func parseIdentifier(line string, start int, lineNum int) []types.Token {
	// Parse identifier or keyword
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

	if value == "true" || value == "false" {
		kind = types.TokenKindBoolean
	} else if value == "null" || value == "undefined" {
		kind = types.TokenKindConstant
	} else if keywords[value] {
		kind = types.TokenKindKeyword
	}

	return []types.Token{{
		Kind:  kind,
		Value: value,
		Line:  lineNum,
		Col:   start + 1,
	}}
}

func parseOperator(line string, start int, lineNum int) []types.Token {
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

func isOperator(r rune) bool {
	operators := "+-*/%=<>!&|^~?:."
	return strings.ContainsRune(operators, r)
}

func isPunctuation(r rune) bool {
	punctuation := "(){}[],.;:"
	return strings.ContainsRune(punctuation, r)
}
