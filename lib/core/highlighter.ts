import type {
  Token,
  Highlighter,
  CodeToTokensOptions,
  CodeToHtmlOptions,
  WorkerMessage,
  WorkerResponse,
  Theme
} from '../types'

export async function createHighlighter(): Promise<Highlighter> {
  const worker = new Worker(new URL('./worker.js', import.meta.url))

  return {
    async codeToHtml(code: string, opts: CodeToHtmlOptions): Promise<string> {
      const result = await postMessage(worker, {
        type: 'highlight',
        code,
        lang: opts.lang,
        mode: 'html'
      })
      return result as string
    },

    async codeToTokens(
      code: string,
      opts: CodeToTokensOptions
    ): Promise<Token[]> {
      const result = await postMessage(worker, {
        type: 'highlight',
        code,
        lang: opts.lang,
        mode: 'tokens'
      })
      return result as Token[]
    },

    dispose(): void {
      worker.postMessage({ type: 'dispose' })
      worker.terminate()
    },

    setTheme
  }
}

function postMessage(worker: Worker, data: WorkerMessage): Promise<unknown> {
  return new Promise((resolve, reject) => {
    const handler = (event: MessageEvent<WorkerResponse>) => {
      const response = event.data

      if (response.type === 'result') {
        resolve(response.data)
      } else if (response.type === 'error') {
        reject(new Error(response.message))
      } else if (response.type === 'disposed') {
        resolve(null)
      }

      worker.removeEventListener('message', handler)
    }

    worker.addEventListener('message', handler)
    worker.postMessage(data)
  })
}

function setTheme(theme: Theme): void {
  document.documentElement.setAttribute('data-theme', theme)
}
