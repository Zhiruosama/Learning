//infer后跟一个变量名
// interface User {
//     name: string
//     age: number
// }

// type PromiseType = Promise<User>

// type GetPromiseType<T> = T extends Promise<infer U> ? U : T

// type T = GetPromiseType<PromiseType>

//infer协变
//产生协变会返回联合类型
// let obj1 = {
//     name: 'ceshi',
//     age: 18
// }
// type Bar<T> = T extends { name: infer U, age: infer U } ? U : T
// type T = Bar<typeof obj1>

//infer逆变 出现在函数的参数上
//逆变返回的是交叉类型
type Bar<T> = T extends {
    a: (x: infer U) => void,
    b: (x: infer U) => void
} ? U : never

type T = Bar<{ a: (x: number) => void, b: (x: string) => void }>