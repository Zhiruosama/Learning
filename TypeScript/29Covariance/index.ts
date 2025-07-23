//主类型
interface A {
    name: string
    age: number
}
//子类型
interface B {
    name: string
    age: number
    sex: string
}

let a: A = {
    name: "123",
    age: 33
}

let b: B = {
    name: "234",
    age: 22,
    sex: "男"
}

//协变
a = b
//b一定包含a中所有类型

//逆变
let fna = (params: A) => {

}
let fnb = (params: B) => {

}
// fna = fnb
//函数逆变 顺序不对
fnb = fna