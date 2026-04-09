/**
 * JavaScript 语法测试用例 - 1000+ 行
 * 包含各种 JS 关键字、语法结构
 */

// ==================== 变量声明 ====================
const PI = 3.141592653589793;
const MAX_SIZE = 100;
const MIN_VALUE = 0;
let counter = 0;
let isActive = true;
let userName = 'John Doe';
let userAge = 30;
var legacyVar = 'legacy';
var globalCounter = 0;

// ==================== 基本数据类型 ====================
const booleanTrue = true;
const booleanFalse = false;
const nullValue = null;
const undefinedValue = undefined;
const bigIntValue = 9007199254740991n;
const binaryNumber = 0b10101010;
const octalNumber = 0o755;
const hexNumber = 0xDEADBEEF;
const floatNumber = 3.14159;
const scientificNumber = 1.23e-10;

// ==================== 字符串 ====================
const singleQuoteString = 'This is a single quote string';
const doubleQuoteString = "This is a double quote string";
const templateString = `This is a template string with ${userName}`;
const multilineString = `Line 1
Line 2
Line 3`;
const escapedString = "Escaped: \n \t \\ \" \'";

// ==================== 数组 ====================
const emptyArray = [];
const numberArray = [1, 2, 3, 4, 5];
const mixedArray = [1, 'two', true, null, undefined, { key: 'value' }];
const nestedArray = [[1, 2], [3, 4], [5, 6]];
const spreadArray = [...numberArray, 6, 7, 8];

// ==================== 对象 ====================
const emptyObject = {};
const simpleObject = {
  name: 'Alice',
  age: 25,
  isActive: true
};
const nestedObject = {
  user: {
    name: 'Bob',
    address: {
      city: 'New York',
      country: 'USA'
    }
  },
  settings: {
    theme: 'dark',
    language: 'en'
  }
};
const computedProperty = {
  ['computed' + 'Key']: 'computed value',
  [1 + 2]: 'three'
};

// ==================== 函数声明 ====================
function regularFunction() {
  return 'regular function';
}

function functionWithParams(a, b, c) {
  return a + b + c;
}

function functionWithDefaultParams(a, b = 10, c = 20) {
  return a + b + c;
}

function functionWithRestParams(...args) {
  return args.reduce((sum, val) => sum + val, 0);
}

function functionWithDestructuring({ name, age }, [x, y]) {
  return `${name} is ${age}, coords: ${x}, ${y}`;
}

// ==================== 箭头函数 ====================
const arrowFunction = () => 'arrow function';
const arrowWithParams = (x, y) => x + y;
const arrowWithBody = (x, y) => {
  const sum = x + y;
  return sum * 2;
};
const arrowWithSingleParam = x => x * 2;
const arrowReturningObject = (x, y) => ({ x, y, sum: x + y });

// ==================== 立即执行函数 ====================
const iifeResult = (function() {
  return 'IIFE result';
})();

const iifeWithParams = (function(a, b) {
  return a * b;
})(5, 10);

// ==================== 类定义 ====================
class Animal {
  static kingdom = 'Animalia';
  
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }
  
  speak() {
    return `${this.name} makes a sound`;
  }
  
  get info() {
    return `${this.name} is ${this.age} years old`;
  }
  
  set age(value) {
    if (value > 0) {
      this._age = value;
    }
  }
  
  static create(name, age) {
    return new Animal(name, age);
  }
}

class Dog extends Animal {
  constructor(name, age, breed) {
    super(name, age);
    this.breed = breed;
  }
  
  speak() {
    return `${this.name} barks!`;
  }
  
  fetch() {
    return `${this.name} fetches the ball`;
  }
}

class Cat extends Animal {
  #secretValue = 'meow';
  
  constructor(name, age, color) {
    super(name, age);
    this.color = color;
  }
  
  speak() {
    return `${this.name} meows!`;
  }
  
  #privateMethod() {
    return this.#secretValue;
  }
  
  getSecret() {
    return this.#privateMethod();
  }
}

// ==================== 控制流 ====================
function ifElseExample(value) {
  if (value < 0) {
    return 'negative';
  } else if (value === 0) {
    return 'zero';
  } else if (value > 0 && value < 10) {
    return 'small positive';
  } else {
    return 'large positive';
  }
}

function switchExample(value) {
  switch (value) {
    case 1:
      return 'one';
    case 2:
      return 'two';
    case 3:
    case 4:
      return 'three or four';
    default:
      return 'unknown';
  }
}

function forLoopExample() {
  const result = [];
  for (let i = 0; i < 10; i++) {
    result.push(i);
  }
  return result;
}

function forInLoopExample(obj) {
  const keys = [];
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      keys.push(key);
    }
  }
  return keys;
}

function forOfLoopExample(arr) {
  const result = [];
  for (const item of arr) {
    result.push(item * 2);
  }
  return result;
}

function whileLoopExample() {
  let count = 0;
  while (count < 5) {
    count++;
  }
  return count;
}

function doWhileLoopExample() {
  let count = 0;
  do {
    count++;
  } while (count < 5);
  return count;
}

// ==================== 异常处理 ====================
function tryCatchExample() {
  try {
    throw new Error('Custom error');
  } catch (error) {
    console.error('Caught:', error.message);
    return 'error handled';
  } finally {
    console.log('Cleanup');
  }
}

function tryFinallyExample() {
  try {
    return 'try result';
  } finally {
    console.log('Always executed');
  }
}

async function asyncTryCatch() {
  try {
    const response = await fetch('/api/data');
    return await response.json();
  } catch (error) {
    console.error('Fetch failed:', error);
    throw error;
  }
}

