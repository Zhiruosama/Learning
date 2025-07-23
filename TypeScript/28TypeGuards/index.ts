// 类型守卫 (Type Guards) 是一种特殊的机制
// 它允许你在运行时（runtime）检查变量的类型
// 并帮助 TypeScript 编译器在代码的特定作用域内收窄 (narrow) 变量的类型

//1.类型收缩|收窄
//typeof是有缺陷的 数组 对象 函数 null都返回object
const isString = (str: any) => typeof str === 'string'

const isArr = (arr: any) => arr instanceof Array //这样就可以区分数组了

//2.类型谓词|自定义守卫

const isObj = (arg: any) => ({}).toString.call(arg) === '[Object Object]'
const isNum = (num: any) => typeof num === 'number'
const isStr = (str: any) => typeof str === 'string'
const isFn = (fn: any) => typeof fn === 'function'
const fn = (data: any) => {
    if (isObj(data)) {
        let val;
        //遍历属性不能用for in 因为for in会遍历原型上的属性
        Object.keys(data).forEach(key => {
            val = data[key]
            if (isNum(val)) {
                data[key] = val.toFixed(2)
            }
            if (isStr(val)) {
                data[key] = val.trim()
            }
            if (isFn(val)) {
                data[key]()
            }
        })
    }
}

let obj = {
    a: 100.2222,
    b: ' test ',
    c: function () {
        console.log(this);
        return this.a
    }
}

fn(obj)