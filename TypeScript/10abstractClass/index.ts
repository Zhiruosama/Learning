// 基类 抽象类
// abstract 所定义的抽象类
// abstract 所定义的方法 都只能描述不能进行实现
// 抽象类无法被实例化
abstract class Vue {
    name: string
    constructor(name?: string) {
        this.name = name
    }
    getName(): string {
        return this.name
    }
    abstract init(name: string): void
}

// 定义的抽象方法必须在派生类实现
class React extends Vue {
    constructor() {
        super()
    }
    init(name: string) {

    }
    setName(name: string) {
        this.name = name;
    }
}

const react = new React()
react.setName('测试-test')

console.log(react.getName());
