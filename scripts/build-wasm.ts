#!/usr/bin/env bun
import { $ } from 'bun'

const projectRoot = `${import.meta.dir}/..`
const wasmOutputPath = `${projectRoot}/bin/coder-mate.wasm`

console.log('Building WASM with TinyGo...')

try {
  await $`cd ${projectRoot}/core && go build -o ${wasmOutputPath} .`
  console.log(`Build complete: ${wasmOutputPath}`)
} catch (error) {
  console.error('Build failed:', error)
  throw error
}
