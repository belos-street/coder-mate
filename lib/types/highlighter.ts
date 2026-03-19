import type { Language } from './language'
import type { Theme } from './theme'
import type { Token } from './token'

export type CodeToHtmlOptions = {
  lang: Language
}

export type CodeToTokensOptions = {
  lang: Language
}

export type Highlighter = {
  codeToHtml(code: string, options: CodeToHtmlOptions): Promise<string>
  codeToTokens(code: string, options: CodeToTokensOptions): Promise<Token[]>
  dispose(): void
  setTheme(theme: Theme): void
}
