// 函数
// 基本的使用跟其它的语言是一样的
// function add(a: number, b: number): number {
//     return a + b;
// }
const add = (a: number = 10, b: number = 20): number => {
    return a + b;
}
// 默认值直接在类型后面=进行设置
// 默认值跟可选是不能同时存在的

// 对于参数是对象 那一样通过interface定义类型
interface Obj {
    arr: number[],
    add: (this: Obj, num: number) => number
}

let obj: Obj = {
    arr: [1, 2, 3],
    add(this: Obj, num: number) {
        // ts可以定义this的类型
        // 定义的时候第一个参数必须是this的类型
        this.arr.push(4);
        return this.arr[3]
    },
}
console.log(obj.add(3))

// 函数重载 跟一般的后端语言一样
// 就是根据传入的内容来进行重载