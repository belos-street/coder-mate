package main

import (
	"syscall/js"

	"github.com/belos-street/coder-mate/core/parser"
	"github.com/belos-street/coder-mate/core/renderer"
	"github.com/belos-street/coder-mate/core/types"
)

func main() {
	js.Global().Set("highlightCode", js.FuncOf(highlightCode))

	select {}
}

func highlightCode(this js.Value, args []js.Value) interface{} {
	if len(args) < 2 {
		return "error: insufficient arguments"
	}

	code := args[0].String()
	language := types.Language(args[1].String())

	tokens := parser.Tokenize(code, language)
	result := renderer.Render(tokens)

	return result
}
