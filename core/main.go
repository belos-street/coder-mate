//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/belos-street/coder-mate/core/parser"
	"github.com/belos-street/coder-mate/core/renderer"
	"github.com/belos-street/coder-mate/core/types"
)

var globalParser *parser.Parser

func main() {
	globalParser = parser.New()
	globalParser.Register(types.LangJavaScript, parser.NewJavaScriptParser())
	js.Global().Set("highlightCode", js.FuncOf(highlightCode))
	select {}
}

func highlightCode(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return "error: insufficient arguments"
	}

	code := args[0].String()
	language := types.Language(args[1].String())
	mode := types.Mode(args[2].String())

	tokens := globalParser.Parse(code, language)

	if mode == types.ModeHTML {
		return renderer.Render(tokens)
	}

	jsonBytes, err := json.Marshal(tokens)
	if err != nil {
		return "error: failed to marshal tokens"
	}

	return string(jsonBytes)
}
