// interface 接口
// 用来做类型约束 来定义一个类型
interface test {
    name: string
    age: number
}
//同名会重合
interface test {
    log: string
}

interface test {
    [propName: string]: any
    // 索引签名
    // 即剩余的属性不再做校验
}

let a: test = {
    name: "str",
    age: 123,
    // gender:"male" 报错 因为定义的test中没有这个类型
    log: "test"
}

interface test1 {
    name: string
    age?: number
    // 问号表示可选
    readonly cb: () => boolean
}

let b: test1 = {
    name: "str",
    age: 18,
    cb: () => {
        return false
    }
}

// 通过extends进行继承 跟class类似
interface a extends test1 {
    xxx: string
}

let c: a = {
    name: "str",
    age: 18,
    cb: () => {
        return false
    },
    xxx: "xxx"
}

// 函数类型定义
interface Fn {
    (name: string): number[]
}

const fn: Fn = function (name: string) {
    return [1, 2, 3]
}