#!/usr/bin/env bun
import { cpSync, existsSync } from 'fs'
import { execSync } from 'child_process'

const PROJECT_ROOT = import.meta.dir.replace('/scripts', '')
const DIST_DIR = `${PROJECT_ROOT}/dist`
const BIN_DIR = `${PROJECT_ROOT}/bin`
const WASM_PATH = `${BIN_DIR}/coder-mate.wasm`

function buildJs() {
  console.log('Building JS...')
  execSync(
    'bun build ./lib/main.ts --outdir ./dist --format esm --target browser',
    {
      cwd: PROJECT_ROOT
    }
  )
  execSync(
    'bun build ./lib/core/worker.ts --outfile ./dist/worker.js --format esm --target browser',
    {
      cwd: PROJECT_ROOT
    }
  )
  execSync('tsc -p tsconfig.build.json', { cwd: PROJECT_ROOT })

  if (!existsSync(WASM_PATH)) {
    buildWasm()
  } else {
    cpSync(WASM_PATH, `${DIST_DIR}/coder-mate.wasm`)
  }
  console.log('JS build complete')
}

function buildWasm() {
  console.log('Building WASM...')
  execSync('bash ./scripts/build-wasm', {
    cwd: PROJECT_ROOT
  })
  cpSync(WASM_PATH, `${DIST_DIR}/coder-mate.wasm`)
  console.log('WASM build complete')
}

buildJs()
console.log('All builds complete!')
