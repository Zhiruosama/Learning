// Object是原型链的顶部 所以Object可以赋值任意类型

let a: Object = 123
let a1: Object = 'abc'
let a3: Object = []
let a4: Object = {}
let a5: Object = () => 123

// object 表示 非原始类型
// 只能赋值引用类型
let a6: object = []
let a7: object = {}
let a8: object = () => 123
// let a9: object = 123 报错

// {}类型 可以对任意类型赋值 但不能修改值 没啥用
let a9: {} = { name: 1 }
// a9.age = 'abc'