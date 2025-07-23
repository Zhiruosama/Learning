let person = { name: "测试", age: 21 }
//proxy支持对象 数组 函数 set map
//ES6新增 作拦截器
// person.name//取值
// person.name = 'xxx'//赋值
// let personProxy = new Proxy(person, {
//     get() {
//         //拦截取值
//     },
//     set(target, key, value, receiver) {
//         //拦截赋值
//         return true
//     },
//     apply() {
//         //拦截函数调用
//     },
//     has() {
//         //拦截in操作符
//     },
//     ownKeys() {
//         //拦截for in
//     },
//     construct() {
//         //拦截new操作符
//     },
//     deleteProperty(target, p) {
//         //拦截删除操作
//     }
// })

// let personProxy = new Proxy(person, {
//     get(target, key, receiver) {
//         if (target.age <= 18) {
//             return Reflect.get(target, key, receiver)
//         } else {
//             return '成年了'
//         }
//     }
// })

// console.log(personProxy.age);

console.log(Reflect.get(person, 'name', person));
