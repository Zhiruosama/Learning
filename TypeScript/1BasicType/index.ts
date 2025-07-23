// TS的基本类型跟JS完全重叠 因为TS也就是JS的超集
// 基础类型：Boolean、Number、String、null、undefined以及 ES6 的  Symbol 和 ES10 的 BigInt

// 严格模式下会严格区分各个类型
// null跟undefined这类就无妨划等
// 严格模式会更加有利于开发 所以一般选择开启

// tsc --init初始化
// tsc --w index.ts监听并编译为js
// node index.js进行运行

let notANumber: number = NaN;//Nan
let num: number = 123;//普通数字
let infinityNumber: number = Infinity;//无穷大
let decimal: number = 6;//十进制
let hex: number = 0xf00d;//十六进制
let binary: number = 0b1010;//二进制
let octal: number = 0o744;//八进制

// 严格模式下null不能赋予void类型
// undefined可以