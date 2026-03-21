import type { Language } from './language'
import type { Token } from './token'

declare global {
  function highlightCode(code: string, lang: Language, mode: 'html'): string
  function highlightCode(code: string, lang: Language, mode: 'tokens'): Token[]
}
