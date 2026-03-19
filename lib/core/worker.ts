import '../loader/wasm_exec.js'

import type { WorkerMessage, WorkerResponse } from '../types'

let wasmModule: WebAssembly.Module | null = null
let wasmInstance: WebAssembly.Instance | null = null
let isInitialized = false

async function initWasm(): Promise<void> {
  if (isInitialized) return

  const response = await fetch('./coder-mate.wasm')
  if (!response.ok) {
    throw new Error(`Failed to load WASM: ${response.status}`)
  }

  const buffer = await response.arrayBuffer()
  wasmModule = await WebAssembly.compile(buffer)

  const go = new (globalThis as any).Go()
  wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject)

  go.run(wasmInstance)
  isInitialized = true
}

self.onmessage = async (event: MessageEvent<WorkerMessage>) => {
  const { type } = event.data

  try {
    if (type === 'dispose') {
      wasmModule = null
      wasmInstance = null
      isInitialized = false
      const response: WorkerResponse = { type: 'disposed' }
      self.postMessage(response)
      return
    }

    if (type === 'highlight') {
      const { code, lang, mode } = event.data
      await initWasm()
      const result = globalThis.highlightCode(code, lang, mode as any)
      const response: WorkerResponse = { type: 'result', data: result }
      self.postMessage(response)
      return
    }
  } catch (error) {
    const response: WorkerResponse = { type: 'error', message: String(error) }
    self.postMessage(response)
  }
}
