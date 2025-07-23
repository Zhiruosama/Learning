// 模块解析
// es6模块化规范出来之后上面这些模块化规范就用的比较少了
// 现在主要使用 import export 
// 1.默认导出 导出的可以是任何类型
export default 1
// 一个模块只能出现一个默认导出
// 2.分别导出
// export let x = 2
// export const add = (a: any, b: any) => {
//     return a + b
// }
// export let arr = [1, 2, 3, 4]
let x = 2
const add = (a: any, b: any) => {
    return a + b
}
let arr = [1, 2, 3, 4]
//3.解构导出
export{
    x,
    add,
    arr
}