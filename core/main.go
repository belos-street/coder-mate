package main

import (
	"syscall/js"

	"github.com/belos-street/coder-mate/core/parser"
	"github.com/belos-street/coder-mate/core/renderer"
)

func main() {
	js.Global().Set("highlightCode", js.FuncOf(highlightCode))

	select {}
}

func highlightCode(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return "error: insufficient arguments"
	}

	code := args[0].String()
	language := args[1].String()
	theme := args[2].String()

	parsed := parser.Parse(code, language, theme)
	result := renderer.Render(parsed, renderer.GetDefaultTheme())

	return result
}