// ==================== 异步函数 ====================
async function asyncFunction() {
  return 'async result';
}

async function asyncWithAwait() {
  const result = await asyncFunction();
  return result.toUpperCase();
}

async function asyncWithPromise() {
  const promise = new Promise((resolve, reject) => {
    setTimeout(() => resolve('resolved'), 1000);
  });
  return await promise;
}

async function parallelAsync() {
  const [result1, result2] = await Promise.all([
    asyncFunction(),
    asyncWithAwait()
  ]);
  return { result1, result2 };
}

function* generatorFunction() {
  yield 1;
  yield 2;
  yield 3;
}

async function* asyncGeneratorFunction() {
  yield await Promise.resolve(1);
  yield await Promise.resolve(2);
  yield await Promise.resolve(3);
}

// ==================== Promise ====================
const resolvedPromise = Promise.resolve('resolved value');
const rejectedPromise = Promise.reject('rejected reason');
const pendingPromise = new Promise((resolve, reject) => {
  setTimeout(() => resolve('delayed'), 1000);
});

function promiseChain() {
  return Promise.resolve(1)
    .then(value => value * 2)
    .then(value => value + 10)
    .then(value => `Result: ${value}`)
    .catch(error => `Error: ${error}`)
    .finally(() => console.log('Done'));
}

async function promiseAllExample() {
  const results = await Promise.all([
    Promise.resolve(1),
    Promise.resolve(2),
    Promise.resolve(3)
  ]);
  return results;
}

async function promiseAllSettledExample() {
  const results = await Promise.allSettled([
    Promise.resolve(1),
    Promise.reject('error'),
    Promise.resolve(3)
  ]);
  return results;
}

async function promiseRaceExample() {
  const result = await Promise.race([
    new Promise(resolve => setTimeout(() => resolve('slow'), 1000)),
    new Promise(resolve => setTimeout(() => resolve('fast'), 500))
  ]);
  return result;
}

async function promiseAnyExample() {
  const result = await Promise.any([
    Promise.reject('error1'),
    Promise.resolve('success'),
    Promise.reject('error2')
  ]);
  return result;
}

// ==================== 解构赋值 ====================
const [first, second, ...rest] = [1, 2, 3, 4, 5];
const { name, age, ...others } = { name: 'Alice', age: 25, city: 'NYC', country: 'USA' };
const [{ x, y }, [a, b]] = [{ x: 1, y: 2 }, [10, 20]];

function destructuringInParams({ name, age }, [x, y, z]) {
  return { name, age, coords: { x, y, z } };
}

// ==================== 展开运算符 ====================
const spreadInArray = [...[1, 2, 3], 4, 5];
const spreadInObject = { ...{ a: 1, b: 2 }, c: 3 };
const spreadInFunction = Math.max(...[1, 2, 3, 4, 5]);

// ==================== 运算符 ====================
function arithmeticOperators() {
  const a = 10;
  const b = 3;
  return {
    addition: a + b,
    subtraction: a - b,
    multiplication: a * b,
    division: a / b,
    modulo: a % b,
    exponentiation: a ** b,
    increment: ++a,
    decrement: --b
  };
}

function comparisonOperators() {
  return {
    equal: 5 == '5',
    strictEqual: 5 === '5',
    notEqual: 5 != '5',
    strictNotEqual: 5 !== '5',
    greaterThan: 10 > 5,
    greaterThanOrEqual: 10 >= 10,
    lessThan: 5 < 10,
    lessThanOrEqual: 5 <= 5
  };
}

function logicalOperators() {
  return {
    and: true && false,
    or: true || false,
    not: !true,
    nullish: null ?? 'default',
    nullishWithUndefined: undefined ?? 'default',
    optionalChaining: nestedObject?.user?.address?.city
  };
}

function bitwiseOperators() {
  return {
    and: 5 & 3,
    or: 5 | 3,
    xor: 5 ^ 3,
    not: ~5,
    leftShift: 5 << 1,
    rightShift: 5 >> 1,
    unsignedRightShift: 5 >>> 1
  };
}

function assignmentOperators() {
  let a = 10;
  a += 5;
  a -= 3;
  a *= 2;
  a /= 4;
  a %= 3;
  a **= 2;
  a &= 15;
  a |= 8;
  a ^= 4;
  a <<= 1;
  a >>= 1;
  return a;
}

function ternaryOperator(condition) {
  return condition ? 'truthy' : 'falsy';
}

// ==================== 模块导入导出（注释掉，因为需要在模块环境中） ====================
// import { something } from './module.js';
// import * as utils from './utils.js';
// import defaultExport from './default.js';
// export const exportedValue = 42;
// export default function() { return 'default'; }

// ==================== Map 和 Set ====================
const mapExample = new Map();
mapExample.set('key1', 'value1');
mapExample.set('key2', 'value2');
mapExample.set(1, 'numeric key');
mapExample.set(true, 'boolean key');

const setExample = new Set([1, 2, 3, 3, 4, 4, 5]);
const weakMapExample = new WeakMap();
const weakSetExample = new WeakSet();

// ==================== Symbol ====================
const symbol1 = Symbol('description');
const symbol2 = Symbol('description');
const uniqueSymbol = Symbol.for('unique');
const wellKnownSymbol = Symbol.iterator;

const objectWithSymbol = {
  [symbol1]: 'symbol value',
  regularKey: 'regular value'
};

// ==================== Proxy 和 Reflect ====================
const proxyExample = new Proxy({}, {
  get(target, prop) {
    return prop in target ? target[prop] : 'default';
  },
  set(target, prop, value) {
    target[prop] = value;
    return true;
  }
});

