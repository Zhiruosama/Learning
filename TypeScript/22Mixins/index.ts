// 混入
// Mixins (混入) 是一种在面向对象编程中复用代码的技术。
// 它的核心思想是将一个或多个对象（或类）的属性和方法复制到另一个对象（或类）中
// 从而增强目标对象的功能，而无需使用传统的继承机制。

interface A {
    age: number
}
interface B {
    name: string
}
let a: A = {
    age: 18
}
let b: B = {
    name: '测试'
}

// 如何合并
//1. 扩展运算符 浅拷贝 返回新的类型
let c = { ...a, ...b }
console.log(c);
//2. Object.assign 浅拷贝 交叉类型
let c2 = Object.assign({}, a, b)
console.log(c2);

// 2.类的混入
// 插件类型的混入
class Logger {
    log(msg: string) {
        console.log(msg);
    }
}
class Html {
    render() {
        console.log('render');
    }
}
class App {
    run() {
        console.log('run');
    }
}

type Custructor<T> = new (...args: any[]) => T
function pluginMinxins<T extends Custructor<App>>(Base: T) {
    return class extends Base {
        private Logger = new Logger()
        private Html = new Html()
        constructor(...args: any[]) {
            super(...args)
            this.Logger = new Logger()
            this.Html = new Html()
        }
        run() {
            this.Logger.log('run')
        }
    }
}

const mixins = pluginMinxins(App)
const app = new mixins()
app.run()