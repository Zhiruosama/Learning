//nodejs 定义全局变量
// global.name = 'index.js'
// require('./child.js')
//浏览器环境定义全局变量用window
//nodejs环境定义全局变量用global
//globalThis根据环境自动判断 实现跨平台

//js三部分组成 ECMAScript DOM BOM
//nodejs环境下用不了DOM BOM相关的API
//ECMAScript相关的大部分可以用

//node环境内置API
// dirname 当前文件所在的目录
// 执行脚本的目录
// console.log(__dirname);
// file 文件名
// extname 文件后缀名
// filename 当前文件的绝对路径
// process 处理进程
// setTimeout(() => {
//     console.log('结束');
// }, 5000)
// setTimeout(() => {
//     process.exit()
// }, 2000)