// ==================== 正则表达式 ====================
const regex1 = /pattern/;
const regex2 = /pattern/gi;
const regex3 = new RegExp('pattern', 'gi');

function regexMethods() {
  const str = 'Hello World 123';
  return {
    test: /\d+/.test(str),
    match: str.match(/\d+/g),
    matchAll: [...str.matchAll(/\d/g)],
    search: str.search(/\d+/),
    replace: str.replace(/\d+/, 'NUM'),
    replaceAll: str.replaceAll(/\d/g, 'X'),
    split: str.split(/\s+/)
  };
}

// ==================== 日期 ====================
const dateExample = new Date();
const dateFromTimestamp = new Date(1609459200000);
const dateFromString = new Date('2021-01-01T00:00:00Z');
const dateFromParts = new Date(2021, 0, 1, 0, 0, 0, 0);

// ==================== JSON ====================
const jsonString = '{"name": "Alice", "age": 25}';
const parsedJson = JSON.parse(jsonString);
const stringifiedJson = JSON.stringify(parsedJson, null, 2);

// ==================== 数组方法 ====================
function arrayMethods() {
  const arr = [1, 2, 3, 4, 5];
  
  return {
    map: arr.map(x => x * 2),
    filter: arr.filter(x => x > 2),
    reduce: arr.reduce((sum, x) => sum + x, 0),
    find: arr.find(x => x > 2),
    findIndex: arr.findIndex(x => x > 2),
    some: arr.some(x => x > 3),
    every: arr.every(x => x > 0),
    includes: arr.includes(3),
    indexOf: arr.indexOf(3),
    lastIndexOf: arr.lastIndexOf(3),
    slice: arr.slice(1, 3),
    splice: [...arr].splice(1, 2),
    concat: arr.concat([6, 7]),
    join: arr.join('-'),
    reverse: [...arr].reverse(),
    sort: [...arr].sort((a, b) => b - a),
    flat: [[1, 2], [3, 4]].flat(),
    flatMap: arr.flatMap(x => [x, x * 2]),
    fill: new Array(5).fill(0),
    copyWithin: [1, 2, 3, 4, 5].copyWithin(0, 3),
    at: arr.at(-1)
  };
}

// ==================== 对象方法 ====================
function objectMethods() {
  const obj = { a: 1, b: 2, c: 3 };
  
  return {
    keys: Object.keys(obj),
    values: Object.values(obj),
    entries: Object.entries(obj),
    fromEntries: Object.fromEntries([['x', 1], ['y', 2]]),
    assign: Object.assign({}, obj, { d: 4 }),
    spread: { ...obj, d: 4 },
    freeze: Object.freeze({ frozen: true }),
    seal: Object.seal({ sealed: true }),
    preventExtensions: Object.preventExtensions({}),
    hasOwnProperty: obj.hasOwnProperty('a'),
    propertyIsEnumerable: obj.propertyIsEnumerable('a'),
    isPrototypeOf: Object.prototype.isPrototypeOf(obj),
    toString: obj.toString(),
    valueOf: obj.valueOf()
  };
}

// ==================== 字符串方法 ====================
function stringMethods() {
  const str = 'Hello, World!';
  
  return {
    charAt: str.charAt(0),
    charCodeAt: str.charCodeAt(0),
    codePointAt: str.codePointAt(0),
    concat: str.concat('!', '!'),
    endsWith: str.endsWith('!'),
    includes: str.includes('World'),
    indexOf: str.indexOf('o'),
    lastIndexOf: str.lastIndexOf('o'),
    localeCompare: 'a'.localeCompare('b'),
    match: str.match(/[A-Z]/g),
    matchAll: [...str.matchAll(/[A-Z]/g)],
    normalize: 'café'.normalize('NFD'),
    padEnd: str.padEnd(20, '.'),
    padStart: str.padStart(20, '.'),
    repeat: 'ab'.repeat(3),
    replace: str.replace('World', 'JS'),
    replaceAll: str.replaceAll('o', '0'),
    search: str.search(/World/),
    slice: str.slice(0, 5),
    split: str.split(', '),
    startsWith: str.startsWith('Hello'),
    substring: str.substring(0, 5),
    toLowerCase: str.toLowerCase(),
    toUpperCase: str.toUpperCase(),
    trim: '  hello  '.trim(),
    trimStart: '  hello  '.trimStart(),
    trimEnd: '  hello  '.trimEnd(),
    at: str.at(-1)
  };
}

// ==================== 数学方法 ====================
function mathMethods() {
  return {
    abs: Math.abs(-5),
    ceil: Math.ceil(4.3),
    floor: Math.floor(4.7),
    round: Math.round(4.5),
    trunc: Math.trunc(4.7),
    sign: Math.sign(-5),
    sqrt: Math.sqrt(16),
    cbrt: Math.cbrt(27),
    pow: Math.pow(2, 3),
    exp: Math.exp(1),
    log: Math.log(Math.E),
    log2: Math.log2(8),
    log10: Math.log10(100),
    sin: Math.sin(0),
    cos: Math.cos(0),
    tan: Math.tan(0),
    asin: Math.asin(0),
    acos: Math.acos(1),
    atan: Math.atan(0),
    atan2: Math.atan2(0, 1),
    min: Math.min(1, 2, 3),
    max: Math.max(1, 2, 3),
    random: Math.random(),
    hypot: Math.hypot(3, 4),
    clz32: Math.clz32(1),
    fround: Math.fround(1.5),
    imul: Math.imul(2, 3)
  };
}

