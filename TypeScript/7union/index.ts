// 联合类型
// 就是一个变量可以是多种类型
let a: number | string = 123;
a = "123";

// 当后端传来的有数字或布尔值
// 可以通过!!来将其都转换为布尔值
const fn = (something: number | boolean): boolean => {
    return !!something
}

// 交叉类型
// 类似于extends
interface People {
    name: string,
    age: number
}

interface Man {
    sex: number
}

const xiaoman = (man: People & Man): void => {
    console.log(man);
}
xiaoman({
    name: "老王",
    age: 18,
    sex: 1
})

// 类型断言
// 只能用来“欺骗”TS 用处不是很大
interface A {
       run: string
}
 
interface B {
       build: string
}
 
const fn1 = (type: A | B): string => {
       return (type as A).run
       // A.run会报错 因为B没有run这个方法
}
//可以使用类型断言来推断他传入的是A接口的值

