// Symbol 是 TypeScript（也是 JavaScript）中一种原始数据类型，
// 在 ES6 (ECMAScript 2015) 中引入。
// 它主要用于创建独一无二的属性键 (unique property keys)
// 这使得它在防止对象属性名冲突方面非常有用

let a1: symbol = Symbol(1) // 唯一的
let a2: symbol = Symbol(1)
console.log(a1 === a2);// fasle

// for方法会在全局symbol中查找是否注册过这个key 有的话会直接使用 没有就创建s
console.log(Symbol.for('123') === Symbol.for('123'));

let obj = {
    name: 1,
    age: 2,

    // 两者key相同但是不会报错
    [a1]: 111,
    [a2]: 222,
}
console.log(obj);

// for in 不能读到symbol
// Object.keys(obj) 也不能读到symbol
for (let key in obj) {
    console.log(key)
}
// 想要取到Symbol 要用Object.getOwnPropertySymbols
// 还有Relect.ownKeys(obj)

//1.生成器

function* gen() {
    // 生成器是一种特殊类型的函数，它能够暂停执行并在以后恢复执行
    // 每次恢复时产生（"yield"）一个值
    yield Promise.resolve('测试')// 同步异步都支持
    yield '1'
    yield '2'
    yield '3'
}

const man = gen()
console.log(man.next());
console.log(man.next());
console.log(man.next());
console.log(man.next());
console.log(man.next());

//2.迭代器
//3.set map
let set: Set<number> = new Set([1, 1, 2, 2, 3, 3])//set天然去重
console.log(set);
let map: Map<any, any> = new Map()

function args() {
    console.log(arguments);// 伪数组
    // 形式上是一样的 但是没有数组的方法
}

// 但是这些类型都有迭代器
// 所以我们可以调用其自身的迭代器去进行遍历

const each = (value: any) => {
    let It: any = value[Symbol.iterator]()
    let next: any = { done: false }
    while (!next.done) {
        next = It.next()
        if (!next.done) {
            console.log(next.value);
        }
    }
}

// 自己写相对麻烦
// 所以我们可以用迭代器的语法糖
// 也就是for of
for(let value of set){
    console.log(set);
}
// 对象不可以用 因为对象身上没有迭代器

// ...是解构
// 而解构的底层原理也都是调用iteratorgenerics