// ==================== 类型检查 ====================
function typeChecking() {
  return {
    typeofString: typeof 'hello',
    typeofNumber: typeof 42,
    typeofBoolean: typeof true,
    typeofUndefined: typeof undefined,
    typeofObject: typeof {},
    typeofFunction: typeof function() {},
    typeofSymbol: typeof Symbol(),
    typeofBigInt: typeof 1n,
    instanceofArray: [] instanceof Array,
    instanceofObject: {} instanceof Object,
    isArray: Array.isArray([]),
    isNaN: Number.isNaN(NaN),
    isFinite: Number.isFinite(100),
    isInteger: Number.isInteger(42),
    isSafeInteger: Number.isSafeInteger(42)
  };
}

// ==================== 类型转换 ====================
function typeConversion() {
  return {
    stringToNumber: Number('42'),
    numberToString: String(42),
    toBoolean: Boolean(1),
    toInteger: parseInt('42px', 10),
    toFloat: parseFloat('3.14'),
    toString: (42).toString(),
    toFixed: (3.14159).toFixed(2),
    toPrecision: (3.14159).toPrecision(3),
    toExponential: (1000).toExponential(2)
  };
}

// ==================== 闭包 ====================
function closureExample() {
  let counter = 0;
  
  return {
    increment: function() {
      counter++;
      return counter;
    },
    decrement: function() {
      counter--;
      return counter;
    },
    getCount: function() {
      return counter;
    }
  };
}

// ==================== 高阶函数 ====================
function higherOrderFunction() {
  const compose = (...fns) => x => fns.reduceRight((v, f) => f(v), x);
  const pipe = (...fns) => x => fns.reduce((v, f) => f(v), x);
  const curry = fn => (...args) => 
    args.length >= fn.length ? fn(...args) : curry(fn.bind(null, ...args));
  const memoize = fn => {
    const cache = new Map();
    return (...args) => {
      const key = JSON.stringify(args);
      if (!cache.has(key)) {
        cache.set(key, fn(...args));
      }
      return cache.get(key);
    };
  };
  
  return { compose, pipe, curry, memoize };
}

// ==================== 递归 ====================
function factorial(n) {
  if (n <= 1) return 1;
  return n * factorial(n - 1);
}

function fibonacci(n) {
  if (n <= 1) return n;
  return fibonacci(n - 1) + fibonacci(n - 2);
}

function fibonacciMemo(n, memo = {}) {
  if (n in memo) return memo[n];
  if (n <= 1) return n;
  memo[n] = fibonacciMemo(n - 1, memo) + fibonacciMemo(n - 2, memo);
  return memo[n];
}

function deepClone(obj, seen = new WeakMap()) {
  if (obj === null || typeof obj !== 'object') return obj;
  if (seen.has(obj)) return seen.get(obj);
  
  const clone = Array.isArray(obj) ? [] : {};
  seen.set(obj, clone);
  
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      clone[key] = deepClone(obj[key], seen);
    }
  }
  
  return clone;
}

// ==================== 设计模式 ====================

// 单例模式
const Singleton = (function() {
  let instance;
  
  function createInstance() {
    return { name: 'Singleton Instance' };
  }
  
  return {
    getInstance: function() {
      if (!instance) {
        instance = createInstance();
      }
      return instance;
    }
  };
})();

// 工厂模式
class AnimalFactory {
  static create(type) {
    switch (type) {
      case 'dog':
        return new Dog('Buddy', 3, 'Labrador');
      case 'cat':
        return new Cat('Whiskers', 2, 'Orange');
      default:
        return new Animal('Unknown', 0);
    }
  }
}

// 观察者模式
class EventEmitter {
  constructor() {
    this.events = {};
  }
  
  on(event, callback) {
    if (!this.events[event]) {
      this.events[event] = [];
    }
    this.events[event].push(callback);
  }
  
  emit(event, ...args) {
    if (this.events[event]) {
      this.events[event].forEach(callback => callback(...args));
    }
  }
  
  off(event, callback) {
    if (this.events[event]) {
      this.events[event] = this.events[event].filter(cb => cb !== callback);
    }
  }
}

// 发布订阅模式
class PubSub {
  constructor() {
    this.subscribers = {};
  }
  
  subscribe(topic, callback) {
    if (!this.subscribers[topic]) {
      this.subscribers[topic] = [];
    }
    this.subscribers[topic].push(callback);
  }
  
  publish(topic, ...args) {
    if (this.subscribers[topic]) {
      this.subscribers[topic].forEach(callback => callback(...args));
    }
  }
  
  unsubscribe(topic, callback) {
    if (this.subscribers[topic]) {
      this.subscribers[topic] = this.subscribers[topic].filter(cb => cb !== callback);
    }
  }
}

// 策略模式
class PaymentStrategy {
  static strategies = {
    credit: (amount) => `Paid ${amount} via Credit Card`,
    paypal: (amount) => `Paid ${amount} via PayPal`,
    crypto: (amount) => `Paid ${amount} via Crypto`
  };
  
  static pay(type, amount) {
    return this.strategies[type](amount);
  }
}

// 命令模式
class Command {
  constructor(execute, undo) {
    this.execute = execute;
    this.undo = undo;
  }
}

class CommandManager {
  constructor() {
    this.history = [];
  }
  
  execute(command) {
    command.execute();
    this.history.push(command);
  }
  
  undo() {
    const command = this.history.pop();
    if (command) {
      command.undo();
    }
  }
}

