package javascript

// JavaScript tokenizer/lexer with state machine (ES2020 support)

import (
	"strings"
	"unicode"

	"github.com/belos-street/coder-mate/core/types"
)

// Parser state types
const (
	stateGlobal                = "global"
	stateMultilineComment      = "multiline-comment"
	stateStringDouble          = "string-double"
	stateStringSingle          = "string-single"
	stateStringBacktick        = "string-backtick"
	stateTemplateInterpolation = "template-interpolation"
)

// Keywords map
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
	"super": true, "static": true, "constructor": true,
	"globalThis": true,
}

// Parser state
type parser struct {
	stateStack []string
	tokens     []types.Token
	lineNum    int
	colNum     int
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
		p := &parser{
			stateStack: []string{stateGlobal},
			lineNum:    lineNum + 1,
		}
		p.parseLine(line)
		tokenLines = append(tokenLines, p.tokens)
	}

	return tokenLines
}

func (p *parser) currentState() string {
	return p.stateStack[len(p.stateStack)-1]
}

func (p *parser) pushState(state string) {
	p.stateStack = append(p.stateStack, state)
}

func (p *parser) popState() {
	if len(p.stateStack) > 1 {
		p.stateStack = p.stateStack[:len(p.stateStack)-1]
	}
}

func (p *parser) addToken(kind types.TokenKind, value string) {
	p.tokens = append(p.tokens, types.Token{
		Kind:  kind,
		Value: value,
		Line:  p.lineNum,
		Col:   p.colNum + 1,
	})
	p.colNum += len(value)
}

func (p *parser) parseLine(line string) {
	i := 0
	for i < len(line) {
		switch p.currentState() {
		case stateGlobal:
			i = p.parseGlobal(line, i)
		case stateMultilineComment:
			i = p.parseMultilineComment(line, i)
		case stateStringDouble:
			i = p.parseString(line, i, '"', stateStringDouble)
		case stateStringSingle:
			i = p.parseString(line, i, '\'', stateStringSingle)
		case stateStringBacktick:
			i = p.parseStringBacktick(line, i)
		case stateTemplateInterpolation:
			i = p.parseTemplateInterpolation(line, i)
		default:
			i++
		}
	}
}

func (p *parser) parseGlobal(line string, i int) int {
	r := rune(line[i])

	switch {
	case unicode.IsSpace(r):
		i++

	case r == '/' && i+1 < len(line) && line[i+1] == '/':
		p.addToken(types.TokenKindComment, line[i:])
		i = len(line)

	case r == '/' && i+1 < len(line) && line[i+1] == '*':
		p.addToken(types.TokenKindComment, line[i:i+2])
		p.pushState(stateMultilineComment)
		i += 2

	case r == '"':
		i = p.parseString(line, i, '"', stateStringDouble)

	case r == '\'':
		i = p.parseString(line, i, '\'', stateStringSingle)

	case r == '`':
		i = p.parseStringBacktick(line, i)

	case r == '.' && i+1 < len(line) && line[i+1] == '.' && i+2 < len(line) && line[i+2] == '.':
		p.addToken(types.TokenKindOperator, "...")
		i += 3

	case r == '.' && i+1 < len(line) && (unicode.IsDigit(rune(line[i+1])) || line[i+1] == '.'):
		p.addToken(types.TokenKindOperator, ".")
		i++

	case r == '?' && i+1 < len(line) && line[i+1] == '.':
		p.addToken(types.TokenKindOperator, "?.")
		i += 2

	case r == '?' && i+1 < len(line) && line[i+1] == '?':
		p.addToken(types.TokenKindOperator, "??")
		i += 2

	case r == '=' && i+1 < len(line) && line[i+1] == '=' && i+2 < len(line) && line[i+2] == '=':
		p.addToken(types.TokenKindOperator, "===")
		i += 3

	case r == '!' && i+1 < len(line) && line[i+1] == '=' && i+2 < len(line) && line[i+2] == '=':
		p.addToken(types.TokenKindOperator, "!==")
		i += 3

	case r == '=' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "==")
		i += 2

	case r == '!' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "!=")
		i += 2

	case r == '<' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "<=")
		i += 2

	case r == '>' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, ">=")
		i += 2

	case r == '&' && i+1 < len(line) && line[i+1] == '&':
		p.addToken(types.TokenKindOperator, "&&")
		i += 2

	case r == '|' && i+1 < len(line) && line[i+1] == '|':
		p.addToken(types.TokenKindOperator, "||")
		i += 2

	case r == '+' && i+1 < len(line) && line[i+1] == '+':
		p.addToken(types.TokenKindOperator, "++")
		i += 2

	case r == '-' && i+1 < len(line) && line[i+1] == '-':
		p.addToken(types.TokenKindOperator, "--")
		i += 2

	case r == '=' && i+1 < len(line) && line[i+1] == '>':
		p.addToken(types.TokenKindOperator, "=>")
		i += 2

	case r == '+' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "+=")
		i += 2

	case r == '-' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "-=")
		i += 2

	case r == '*' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "*=")
		i += 2

	case r == '/' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "/=")
		i += 2

	case r == '%' && i+1 < len(line) && line[i+1] == '=':
		p.addToken(types.TokenKindOperator, "%=")
		i += 2

	case isOperator(r):
		p.addToken(types.TokenKindOperator, string(r))
		i++

	case isPunctuation(r):
		p.addToken(types.TokenKindPunctuation, string(r))
		i++

	case unicode.IsDigit(r):
		i = p.parseNumber(line, i)

	case r == '_' || unicode.IsLetter(r):
		i = p.parseIdentifier(line, i)

	default:
		i++
	}

	return i
}

