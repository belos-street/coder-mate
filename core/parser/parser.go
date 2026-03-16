package parser

import (
	"github.com/belos-street/coder-mate/core/types"
)

func Parse(code string, language types.Language) string {
	return ""
}

func DetectLanguage(code string) types.Language {
	return types.LangUnknown
}

func Tokenize(code string, language types.Language) []types.Token {
	tokens := []types.Token{
		{Kind: types.TokenKindKeyword, Value: "const", Line: 1, Col: 1},
		{Kind: types.TokenKindIdentifier, Value: "x", Line: 1, Col: 7},
		{Kind: types.TokenKindOperator, Value: "=", Line: 1, Col: 9},
		{Kind: types.TokenKindNumber, Value: "42", Line: 1, Col: 11},
		{Kind: types.TokenKindKeyword, Value: "function", Line: 2, Col: 1},
		{Kind: types.TokenKindIdentifier, Value: "hello", Line: 2, Col: 10},
		{Kind: types.TokenKindPunctuation, Value: "(", Line: 2, Col: 15},
		{Kind: types.TokenKindPunctuation, Value: ")", Line: 2, Col: 16},
		{Kind: types.TokenKindPunctuation, Value: "{", Line: 2, Col: 18},
		{Kind: types.TokenKindComment, Value: "// This is a comment", Line: 3, Col: 3},
		{Kind: types.TokenKindKeyword, Value: "return", Line: 4, Col: 3},
		{Kind: types.TokenKindString, Value: "\"Hello World\"", Line: 4, Col: 10},
		{Kind: types.TokenKindPunctuation, Value: "}", Line: 5, Col: 1},
	}
	return tokens
}
