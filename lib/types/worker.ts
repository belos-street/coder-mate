import type { Language } from './language'
import type { Token } from './token'

export type WorkerMessage =
  | { type: 'highlight'; code: string; lang: Language; mode: 'html' | 'tokens' }
  | { type: 'dispose' }

export type WorkerResponse =
  | { type: 'result'; data: string | Token[] }
  | { type: 'error'; message: string }
  | { type: 'disposed' }
