// util 模块是一个实用工具模块，它提供了一系列常用的函数
// 用于处理各种编程任务，例如类型检查、格式化输出、继承实现等
import util from 'node:util'
import { exec } from 'node:child_process'

// exec('node -v', (err, stdout, stderr) => {
//     if (err) {
//         return err
//     }
//     console.log(stdout);
// })

// const execPromise = util.promisify(exec)//返回一个新的函数
// //如果返回多个参数 res会是一个对象 如果只返回了一个参数 就直接返回
// execPromise('node -v').then(res => {
//     console.log(res, 'res');
// }).catch(err => {
//     console.log(err, 'err');
// })

//callbackify 将一个函数转换为回调函数风格
// const fn = (typee) => {
//     if (type === 1) {
//         return Promise.resolve('sucess')
//     } else {
//         return Promise.reject('error')
//     }
// }

// const callback = util.callbackify(fn)

// callback(1, (err, value) => {
//     console.log(err, value);
// })

//format 类似C语言中的printf
console.log(util.format('%s---%s', 'xm', 'zs'));