package javascript

import (
	"strings"
	"testing"

	"github.com/belos-street/coder-mate/core/types"
)

func TestParse_Keyword(t *testing.T) {
	code := "const x = 1"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 4 {
		t.Fatalf("expected 4 tokens, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindKeyword || tokens[0].Value != "const" {
		t.Errorf("expected first token to be keyword 'const'")
	}
}

func TestParse_Number(t *testing.T) {
	code := "42"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindNumber || tokens[0].Value != "42" {
		t.Errorf("expected token to be number '42'")
	}
}

func TestParse_String(t *testing.T) {
	code := `"hello world"`
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindString {
		t.Errorf("expected token to be string")
	}
}

func TestParse_Identifier(t *testing.T) {
	code := "myVariable"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindIdentifier || tokens[0].Value != "myVariable" {
		t.Errorf("expected token to be identifier 'myVariable'")
	}
}

func TestParse_Operator(t *testing.T) {
	code := "a + b"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}

	if tokens[1].Kind != types.TokenKindOperator || tokens[1].Value != "+" {
		t.Errorf("expected second token to be operator '+'")
	}
}

func TestParse_Comment(t *testing.T) {
	code := "// this is a comment"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindComment {
		t.Errorf("expected token to be comment")
	}
}

func TestParse_Punctuation(t *testing.T) {
	code := "()"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindPunctuation || tokens[0].Value != "(" {
		t.Errorf("expected first token to be '('")
	}
}

func TestParse_Boolean(t *testing.T) {
	code := "true"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindBoolean || tokens[0].Value != "true" {
		t.Errorf("expected token to be boolean 'true'")
	}
}

func TestParse_Null(t *testing.T) {
	code := "null"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Kind != types.TokenKindConstant || tokens[0].Value != "null" {
		t.Errorf("expected token to be constant 'null'")
	}
}

func TestParse_EmptyLine(t *testing.T) {
	code := "a\n\nb"
	tokenLines := Parse(code)

	if len(tokenLines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(tokenLines))
	}

	if len(tokenLines[0]) != 1 {
		t.Errorf("expected first line to have 1 token")
	}

	if len(tokenLines[1]) != 0 {
		t.Errorf("expected second line to be empty")
	}

	if len(tokenLines[2]) != 1 {
		t.Errorf("expected third line to have 1 token")
	}
}

func TestParse_ES2020_BigInt(t *testing.T) {
	code := "const num = 123n"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	hasBigInt := false
	for _, token := range tokens {
		if token.Kind == types.TokenKindNumber && token.Value == "123n" {
			hasBigInt = true
			break
		}
	}

	if !hasBigInt {
		t.Error("expected to find BigInt literal '123n'")
	}
}

func TestParse_ES2020_OptionalChaining(t *testing.T) {
	code := "const city = user?.address?.city"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	hasOptionalChaining := false
	for _, token := range tokens {
		if token.Kind == types.TokenKindOperator && token.Value == "?." {
			hasOptionalChaining = true
			break
		}
	}

	if !hasOptionalChaining {
		t.Error("expected to find optional chaining operator '?.'")
	}
}

func TestParse_ES2020_NullishCoalescing(t *testing.T) {
	code := "const age = user?.age ?? 18"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	hasNullishCoalescing := false
	for _, token := range tokens {
		if token.Kind == types.TokenKindOperator && token.Value == "??" {
			hasNullishCoalescing = true
			break
		}
	}

	if !hasNullishCoalescing {
		t.Error("expected to find nullish coalescing operator '??'")
	}
}

func TestParse_ES2020_TemplateString(t *testing.T) {
	code := "const msg = `Hello, ${name}!`"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	hasInterpolation := false
	for _, token := range tokens {
		if token.Kind == types.TokenKindInterpolation {
			hasInterpolation = true
			break
		}
	}

	if !hasInterpolation {
		t.Errorf("expected to find template interpolation '${', got tokens: %+v", tokens)
	}
}

func TestParse_ES2020_GlobalThis(t *testing.T) {
	code := "console.log(globalThis)"
	tokenLines := Parse(code)

	if len(tokenLines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	hasGlobalThis := false
	for _, token := range tokens {
		if token.Kind == types.TokenKindKeyword && token.Value == "globalThis" {
			hasGlobalThis = true
			break
		}
	}

	if !hasGlobalThis {
		t.Error("expected to find 'globalThis' keyword")
	}
}

func TestParse_MultilineComment(t *testing.T) {
	code := `/* This is a
   multi-line comment */`
	tokenLines := Parse(code)

	if len(tokenLines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(tokenLines))
	}

	tokens := tokenLines[0]
	if len(tokens) != 1 {
		t.Errorf("expected first line to have 1 token")
	}

	if tokens[0].Kind != types.TokenKindComment {
		t.Errorf("expected first token to be comment")
	}

	if !strings.HasPrefix(tokens[0].Value, "/*") {
		t.Errorf("expected first token value to start with '/*', got '%s'", tokens[0].Value)
	}
}

