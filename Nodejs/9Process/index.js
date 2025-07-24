//process不需要引入
//1.CPU架构
// console.log(process.arch);
//2.操作系统获取
// console.log(process.platform);
//3.argv 返回数组
// console.log(process.argv, process.argv.includes('--version') ? '1.0.0' : '无');
// [
//   'C:\\Users\\15126\\AppData\\Local\\fnm_multishells\\7088_1753326423634\\node.exe',
//   'D:\\Github\\Learning\\Nodejs\\9Process\\index.js'
// ]

//4.cwd() 获取工作目录 esm模式下用不了__dirname 可以用cwd代替
// console.log(process.cwd());

//5.内存信息
// console.log(process.memoryUsage());
// {
//   rss: 29237248, 常驻集大小 物理内存的存量
//   heapTotal: 5537792, V8分配到堆内存的总大小包括未使用的内存
//   heapUsed: 3821048, 已使用的内存
//   external: 1208756, 外部的内存 C C++使用
//   arrayBuffers: 10515 二进制的总量
// }

//5.exit 退出进程
//6.kill 杀死进程 需要Pid 即进程id
// setTimeout(() => {
//     console.log(5);

// }, 5000)
// setTimeout(() => {
//     process.exit()
// }, 2000)

//7.env环境变量 获取操作系统所有的环境变量
//修改只会在当前进程生效 不会影响系统的环境变量
//可以区分开发环境和生产环境