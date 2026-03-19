package parser

import (
	"github.com/belos-street/coder-mate/core/types"
)

// LanguageParser 定义语言解析器的接口
// 所有语言解析器（如 JavaScript、Python 等）都必须实现此接口
// 接口提供两个核心方法：Parse 用于解析代码，Detect 用于检测语言类型
type LanguageParser interface {
	// Parse 解析给定的代码字符串，返回 Token 列表
	// 参数: code - 待解析的源代码字符串
	// 返回: []types.Token - 解析后的 Token 数组
	Parse(code string) []types.Token

	// Detect 检测给定的代码是否属于该解析器对应的语言
	// 参数: code - 待检测的代码字符串
	// 返回: bool - 如果代码属于该语言返回 true，否则返回 false
	Detect(code string) bool
}

// Parser 语言解析器工厂，负责管理所有已注册的语言解析器
// 使用注册模式将具体的语言解析器添加到内部映射中
type Parser struct {
	// parsers 存储所有已注册的语言解析器，key 为语言类型，value 为解析器实例
	parsers map[types.Language]LanguageParser
}

// New 创建一个新的 Parser 实例
// 返回: *Parser - 初始化好的 Parser 指针，内部包含空的解析器映射
func New() *Parser {
	return &Parser{
		parsers: make(map[types.Language]LanguageParser),
	}
}

// Register 注册一个语言解析器到 Parser 中
// 参数: lang - 语言类型（如 JavaScript、Python 等）
// 参数: parser - 实现了 LanguageParser 接口的解析器实例
// 注册后可以通过 Parser.Parse() 方法使用该解析器
func (p *Parser) Register(lang types.Language, parser LanguageParser) {
	p.parsers[lang] = parser
}

// Parse 根据指定语言解析代码
// 参数: code - 待解析的源代码字符串
// 参数: language - 目标语言类型
// 返回: []types.Token - 解析后的 Token 数组
// 如果指定语言未注册解析器，返回空 Token 数组
func (p *Parser) Parse(code string, language types.Language) []types.Token {
	if parser, ok := p.parsers[language]; ok {
		return parser.Parse(code)
	}
	return []types.Token{}
}

// DetectLanguage 自动检测给定代码的语言类型
// 参数: code - 待检测的代码字符串
// 返回: types.Language - 检测到的语言类型
// 如果所有已注册解析器都无法识别，返回 types.LangUnknown
func (p *Parser) DetectLanguage(code string) types.Language {
	for lang, parser := range p.parsers {
		if parser.Detect(code) {
			return lang
		}
	}
	return types.LangUnknown
}
