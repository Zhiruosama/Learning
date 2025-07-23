// 1. map 数组转换器
// 认识 map()：数组转换器
// map() 是 JavaScript 数组的一个高阶函数，这意味着它接受一个函数作为参数。
// 它的核心作用是创建一个新数组，这个新数组的元素是原数组中每个元素调用你提供的函数后返回的结果。

// eg:
// const numbers = [1, 2, 3, 4];

// 使用 map() 将每个数字翻倍
const doubledNumbers = numbers.map(function(number) {
    return number * 2;
});

console.log(doubledNumbers); // 输出: [2, 4, 6, 8]
console.log(numbers);       // 输出: [1, 2, 3, 4] (原数组未改变)

// 甚至可以arr.map(String)来将数组中的所有数字转换为字符

// 2.reduce 数组的“折叠”与“累积”器
// reduce() 方法是 JavaScript 数组中用于对数组中的所有元素执行一个累加器（reducer）函数
// reduce要传一个回调函数
// 回调可以传四个参数 核心的是累积器：及上次回调完的结果 还有当前正在处理的数组元素
// 剩下的index array可选

const numbers = [1, 2, 3, 4, 5];

// 求和，提供初始值 0
const sum = numbers.reduce((accumulator, currentValue) => {
    return accumulator + currentValue;
}, 0); // 初始值设置为 0

console.log(sum); // 输出: 15 (1 + 2 + 3 + 4 + 5)

// 3.filter 数组筛选器
// filter() 方法顾名思义，就像一个筛选器。它会遍历数组中的每个元素，并根据你提供的回调函数（一个测试函数）来判断该元素是否应该被“留下”
// 例如，在一个Array中，删掉偶数，只保留奇数，可以这么写：

// let arr = [1, 2, 4, 5, 6, 9, 10, 15];
let r = arr.filter(function (x) {
    return x % 2 !== 0;
});
r; // [1, 5, 9, 15]

// 4.sort 排序函数
// Array的sort()方法默认把所有元素先转换为String再排序
// 所以几乎不会使用sort的默认方法 结果会很抽象
// 所以我们需要自己传入一个回调函数来设定规则
// 通常规定，对于两个元素x和y，如果认为x < y，则返回-1，如果认为x == y，则返回0，如果认为x > y，则返回1
// 要按数字大小进行排序 可以这样写
let arr = [10, 20, 1, 2];

arr.sort(function (x, y) {
    if (x < y) {
        return -1;
    }
    if (x > y) {
        return 1;
    }
    return 0;
});

console.log(arr); // [1, 2, 10, 20]

// 需要注意的是 sort()方法会直接对Array进行修改

