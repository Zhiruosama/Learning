// TypeScript 会在没有明确的指定类型的时候推测出一个类型，这就是类型推论
let str = '123'
// 如果你声明变量没有定义类型也没有赋值这时候TS会推断成any类型可以进行任何操作
let any

// 可以通过type关键字来给类型起别名
type s = string | number
// 这样看有点像interface
// 当然两者的用处是大不相同的
// interface在重名时会自动合并 但type不会

// extends在type中是包含的意思
// 左边的值会作为右边类型的子类型
// 1.顶级类型any unknow
// 2.Object
// 3.Number...
// 4.number string
// 5.never
type num = 1 extends number ? 1 : 0