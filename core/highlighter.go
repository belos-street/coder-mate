package main

import (
	"syscall/js"
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

	result := parseCode(code, language, theme)
	return result
}

func parseCode(code, language, theme string) string {
	return "Hello from WASM! Code: " + code + ", Language: " + language + ", Theme: " + theme
}