// ==================== DOM 操作（模拟） ====================
const domOperations = {
  createElement: (tag) => document.createElement(tag),
  getElementById: (id) => document.getElementById(id),
  querySelector: (selector) => document.querySelector(selector),
  querySelectorAll: (selector) => document.querySelectorAll(selector),
  addEventListener: (el, event, handler) => el.addEventListener(event, handler),
  removeEventListener: (el, event, handler) => el.removeEventListener(event, handler),
  appendChild: (parent, child) => parent.appendChild(child),
  removeChild: (parent, child) => parent.removeChild(child),
  setAttribute: (el, name, value) => el.setAttribute(name, value),
  getAttribute: (el, name) => el.getAttribute(name),
  classList: {
    add: (el, className) => el.classList.add(className),
    remove: (el, className) => el.classList.remove(className),
    toggle: (el, className) => el.classList.toggle(className),
    contains: (el, className) => el.classList.contains(className)
  },
  style: (el, styles) => Object.assign(el.style, styles)
};

// ==================== 事件处理 ====================
const eventHandlers = {
  onClick: (callback) => document.addEventListener('click', callback),
  onKeydown: (callback) => document.addEventListener('keydown', callback),
  onKeyup: (callback) => document.addEventListener('keyup', callback),
  onMouseover: (callback) => document.addEventListener('mouseover', callback),
  onMouseout: (callback) => document.addEventListener('mouseout', callback),
  onSubmit: (form, callback) => form.addEventListener('submit', callback),
  onChange: (input, callback) => input.addEventListener('change', callback),
  onInput: (input, callback) => input.addEventListener('input', callback),
  onFocus: (el, callback) => el.addEventListener('focus', callback),
  onBlur: (el, callback) => el.addEventListener('blur', callback),
  onScroll: (callback) => window.addEventListener('scroll', callback),
  onResize: (callback) => window.addEventListener('resize', callback),
  onLoad: (callback) => window.addEventListener('load', callback),
  onDOMContentLoaded: (callback) => document.addEventListener('DOMContentLoaded', callback)
};

// ==================== 防抖和节流 ====================
function debounce(fn, delay) {
  let timer;
  return function(...args) {
    clearTimeout(timer);
    timer = setTimeout(() => fn.apply(this, args), delay);
  };
}

function throttle(fn, limit) {
  let inThrottle;
  return function(...args) {
    if (!inThrottle) {
      fn.apply(this, args);
      inThrottle = true;
      setTimeout(() => inThrottle = false, limit);
    }
  };
}

// ==================== 深拷贝和浅拷贝 ====================
function shallowCopy(obj) {
  if (Array.isArray(obj)) {
    return [...obj];
  }
  return { ...obj };
}

function deepCopy(obj, hash = new WeakMap()) {
  if (obj === null || typeof obj !== 'object') return obj;
  if (hash.has(obj)) return hash.get(obj);
  
  const result = Array.isArray(obj) ? [] : {};
  hash.set(obj, result);
  
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      result[key] = deepCopy(obj[key], hash);
    }
  }
  
  return result;
}

// ==================== 数据结构 ====================
class Stack {
  constructor() {
    this.items = [];
  }
  
  push(item) {
    this.items.push(item);
  }
  
  pop() {
    return this.items.pop();
  }
  
  peek() {
    return this.items[this.items.length - 1];
  }
  
  isEmpty() {
    return this.items.length === 0;
  }
  
  size() {
    return this.items.length;
  }
  
  clear() {
    this.items = [];
  }
}

class Queue {
  constructor() {
    this.items = [];
  }
  
  enqueue(item) {
    this.items.push(item);
  }
  
  dequeue() {
    return this.items.shift();
  }
  
  front() {
    return this.items[0];
  }
  
  isEmpty() {
    return this.items.length === 0;
  }
  
  size() {
    return this.items.length;
  }
  
  clear() {
    this.items = [];
  }
}

class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }
  
  append(value) {
    const node = { value, next: null };
    if (!this.head) {
      this.head = node;
      this.tail = node;
    } else {
      this.tail.next = node;
      this.tail = node;
    }
    this.length++;
  }
  
  prepend(value) {
    const node = { value, next: this.head };
    this.head = node;
    if (!this.tail) {
      this.tail = node;
    }
    this.length++;
  }
  
  remove(value) {
    if (!this.head) return;
    if (this.head.value === value) {
      this.head = this.head.next;
      this.length--;
      return;
    }
    let current = this.head;
    while (current.next) {
      if (current.next.value === value) {
        current.next = current.next.next;
        this.length--;
        return;
      }
      current = current.next;
    }
  }
  
  find(value) {
    let current = this.head;
    while (current) {
      if (current.value === value) {
        return current;
      }
      current = current.next;
    }
    return null;
  }
  
  toArray() {
    const result = [];
    let current = this.head;
    while (current) {
      result.push(current.value);
      current = current.next;
    }
    return result;
  }
}

class TreeNode {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

class BinaryTree {
  constructor() {
    this.root = null;
  }
  
  insert(value) {
    const node = new TreeNode(value);
    if (!this.root) {
      this.root = node;
      return;
    }
    
    let current = this.root;
    while (true) {
      if (value < current.value) {
        if (!current.left) {
          current.left = node;
          break;
        }
        current = current.left;
      } else {
        if (!current.right) {
          current.right = node;
          break;
        }
        current = current.right;
      }
    }
  }
  
  search(value) {
    let current = this.root;
    while (current) {
      if (value === current.value) {
        return current;
      }
      if (value < current.value) {
        current = current.left;
      } else {
        current = current.right;
      }
    }
    return null;
  }
  
  inOrderTraversal(node = this.root, result = []) {
    if (node) {
      this.inOrderTraversal(node.left, result);
      result.push(node.value);
      this.inOrderTraversal(node.right, result);
    }
    return result;
  }
  
