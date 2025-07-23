// 记得在package.json中的type修改模式

// Common js
// 五种模式
//1.引入自己编写的模块
// require('./test.js')
//2.引入第三方模块
// const md5=require('md5')
// console.log(md5('12345'));
//3.nodejs内置模块 fs http net os child_process
// const fs=require('node:fs')//高版本nodejs node:fs 低版本fs
// console.log(fs);
//4.C++扩展 addon napi node-gyp .node
//5.引入json文件
// const data = require('./data.json')
// console.log(data, 1);
// const data = require('./test.js')
// console.log(data);


// module
// import obj from './test.js'
// console.log(obj);
// esm不支持json文件
// nodejs 18可以强行引入 但会有警告
// import json from './data.json' assert {type: "json"}
// console.log(json);

// 对比
// cjs中顶层的this指向模块本身 而ES6中顶层this指向undefined
// cjs是基于运行时加载 esm是基于编译时的异步加载
// cjs可以修改值 esm不能修改值
// cjs不可以支持tree shaking esm支持
// import在函数模式下可以掺杂在逻辑里面