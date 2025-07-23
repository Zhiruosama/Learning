// 类型守卫 (Type Guards) 是一种特殊的机制
// 它允许你在运行时（runtime）检查变量的类型
// 并帮助 TypeScript 编译器在代码的特定作用域内收窄 (narrow) 变量的类型
//1.类型收缩|收窄
//typeof是有缺陷的 数组 对象 函数 null都返回object
var isString = function (str) { return typeof str === 'string'; };
var isArr = function (arr) { return arr instanceof Array; }; //这样就可以区分数组了
//2.类型谓词|自定义守卫
var isObj = function (arg) { return ({}).toString.call(arg) === '[Object Object]'; };
var isNum = function (num) { return typeof num === 'number'; };
var isStr = function (str) { return typeof str === 'string'; };
var isFn = function (fn) { return typeof fn === 'function'; };
var fn = function (data) {
    if (isObj(data)) {
        var val_1;
        //遍历属性不能用for in 因为for in会遍历原型上的属性
        Object.keys(data).forEach(function (key) {
            val_1 = data[key];
            if (isNum(val_1)) {
                data[key] = val_1.toFixed(2);
            }
            if (isStr(val_1)) {
                data[key] = val_1.trim();
            }
            if (isFn(val_1)) {
                val_1();
            }
        });
    }
};
var obj = {
    a: 100.2222,
    b: ' test ',
    c: function () {
        console.log(this);
        return this.a;
    }
};
fn(obj);
