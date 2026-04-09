// core-fast/language/javascript/rule.ts
var grammarRules = {
  initial: [
    { regex: /(\/\/.*)/y, token: "token-comment", state: "initial" },
    { regex: /(".*?"|'.*?')/y, token: "token-string", state: "initial" },
    {
      regex: /(let|const|var|function|return|if|else|for|while)/y,
      token: "token-keyword",
      state: "initial"
    },
    { regex: /(\d+(\.\d+)?)/y, token: "token-number", state: "initial" },
    {
      regex: /([a-zA-Z_$][a-zA-Z0-9_$]*)/y,
      token: "token-ident",
      state: "initial"
    },
    {
      regex: /([;,.(){}[\]=+\-*\/<>])/y,
      token: "token-punctuation",
      state: "initial"
    },
    { regex: /( +|\t+)/y, token: "token-whitespace", state: "initial" },
    { regex: /(\n)/y, token: "token-newline", state: "initial" }
  ]
};

// core-fast/language/javascript/main.ts
var generateTokens = (code) => {
  const tokens = [[]];
  let currentState = "initial";
  let currentLine = 1;
  let currentCol = 0;
  const codeLength = code.length;
  const rules = grammarRules[currentState];
  let position = 0;
  while (position < codeLength) {
    let matched = false;
    for (const rule of rules) {
      rule.regex.lastIndex = position;
      const match = rule.regex.exec(code);
      debugger;
    }
  }
};
export {
  generateTokens
};
