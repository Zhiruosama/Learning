// 闭包（Closure），简单来说，就是函数和其创建时所处的词法环境（Lexical Environment）的组合。
// 用更形象的话说，一个闭包就是“一个函数记得并能够访问它被创建时所能访问的外部变量，即使那个外部作用域已经关闭了。”

// 简单来说，闭包就是让函数“记住”并能访问它被创建时所处的环境中的变量。即使创建它的那个环境（比如外部函数）已经运行完了，这个函数仍然能访问到那些变量。

// 闭包的作用
// 闭包主要有两大作用：

// 数据私有化 / 封装状态：
// 你可以创建只属于某个特定函数实例的“私有”数据。这些数据外部无法直接访问，只能通过闭包内部暴露的方法来操作。这就像给数据加了一层保护罩。

// 简单例子：计数器

// function createCounter() {
//     let count = 0; // 这个 count 是私有的，外面摸不着

//     return function() { // 返回的这个匿名函数就是闭包
//         count++; // 每次调用，它都操作同一个私有的 count
//         console.log(count);
//     };
// }

// const myCounter = createCounter(); // 创建一个计数器实例
// myCounter(); // 输出 1
// myCounter(); // 输出 2
// // 你不能直接访问 myCounter.count
// 这里的 count 只被 myCounter 这个闭包“记住”，其他任何地方都不能直接改它，只能通过调用 myCounter() 来增加。

// 让函数拥有“记忆”：
// 闭包允许函数记住它被创建时的特定参数或上下文，从而在后续调用时继续使用这些信息。这在需要配置或预设行为的场景中很有用。

// 简单例子：生成问候语

function greet(greeting) {
    // 这个 greeting 参数被返回的函数“记住”了
    return function(name) { // 返回的这个函数是闭包
        console.log(`${greeting}, ${name}!`);
    };
}

const sayHello = greet('Hello'); // sayHello 记住了 'Hello' 这个问候语
const sayHi = greet('Hi');       // sayHi 记住了 'Hi' 这个问候语

// sayHello('Alice'); // 输出 "Hello, Alice!"
// sayHi('Bob');      // 输出 "Hi, Bob!"
// sayHello 和 sayHi 是两个不同的闭包，它们各自“记住”了创建时传入的 greeting 参数，所以能生成不同的问候语。

// 总的来说，闭包就是让你的函数能保存住它自己的“小秘密”（变量），并且只通过它自己来操作这些秘密。它让 JavaScript 的函数变得更强大，能够创建更模块化、更灵活、更有状态的代码。