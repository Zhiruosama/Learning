// 装饰器
// 装饰器 是一种特殊类型的声明，它能附加到类声明、方法、访问器、属性或参数上。
// 简单来说，装饰器就是一个函数，它能修改或增强它所装饰的目标（比如一个类或一个方法）的行为
import axios from 'axios'
//1.类装饰器
const Base: ClassDecorator = (target) => {
    // target是Http的构造函数
    // 这样就可以直接为类增加新内容而不破坏原有结构
    target.prototype.ceshi = '测试'
    target.prototype.fn = () => {
        console.log('新增');
    }
}
// 函数柯里化
const Get = (url: string) => {
    const fn: MethodDecorator = (target, key, descriptor:TypedPropertyDescriptor<any>) => {
        axios.get(url).then(res => {
            descriptor.value(res.data)
        })
    }
    return fn
}

@Base
class Http {
    @Get('接口地址')
    getList(data: any) {
        console.log(data);
    }
    // @Post()
    // create() {

    // }
}
const http = new Http() as any
http.fn()