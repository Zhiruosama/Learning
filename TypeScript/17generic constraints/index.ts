// 泛型约束
// 语法 在泛型后 extends 类型
function add<T extends number>(a: T, b: T) {
    return a + b
}

interface Len {
    length: number
}

function fn<T extends Len>(a: T) {
    a.length
}
// 有length属性的类型就不会报错
fn([1, 2, 3])

let obj = {
    name: "测试",
    sex: "男"
}
// keyof
// 用于约束对象的key
type Key = keyof typeof obj
function ob<T extends object, K extends keyof T>(obj: T, key: K) {
    return obj[key]
}

ob(obj, 'name')

interface Data {
    name: string,
    age: number,
    sex: string
}
// for in for(ley key in obj)
type Options<T extends object> = {
    [Key in keyof T]?: T[Key]
}

type B =Options<Data>