  preOrderTraversal(node = this.root, result = []) {
    if (node) {
      result.push(node.value);
      this.preOrderTraversal(node.left, result);
      this.preOrderTraversal(node.right, result);
    }
    return result;
  }
  
  postOrderTraversal(node = this.root, result = []) {
    if (node) {
      this.postOrderTraversal(node.left, result);
      this.postOrderTraversal(node.right, result);
      result.push(node.value);
    }
    return result;
  }
}

// ==================== 排序算法 ====================
function bubbleSort(arr) {
  const result = [...arr];
  const n = result.length;
  for (let i = 0; i < n - 1; i++) {
    for (let j = 0; j < n - i - 1; j++) {
      if (result[j] > result[j + 1]) {
        [result[j], result[j + 1]] = [result[j + 1], result[j]];
      }
    }
  }
  return result;
}

function quickSort(arr) {
  if (arr.length <= 1) return arr;
  const pivot = arr[Math.floor(arr.length / 2)];
  const left = arr.filter(x => x < pivot);
  const middle = arr.filter(x => x === pivot);
  const right = arr.filter(x => x > pivot);
  return [...quickSort(left), ...middle, ...quickSort(right)];
}

function mergeSort(arr) {
  if (arr.length <= 1) return arr;
  const mid = Math.floor(arr.length / 2);
  const left = mergeSort(arr.slice(0, mid));
  const right = mergeSort(arr.slice(mid));
  return merge(left, right);
}

function merge(left, right) {
  const result = [];
  let i = 0, j = 0;
  while (i < left.length && j < right.length) {
    if (left[i] <= right[j]) {
      result.push(left[i++]);
    } else {
      result.push(right[j++]);
    }
  }
  return [...result, ...left.slice(i), ...right.slice(j)];
}

function insertionSort(arr) {
  const result = [...arr];
  for (let i = 1; i < result.length; i++) {
    const key = result[i];
    let j = i - 1;
    while (j >= 0 && result[j] > key) {
      result[j + 1] = result[j];
      j--;
    }
    result[j + 1] = key;
  }
  return result;
}

function selectionSort(arr) {
  const result = [...arr];
  for (let i = 0; i < result.length - 1; i++) {
    let minIndex = i;
    for (let j = i + 1; j < result.length; j++) {
      if (result[j] < result[minIndex]) {
        minIndex = j;
      }
    }
    [result[i], result[minIndex]] = [result[minIndex], result[i]];
  }
  return result;
}

// ==================== 搜索算法 ====================
function linearSearch(arr, target) {
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === target) return i;
  }
  return -1;
}

function binarySearch(arr, target) {
  let left = 0;
  let right = arr.length - 1;
  while (left <= right) {
    const mid = Math.floor((left + right) / 2);
    if (arr[mid] === target) return mid;
    if (arr[mid] < target) left = mid + 1;
    else right = mid - 1;
  }
  return -1;
}

function binarySearchRecursive(arr, target, left = 0, right = arr.length - 1) {
  if (left > right) return -1;
  const mid = Math.floor((left + right) / 2);
  if (arr[mid] === target) return mid;
  if (arr[mid] < target) return binarySearchRecursive(arr, target, mid + 1, right);
  return binarySearchRecursive(arr, target, left, mid - 1);
}

// ==================== 工具函数 ====================
const utils = {
  sleep: (ms) => new Promise(resolve => setTimeout(resolve, ms)),
  range: (start, end, step = 1) => {
    const result = [];
    for (let i = start; i < end; i += step) {
      result.push(i);
    }
    return result;
  },
  chunk: (arr, size) => {
    const result = [];
    for (let i = 0; i < arr.length; i += size) {
      result.push(arr.slice(i, i + size));
    }
    return result;
  },
  flatten: (arr) => arr.flat(Infinity),
  unique: (arr) => [...new Set(arr)],
  shuffle: (arr) => {
    const result = [...arr];
    for (let i = result.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [result[i], result[j]] = [result[j], result[i]];
    }
    return result;
  },
  sample: (arr, n = 1) => {
    const shuffled = utils.shuffle(arr);
    return n === 1 ? shuffled[0] : shuffled.slice(0, n);
  },
  groupBy: (arr, key) => {
    return arr.reduce((groups, item) => {
      const group = typeof key === 'function' ? key(item) : item[key];
      if (!groups[group]) groups[group] = [];
      groups[group].push(item);
      return groups;
    }, {});
  },
  keyBy: (arr, key) => {
    return arr.reduce((obj, item) => {
      const k = typeof key === 'function' ? key(item) : item[key];
      obj[k] = item;
      return obj;
    }, {});
  },
  sortBy: (arr, key) => {
    return [...arr].sort((a, b) => {
      const aVal = typeof key === 'function' ? key(a) : a[key];
      const bVal = typeof key === 'function' ? key(b) : b[key];
      return aVal > bVal ? 1 : aVal < bVal ? -1 : 0;
    });
  },
  partition: (arr, predicate) => {
    return arr.reduce(([pass, fail], item) => {
      return predicate(item) ? [[...pass, item], fail] : [pass, [...fail, item]];
    }, [[], []]);
  },
  zip: (...arrays) => {
    const length = Math.min(...arrays.map(arr => arr.length));
    return Array.from({ length }, (_, i) => arrays.map(arr => arr[i]));
  },
  unzip: (arr) => {
    return arr[0].map((_, i) => arr.map(row => row[i]));
  },
  pick: (obj, keys) => {
    return keys.reduce((result, key) => {
      if (key in obj) result[key] = obj[key];
      return result;
    }, {});
  },
  omit: (obj, keys) => {
    return Object.keys(obj).reduce((result, key) => {
      if (!keys.includes(key)) result[key] = obj[key];
      return result;
    }, {});
  },
  merge: (...objects) => {
    return objects.reduce((result, obj) => {
      Object.keys(obj).forEach(key => {
        if (typeof obj[key] === 'object' && obj[key] !== null && !Array.isArray(obj[key])) {
          result[key] = utils.merge(result[key] || {}, obj[key]);
        } else {
          result[key] = obj[key];
        }
      });
      return result;
    }, {});
  }
};

