import './wasm_exec.js'

const isDev = process.env.NODE_ENV === 'development'

export const WASM_PATH = isDev ? '../bin/coder-mate.wasm' : './coder-mate.wasm'

let wasmModule: WebAssembly.Module | null = null
let wasmInstance: WebAssembly.Instance | null = null
let isInitialized = false
let instanceCount = 0

export async function initWasm(): Promise<void> {
  if (isInitialized) return

  try {
    const response = await fetch(WASM_PATH)
    if (!response.ok) {
      throw new Error(`Failed to load WASM: ${response.status}`)
    }

    const buffer = await response.arrayBuffer()
    wasmModule = await WebAssembly.compile(buffer)

    const go = new (globalThis as any).Go()
    wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject)

    go.run(wasmInstance)
    isInitialized = true
  } catch (error) {
    throw new Error(`WASM initialization failed: ${error}`)
  }
}

export function ensureWasm(): void {
  if (!isInitialized) {
    throw new Error('WASM not initialized. Call createHighlighter first.')
  }
}

export function getInstanceCount(): number {
  return instanceCount
}

export function incrementInstance(): void {
  instanceCount++
}

export function decrementInstance(): void {
  instanceCount--
}

export function resetWasm(): void {
  wasmModule = null
  wasmInstance = null
  isInitialized = false
}
