package renderer

import (
	"strings"

	"github.com/belos-street/coder-mate/core/types"
)

func Render(tokens []types.Token) string {
	var result string
	for _, token := range tokens {
		result += `<span class="token-` + string(token.Kind) + `">` + escapeHTML(token.Value) + `</span>`
	}
	return result
}

func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, `"`, "&quot;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	return s
}
