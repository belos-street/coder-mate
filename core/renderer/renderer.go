package renderer

import (
	"fmt"
	"strings"
)

type Theme struct {
	Name      string
	Highlight string
	Background string
}

func Render(code string, theme Theme) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("<pre><code class=\"theme-%s\">", theme.Name))
	builder.WriteString(code)
	builder.WriteString("</code></pre>")
	return builder.String()
}

func GetDefaultTheme() Theme {
	return Theme{
		Name:       "default",
		Highlight: "#000000",
		Background: "#ffffff",
	}
}
