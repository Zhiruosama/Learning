// OS模块用于与操作系统进行交互
const os = require('node:os')
//1.platform 获取操作系统平台
// console.log(os.platform());
//2.release获取操作系统版本
// console.log(os.release());
//3.vresion获取操作系统的版本（家庭版 专业版）
// console.log(os.version());

//4.homedir 读取用户的目录 底层原理就是%userprofile%
// console.log(os.homedir());
//5.cpu架构(安卓里用的多)
// console.log(os.arch());
//操作系统线程cpu的信息
// console.log(os.cpus());

