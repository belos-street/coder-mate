import type {
  Token,
  Highlighter,
  CodeToTokensOptions,
  CodeToHtmlOptions
} from '../types'
import {
  initWasm,
  ensureWasm,
  incrementInstance,
  decrementInstance,
  getInstanceCount,
  resetWasm
} from '../loader'

export async function createHighlighter(): Promise<Highlighter> {
  await initWasm()
  incrementInstance()

  return {
    codeToHtml(code: string, opts: CodeToHtmlOptions): string {
      ensureWasm()

      const html = (globalThis as any).highlightCode(
        code,
        opts.lang,
        opts.theme
      )
      return html
    },

    codeToTokens(code: string, opts: CodeToTokensOptions): Token[] {
      ensureWasm()

      const html = (globalThis as any).highlightCode(code, opts.lang)
      return parseTokensFromHtml(html)
    },

    dispose(): void {
      decrementInstance()
      if (getInstanceCount() === 0) {
        resetWasm()
      }
    }
  }
}

function parseTokensFromHtml(html: string): Token[] {
  const parser = new DOMParser()
  const doc = parser.parseFromString(html, 'text/html')
  const spans = doc.querySelectorAll('span')

  const tokens: Token[] = []
  let line = 1
  let col = 1

  spans.forEach((span) => {
    const className = span.className || ''
    const type = className.replace('token-', '') as Token['type']
    const value = span.textContent || ''

    tokens.push({ type, value, line, col })

    const lines = value.split('\n')
    if (lines.length > 1) {
      line += lines.length - 1
      col = lines[lines.length - 1].length + 1
    } else {
      col += value.length
    }
  })

  return tokens
}