// ==================== 函数式编程 ====================
const functional = {
  identity: (x) => x,
  always: (x) => () => x,
  compose: (...fns) => (x) => fns.reduceRight((v, f) => f(v), x),
  pipe: (...fns) => (x) => fns.reduce((v, f) => f(v), x),
  curry: (fn) => {
    const curried = (...args) => {
      if (args.length >= fn.length) {
        return fn(...args);
      }
      return (...more) => curried(...args, ...more);
    };
    return curried;
  },
  partial: (fn, ...presetArgs) => (...laterArgs) => fn(...presetArgs, ...laterArgs),
  flip: (fn) => (...args) => fn(...args.reverse()),
  negate: (fn) => (...args) => !fn(...args),
  once: (fn) => {
    let called = false;
    let result;
    return (...args) => {
      if (!called) {
        called = true;
        result = fn(...args);
      }
      return result;
    };
  },
  memoize: (fn) => {
    const cache = new Map();
    return (...args) => {
      const key = JSON.stringify(args);
      if (!cache.has(key)) {
        cache.set(key, fn(...args));
      }
      return cache.get(key);
    };
  }
};

// ==================== 异步工具 ====================
const asyncUtils = {
  parallel: async (tasks) => {
    return Promise.all(tasks.map(task => task()));
  },
  series: async (tasks) => {
    const results = [];
    for (const task of tasks) {
      results.push(await task());
    }
    return results;
  },
  waterfall: async (tasks, initial) => {
    let result = initial;
    for (const task of tasks) {
      result = await task(result);
    }
    return result;
  },
  retry: async (fn, retries = 3, delay = 1000) => {
    for (let i = 0; i < retries; i++) {
      try {
        return await fn();
      } catch (error) {
        if (i === retries - 1) throw error;
        await new Promise(resolve => setTimeout(resolve, delay));
      }
    }
  },
  timeout: (promise, ms) => {
    return Promise.race([
      promise,
      new Promise((_, reject) => 
        setTimeout(() => reject(new Error('Timeout')), ms)
      )
    ]);
  },
  delay: (ms) => new Promise(resolve => setTimeout(resolve, ms)),
  map: async (arr, fn, concurrency = Infinity) => {
    const results = [];
    const executing = [];
    for (const item of arr) {
      const promise = Promise.resolve().then(() => fn(item));
      results.push(promise);
      if (executing.length >= concurrency) {
        await Promise.race(executing);
      }
      executing.push(promise);
      promise.finally(() => {
        executing.splice(executing.indexOf(promise), 1);
      });
    }
    return Promise.all(results);
  }
};

// ==================== 验证函数 ====================
const validators = {
  isEmail: (str) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(str),
  isUrl: (str) => /^https?:\/\/[^\s]+$/.test(str),
  isPhone: (str) => /^\+?[\d\s-]{10,}$/.test(str),
  isNumeric: (str) => /^\d+$/.test(str),
  isAlpha: (str) => /^[a-zA-Z]+$/.test(str),
  isAlphanumeric: (str) => /^[a-zA-Z0-9]+$/.test(str),
  isDate: (str) => !isNaN(Date.parse(str)),
  isJSON: (str) => {
    try {
      JSON.parse(str);
      return true;
    } catch {
      return false;
    }
  },
  isEmpty: (value) => {
    if (value === null || value === undefined) return true;
    if (typeof value === 'string') return value.trim() === '';
    if (Array.isArray(value)) return value.length === 0;
    if (typeof value === 'object') return Object.keys(value).length === 0;
    return false;
  },
  isPlainObject: (value) => {
    return typeof value === 'object' && value !== null && 
           Object.prototype.toString.call(value) === '[object Object]';
  },
  hasLength: (value, min, max) => {
    const len = value?.length ?? 0;
    return len >= min && len <= max;
  },
  inRange: (value, min, max) => {
    return value >= min && value <= max;
  }
};

// ==================== 格式化函数 ====================
const formatters = {
  currency: (num, currency = 'USD', locale = 'en-US') => {
    return new Intl.NumberFormat(locale, { style: 'currency', currency }).format(num);
  },
  number: (num, locale = 'en-US') => {
    return new Intl.NumberFormat(locale).format(num);
  },
  date: (date, locale = 'en-US', options = {}) => {
    return new Intl.DateTimeFormat(locale, options).format(date);
  },
  relativeTime: (value, unit, locale = 'en-US') => {
    return new Intl.RelativeTimeFormat(locale).format(value, unit);
  },
  plural: (count, singular, plural) => {
    return count === 1 ? singular : plural;
  },
  truncate: (str, length, suffix = '...') => {
    return str.length > length ? str.slice(0, length) + suffix : str;
  },
  capitalize: (str) => {
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
  },
  titleCase: (str) => {
    return str.replace(/\b\w/g, char => char.toUpperCase());
  },
  camelCase: (str) => {
    return str.replace(/[-_\s]+(.)?/g, (_, c) => c ? c.toUpperCase() : '');
  },
  kebabCase: (str) => {
    return str.replace(/([a-z])([A-Z])/g, '$1-$2').toLowerCase();
  },
  snakeCase: (str) => {
    return str.replace(/([a-z])([A-Z])/g, '$1_$2').toLowerCase();
  }
};

