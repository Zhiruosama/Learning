// 学习别的语言的时候都知道 函数内部定义的参数变量在函数外部其实是访问不到的
// JS这边也是这样

//那考虑一个特殊一点的 就是内外的变量名重复 情况会如何呢
function foo() {
    var x = 1;
    function bar() {
        var x = 'A';
        console.log('x in bar() = ' + x); // 'A'
    }
    console.log('x in foo() = ' + x); // 1
    bar();
}

foo();
// 可以发现JS的函数在查找变量时从自身函数定义开始，从“内”向“外”查找。如果内部函数定义了与外部函数重名的变量
// 则内部函数的变量将“屏蔽”外部函数的变量。

// 变量提升
// JavaScript的函数定义有个特点，它会先扫描整个函数体的语句，把所有用var申明的变量“提升”到函数顶部

function foo() {
    var x = 'Hello, ' + y;
    console.log(x);
    var y = 'Bob';
}

foo();

// 这里的var声明会被提升
function foo() {
    var y; // 提升变量y的申明，此时y为undefined
    var x = 'Hello, ' + y;
    console.log(x);
    y = 'Bob';
}
// 最终JS引擎看到的会是这样

// 全局作用域
// 不在任何函数内定义的变量就具有全局作用域
var course = 'Learn JavaScript';
console.log(course); // 'Learn JavaScript'
console.log(window.course); // 'Learn JavaScript'
// 其实几乎所有的函数可以通过window.调用出来
window.alert('调用window.alert()');
// 把alert保存到另一个变量:
let old_alert = window.alert;
// 给alert赋一个新函数:
window.alert = function () {}

alert('无法用alert()显示了!');

// 恢复alert:
window.alert = old_alert;
alert('又可以用alert()了!');
// 此处就是将全局作用域内的alert给修改了 导致其无法使用了
// JS只有一个全局作用域，任何变量如果在当前的函数作用域找不到的话就会往上查找，直到全局作用域
// 如果全局作用域都找不到的话就报错

// 名字空间
// 从前面的全局变量讲解知道
// 所有的全局变量会被绑到window上去
// 所以如果定义重名 很容易出问题
// 所以减少冲突的一个方法是把自己的所有变量和函数全部绑定到一个全局变量中。例如：
// 唯一的全局变量MYAPP:
let MYAPP = {};

// 其他变量:
MYAPP.name = 'myapp';
MYAPP.version = 1.0;

// 其他函数:
MYAPP.foo = function () {
    return 'foo';
};

