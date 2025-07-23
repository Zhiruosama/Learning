// ECMAScript 的内置对象
// Boolean、Number、string、RegExp、Date、Error

let b: Boolean = new Boolean(1)
let n: Number = new Number(true)
let s: String = new String('fortest')
let d: Date = new Date()
let r: RegExp = /^1/
let e: Error = new Error("error!")

// DOM 和 BOM 的内置对象
// Document、HTMLElement、Event、NodeList 等

// Promise对象定义 可以指定返回类型
// 指定返回类型后会有代码提示
function promise(): Promise<number> {
    return new Promise<number>((resolve, reject) => {
        resolve(1)
    })
}

promise().then(res => {
    console.log(res)
})