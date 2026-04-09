export type GrammarRule = {
  regex: RegExp
  token: string
  state: 'initial'
}
export const grammarRules: Record<'initial', GrammarRule[]> = {
  // 初始状态：默认匹配
  initial: [
    { regex: /(\/\/.*)/y, token: 'token-comment', state: 'initial' }, // 单行注释
    { regex: /(".*?"|'.*?')/y, token: 'token-string', state: 'initial' }, // 字符串（单/双引号）
    {
      regex: /(let|const|var|function|return|if|else|for|while)/y,
      token: 'token-keyword',
      state: 'initial'
    }, // 关键字
    { regex: /(\d+(\.\d+)?)/y, token: 'token-number', state: 'initial' }, // 数字（整数/小数）
    {
      regex: /([a-zA-Z_$][a-zA-Z0-9_$]*)/y,
      token: 'token-ident',
      state: 'initial'
    }, // 标识符
    {
      regex: /([;,.(){}[\]=+\-*\/<>])/y,
      token: 'token-punctuation',
      state: 'initial'
    }, // 标点符号
    { regex: /( +|\t+)/y, token: 'token-whitespace', state: 'initial' }, // 空格/制表符
    { regex: /(\n)/y, token: 'token-newline', state: 'initial' } // 换行符（仅用于行号处理，不生成token）
  ]
}
