package parser

import (
	"fmt"
)

func Parse(code, language, theme string) string {
	return fmt.Sprintf("Code: %s, Language: %s, Theme: %s", code, language, theme)
}

func DetectLanguage(code string) string {
	return "unknown"
}
