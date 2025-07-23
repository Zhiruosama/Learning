// 枚举类型
// JS中没有 TS中的新东西
// 跟别的语言中的枚举类型几乎没有区别

// 默认从0开始
enum Color {
    red,
    green,
    blue
}

console.log(Color.red);
console.log(Color.green);
console.log(Color.blue);

// 自己设定值的时候
// 在被设定的值后会自动增长
enum Color1 {
    red = 1,
    green = 5,
    blue
}

enum Color2 {
    red = 'red',
    green = 'green',
    // blue 当上边是字符串时 下边无法自增
}

interface A {
    red: Color.red
}

let obj: A = {
    red: 0
    // red: 1 报错
}

const enum Types {
    sucess,
    fail
}

let code: number = 0
if (code === Types.sucess) {
    console.log('sucess');
}