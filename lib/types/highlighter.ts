import type { Language } from './language'
import type { Theme } from './theme'
import type { Token } from './token'

export type CodeToHtmlOptions = {
  lang: Language
  theme: Theme
}

export type CodeToTokensOptions = {
  lang: Language
}

export type Highlighter = {
  codeToHtml(
    code: string,
    options: {
      lang: Language
    }
  ): string
  codeToTokens(
    code: string,
    options: {
      lang: Language
    }
  ): Token[]
  dispose(): void
}
