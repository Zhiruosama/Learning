//ForEach用于遍历数组的每一项，并对每一项执行给定的回调函数。它是最常用的遍历数组的方法之一

// 1. 基本语法
// JavaScript
array.forEach(function(currentValue, index, array){
  // 执行的代码
}, thisArg);
// array：要遍历的数组。
// function：对每个数组元素执行的回调函数。
// currentValue：当前正在处理的元素的值。
// index（可选）：当前元素的索引。
// array（可选）：原始数组对象。
// thisArg（可选）：执行回调函数时用作 this 的值。

// 2. 用法示例
// JavaScript
// const fruits = ["apple", "banana", "cherry"];

fruits.forEach(function(item, index, arr) {
  console.log(item, index, arr);
});
// 输出：

// Code
// apple 0 [ 'apple', 'banana', 'cherry' ]
// banana 1 [ 'apple', 'banana', 'cherry' ]
// cherry 2 [ 'apple', 'banana', 'cherry' ]

// Set与Array类似，但Set没有索引，因此回调函数的前两个参数都是元素本身：

let s = new Set(['A', 'B', 'C']);
s.forEach(function (element, sameElement, set) {
    console.log(element);
});
// Map的回调函数参数依次为value、key和map本身：

let m = new Map([[1, 'x'], [2, 'y'], [3, 'z']]);
m.forEach(function (value, key, map) {
    console.log(value);
});
// 如果对某些参数不感兴趣，由于JavaScript的函数调用不要求参数必须一致，因此可以忽略它们。例如，只需要获得Array的element：

let a = ['A', 'B', 'C'];
a.forEach(function (element) {
    console.log(element);
});