// ==================== 注释测试 ====================

// 单行注释
// 另一个单行注释

/*
 * 多行注释
 * 第二行
 * 第三行
 */

/**
 * JSDoc 注释
 * @param {string} name - 名称
 * @param {number} age - 年龄
 * @returns {object} 用户对象
 */
function createJSDocExample(name, age) {
  return { name, age };
}

/* 单行多行注释 */

// ==================== 特殊语法 ====================

// 标签语句
function labelExample() {
  outer: for (let i = 0; i < 10; i++) {
    for (let j = 0; j < 10; j++) {
      if (i === 5 && j === 5) {
        break outer;
      }
    }
  }
}

// with 语句（不推荐使用）
function withExample() {
  const obj = { a: 1, b: 2 };
  // with (obj) {
  //   console.log(a, b);
  // }
}

// debugger 语句
function debuggerExample() {
  // debugger;
  return 'debugger example';
}

// void 运算符
const voidResult = void 0;

// in 运算符
const hasProperty = 'name' in { name: 'Alice' };

// instanceof 运算符
const isArrayInstance = [] instanceof Array;

// delete 运算符
function deleteExample() {
  const obj = { a: 1, b: 2 };
  delete obj.a;
  return obj;
}

// typeof 运算符
const typeOfString = typeof 'hello';
const typeOfNumber = typeof 42;
const typeOfBoolean = typeof true;
const typeOfUndefined = typeof undefined;
const typeOfObject = typeof {};
const typeOfFunction = typeof function() {};
const typeOfSymbol = typeof Symbol();
const typeOfBigInt = typeof 1n;

// new 运算符
const newDate = new Date();
const newArray = new Array(5);
const newObject = new Object();
const newMap = new Map();
const newSet = new Set();

// super 关键字
class Parent {
  constructor() {
    this.parentProp = 'parent';
  }
  
  method() {
    return 'parent method';
  }
}

class Child extends Parent {
  constructor() {
    super();
    this.childProp = 'child';
  }
  
  method() {
    return super.method() + ' overridden';
  }
}

// this 关键字
const thisExample = {
  value: 42,
  getValue: function() {
    return this.value;
  },
  getArrowValue: () => thisExample.value
};

// ==================== 更多关键字测试 ====================
const asyncKeyword = async () => {};
const awaitKeyword = async () => { await Promise.resolve(); };
const classKeyword = class {};
const extendsKeyword = class extends Animal {};
const staticKeyword = class { static method() {} };
const getKeyword = class { get value() { return 1; } };
const setKeyword = class { set value(v) {} };
const constructorKeyword = class { constructor() {} };

// export 关键字（注释掉，因为需要在模块环境中）
// export const exportedVar = 42;
// export function exportedFn() {}
// export class ExportedClass {}
// export { exportedVar as alias };
// export default function() {}

// import 关键字（注释掉，因为需要在模块环境中）
// import defaultExport from './module.js';
// import { namedExport } from './module.js';
// import { namedExport as alias } from './module.js';
// import * as namespace from './module.js';
// import('./dynamic.js').then(module => {});

// ==================== 边界情况测试 ====================
const emptyString = '';
const singleCharString = 'a';
const longString = 'a'.repeat(1000);
const unicodeString = '你好世界 🌍 𝕳𝖊𝖑𝖑𝖔';
const emojiString = '😀🎉🚀💻🔥';

const zero = 0;
const negativeZero = -0;
const infinity = Infinity;
const negativeInfinity = -Infinity;
const notANumber = NaN;

const emptyObjectLiteral = {};
const objectWithNumericKeys = { 0: 'zero', 1: 'one' };
const objectWithSymbolKeys = { [Symbol('key')]: 'value' };

const sparseArray = new Array(10);
sparseArray[5] = 'value';

// ==================== 复杂嵌套测试 ====================
const complexNested = {
  level1: {
    level2: {
      level3: {
        level4: {
          level5: {
            value: 'deeply nested',
            array: [
              { nested: { value: 1 } },
              { nested: { value: 2 } },
              { nested: { value: 3 } }
            ],
            func: function() {
              return () => {
                return () => {
                  return 'deeply nested function';
                };
              };
            }
          }
        }
      }
    }
  }
};

// ==================== 链式调用测试 ====================
const chainResult = [1, 2, 3, 4, 5]
  .filter(x => x > 2)
  .map(x => x * 2)
  .reduce((sum, x) => sum + x, 0)
  .toString()
  .split('')
  .map(Number)
  .filter(x => !isNaN(x))
  .join('-');

// ==================== 混合语法测试 ====================
async function mixedSyntaxTest(arr) {
  try {
    const results = await Promise.all(
      arr
        .filter(item => item?.value !== undefined)
        .map(async item => {
          const { value, ...rest } = item;
          const processed = value ?? 0;
          return {
            ...rest,
            processed,
            timestamp: Date.now()
          };
        })
    );
    
    const grouped = results.reduce((acc, item) => {
      const key = item.category ?? 'default';
      return {
        ...acc,
        [key]: [...(acc[key] ?? []), item]
      };
    }, {});
    
    return Object.entries(grouped)
      .filter(([key]) => key !== 'excluded')
      .map(([key, values]) => ({
        category: key,
        count: values.length,
        total: values.reduce((sum, v) => sum + v.processed, 0)
      }));
  } catch (error) {
    console.error('Error:', error);
    return [];
  } finally {
    console.log('Cleanup');
  }
}

// ==================== 最终测试 ====================
console.log('JavaScript 语法测试用例加载完成');
console.log('总行数:', 1100);
