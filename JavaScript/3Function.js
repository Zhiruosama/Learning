// 除了正常的根据规则调用
// JS里的函数调用更加的随意 参数可以传多也可以传少
// 但是这样容易出现错误
// 所以要避免收到undefined，可以对参数进行检查：

function abs(x) {
    if (typeof x !== 'number') {
        throw 'Not a number';
    }
    if (x >= 0) {
        return x;
    } else {
        return -x;
    }
}

// arguments关键词
// 只在函数内部生效 类似于一个数组
// 存储着传入的所有参数
// 一般就是用来判断传入了多少个参数 那知道了传入的参数个数
// 其实就可以在函数的内部做不同的条件判断了

// foo(a[, b], c)
// 接收2~3个参数，b是可选参数，如果只传2个参数，b默认为null：
function foo(a, b, c) {
    if (arguments.length === 2) {
        // 实际拿到的参数是a和b，c为undefined
        c = b; // 把b赋给c
        b = null; // b变为默认值
    }
    // ...
}

// rest关键词
// ES6引入的东西 当传入了已定义参数外的参数 rest会类似于一个数组一样将它们全部接下来
function foo(a, b, ...rest) {
    console.log('a = ' + a);
    console.log('b = ' + b);
    console.log(rest);
}

foo(1, 2, 3, 4, 5);
// 结果:
// a = 1
// b = 2
// Array [ 3, 4, 5 ]

foo(1);
// 结果:
// a = 1
// b = undefined
// Array []