func (p *parser) parseMultilineComment(line string, i int) int {
	if i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
		p.tokens[len(p.tokens)-1].Value += "*/"
		p.popState()
		i += 2
	} else {
		p.tokens[len(p.tokens)-1].Value += string(line[i])
		i++
	}
	return i
}

func (p *parser) parseString(line string, i int, quote rune, state string) int {
	start := i
	i++
	for i < len(line) {
		if line[i] == '\\' && i+1 < len(line) {
			i += 2
			continue
		}
		if line[i] == byte(quote) {
			p.addToken(types.TokenKindString, line[start:i+1])
			p.popState()
			return i + 1
		}
		i++
	}
	if i > start {
		p.addToken(types.TokenKindString, line[start:i])
		p.popState()
	}
	return i
}

func (p *parser) parseStringBacktick(line string, i int) int {
	start := i
	i++
	for i < len(line) {
		if line[i] == '\\' && i+1 < len(line) {
			i += 2
			continue
		}
		if line[i] == '$' && i+1 < len(line) && line[i+1] == '{' {
			p.addToken(types.TokenKindString, line[start:i])
			p.addToken(types.TokenKindInterpolation, "${")
			p.pushState(stateTemplateInterpolation)
			return i + 2
		}
		if line[i] == '`' {
			p.addToken(types.TokenKindString, line[start:i+1])
			p.popState()
			return i + 1
		}
		i++
	}
	if i > start {
		p.addToken(types.TokenKindString, line[start:i])
	}
	return i
}

func (p *parser) parseTemplateInterpolation(line string, i int) int {
	r := rune(line[i])

	switch {
	case r == '}':
		p.addToken(types.TokenKindInterpolation, "}")
		p.popState()
		i++

	case unicode.IsSpace(r):
		i++

	case r == '"' || r == '\'' || r == '`':
		if r == '"' {
			p.pushState(stateStringDouble)
		} else if r == '\'' {
			p.pushState(stateStringSingle)
		} else {
			p.pushState(stateStringBacktick)
		}
		i++

	case unicode.IsDigit(r):
		i = p.parseNumber(line, i)

	case r == '_' || unicode.IsLetter(r):
		i = p.parseIdentifier(line, i)

	case isOperator(r):
		p.addToken(types.TokenKindOperator, string(r))
		i++

	case isPunctuation(r):
		p.addToken(types.TokenKindPunctuation, string(r))
		i++

	default:
		i++
	}

	return i
}

func (p *parser) parseNumber(line string, i int) int {
	start := i
	hasDot := false
	hasE := false

	for i < len(line) {
		r := rune(line[i])

		if unicode.IsDigit(r) {
			i++
			continue
		}

		if r == '.' && !hasDot && !hasE {
			hasDot = true
			i++
			continue
		}

		if (r == 'e' || r == 'E') && !hasE {
			hasE = true
			i++
			if i < len(line) && (line[i] == '+' || line[i] == '-') {
				i++
			}
			continue
		}

		if r == 'n' && i == start+1 {
			i++
			continue
		}

		if r == 'n' && (unicode.IsDigit(rune(line[i-1])) || line[i-1] == '.') {
			i++
			continue
		}

		break
	}

	value := line[start:i]
	if strings.HasPrefix(value, "0b") || strings.HasPrefix(value, "0B") ||
		strings.HasPrefix(value, "0o") || strings.HasPrefix(value, "0O") ||
		strings.HasPrefix(value, "0x") || strings.HasPrefix(value, "0X") {
		p.addToken(types.TokenKindNumber, value)
	} else {
		p.addToken(types.TokenKindNumber, value)
	}

	return i
}

func (p *parser) parseIdentifier(line string, i int) int {
	start := i

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

	p.addToken(kind, value)
	return i
}

func isOperator(r rune) bool {
	operators := "+-*/%=<>!&|^~?:."
	return strings.ContainsRune(operators, r)
}

func isPunctuation(r rune) bool {
	punctuation := "(){}[],.;:"
	return strings.ContainsRune(punctuation, r)
}
