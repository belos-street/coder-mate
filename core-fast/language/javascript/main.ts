import { grammarRules, type GrammarRule } from './rule'

export const generateTokens = (code: string) => {
  const tokens = [[]] // 二维数组，初始第一行
  let currentState: GrammarRule['state'] = 'initial' // 有限状态机初始状态
  let currentLine = 1 // 当前行号（从1开始）
  let currentCol = 0 // 当前列号（从0开始，对应字符起始位置）

  const codeLength = code.length
  const rules = grammarRules[currentState]

  let position = 0 // 当前字符位置（从0开始）
  while (position < codeLength) {
    let matched = false

    for (const rule of rules) {
      rule.regex.lastIndex = position
      const match = rule.regex.exec(code)

      debugger
    }
  }
}
