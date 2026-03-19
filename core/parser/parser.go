package parser

import (
	javascript "github.com/belos-street/coder-mate/core/language"
	"github.com/belos-street/coder-mate/core/types"
)

func Parse(code string, language string) types.TokenLines {
	switch language {
	case "javascript":
		return javascript.Parse(code)
	default:
		return types.TokenLines{}
	}
}
