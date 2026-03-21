export type TokenKind =
  | 'keyword'
  | 'string'
  | 'number'
  | 'identifier'
  | 'operator'
  | 'punctuation'
  | 'comment'
  | 'function'
  | 'variable'
  | 'type'
  | 'constant'
  | 'property'
  | 'tag'
  | 'attribute'
  | 'decorator'
  | 'namespace'
  | 'interpolation'
  | 'regex'
  | 'boolean'
  | 'text'
  | 'unknown'

export type Token = {
  type: TokenKind
  value: string
  line: number
  col: number
}