func TestParse_CompleteCode(t *testing.T) {
	code := `// This is a single line comment
const PI = 3.14;
let name = "John";
let age = 25;

/* This is a
   multi-line comment */

function greet(user) {
	if (user) {
		return "Hello, " + user;
	}
	return "Hello, Guest";
}

class Person {
	constructor(name) {
		this.name = name;
	}

	sayHello() {
		return true;
	}
}

const person = new Person("Alice");
console.log(person.sayHello());`

	tokenLines := Parse(code)

	if len(tokenLines) == 0 {
		t.Fatal("expected token lines to be generated")
	}

	totalTokens := 0
	for _, lineTokens := range tokenLines {
		totalTokens += len(lineTokens)
	}

	if totalTokens == 0 {
		t.Fatal("expected tokens to be generated")
	}

	hasConst := false
	hasLet := false
	hasFunction := false
	hasIf := false
	hasReturn := false
	hasClass := false
	hasThis := false
	hasNew := false

	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindKeyword {
				switch token.Value {
				case "const":
					hasConst = true
				case "let":
					hasLet = true
				case "function":
					hasFunction = true
				case "if":
					hasIf = true
				case "return":
					hasReturn = true
				case "class":
					hasClass = true
				case "this":
					hasThis = true
				case "new":
					hasNew = true
				}
			}
		}
	}

	if !hasConst {
		t.Error("expected to find 'const' keyword")
	}
	if !hasLet {
		t.Error("expected to find 'let' keyword")
	}
	if !hasFunction {
		t.Error("expected to find 'function' keyword")
	}
	if !hasIf {
		t.Error("expected to find 'if' keyword")
	}
	if !hasReturn {
		t.Error("expected to find 'return' keyword")
	}
	if !hasClass {
		t.Error("expected to find 'class' keyword")
	}
	if !hasThis {
		t.Error("expected to find 'this' keyword")
	}
	if !hasNew {
		t.Error("expected to find 'new' keyword")
	}

	hasSingleLineComment := false
	hasMultiLineComment := false
	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindComment {
				if token.Value[0:2] == "//" {
					hasSingleLineComment = true
				}
				if token.Value[0:2] == "/*" {
					hasMultiLineComment = true
				}
			}
		}
	}

	if !hasSingleLineComment {
		t.Error("expected to find single-line comment")
	}
	if !hasMultiLineComment {
		t.Error("expected to find multi-line comment")
	}

	hasString := false
	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindString {
				hasString = true
				break
			}
		}
		if hasString {
			break
		}
	}

	if !hasString {
		t.Error("expected to find string literal")
	}

	hasNumber := false
	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindNumber {
				hasNumber = true
				break
			}
		}
		if hasNumber {
			break
		}
	}

	if !hasNumber {
		t.Error("expected to find number literal")
	}

	hasBoolean := false
	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindBoolean {
				hasBoolean = true
				break
			}
		}
		if hasBoolean {
			break
		}
	}

	if !hasBoolean {
		t.Error("expected to find boolean value")
	}

	hasOperator := false
	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindOperator {
				hasOperator = true
				break
			}
		}
		if hasOperator {
			break
		}
	}

	if !hasOperator {
		t.Error("expected to find operator")
	}

	hasPunctuation := false
	for _, lineTokens := range tokenLines {
		for _, token := range lineTokens {
			if token.Kind == types.TokenKindPunctuation {
				hasPunctuation = true
				break
			}
		}
		if hasPunctuation {
			break
		}
	}

	if !hasPunctuation {
		t.Error("expected to find punctuation")
	}
}

func TestParse_LineByLine(t *testing.T) {
	code := `// This is a single line comment
const PI = 3.14;`

	tokenLines := Parse(code)

	if len(tokenLines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(tokenLines))
	}

	firstLine := tokenLines[0]
	if len(firstLine) != 1 {
		t.Fatalf("expected first line to have 1 token, got %d", len(firstLine))
	}

	if firstLine[0].Kind != types.TokenKindComment {
		t.Errorf("expected first token to be comment")
	}

	secondLine := tokenLines[1]
	if len(secondLine) != 5 {
		t.Fatalf("expected second line to have 5 tokens, got %d", len(secondLine))
	}

	expectedTokens := []struct {
		kind  types.TokenKind
		value string
	}{
		{types.TokenKindKeyword, "const"},
		{types.TokenKindIdentifier, "PI"},
		{types.TokenKindOperator, "="},
		{types.TokenKindNumber, "3.14"},
		{types.TokenKindPunctuation, ";"},
	}

	for i, expected := range expectedTokens {
		if secondLine[i].Kind != expected.kind {
			t.Errorf("token %d: expected kind %s, got %s", i, expected.kind, secondLine[i].Kind)
		}
		if secondLine[i].Value != expected.value {
			t.Errorf("token %d: expected value %s, got %s", i, expected.value, secondLine[i].Value)
		}
	}
}
