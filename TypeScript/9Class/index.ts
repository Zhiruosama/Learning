// 类
// 类的定义
class Person {
    constructor() {

    }
    run() {

    }
}

interface Options {
    el: string | HTMLElement
}

interface VueCls {
    options: Options
    init(): void
}

// implements用于约束
class Vue implements VueCls {
    options: Options
    constructor(options: Options) {
        this.options = options
    }
    init(): void { }
}

// extends用于继承
class Person1 {
    public name: string
    private age: number
    protected some: any
    constructor(name: string, ages: number, some: any) {
        this.name = name
        this.age = ages
        this.some = some
    }
    run() {

    }
}

class Man extends Person1 {
    constructor() {
        super("张三", 18, 1)
        console.log(this.some)
    }
    create() {
        console.log(this.some)
    }
}

// 类的修饰符
// public private protected
// private只能在父类内部使用(私有)
// 实例化以后也是无法使用
// protected只能给子类和内部使用
// 所有方法默认为public
// 即对外开放

// static 静态属性 和 静态方法
// 静态方法只能调用静态方法和静态属性

class Ref {
    _value: any
    constructor(value: any) {
        this._value = value;
    }
    // get set拦截器
    get value() {
        return this._value + 'vvv'
    }
    set value(newVal) {
        this._value = newVal + 'test'
    }
}

const ref = new Ref('哈哈哈')
ref.value = '坏人'
console.log(ref.value);
