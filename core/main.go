//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
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
	mode := args[2].String()

	tokens := parser.Parse(code, language)

	if mode == "html" {
		return renderer.Render(tokens)
	}

	jsonBytes, err := json.Marshal(tokens)
	if err != nil {
		return "error: failed to marshal tokens"
	}

	return string(jsonBytes)
}
