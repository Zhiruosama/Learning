const { exec, execSync, spawn, spawnSync, execFile, execFileSync, fork } = require('child_process')
const path = require('node:path')
//1.exec 异步的方法 回调函数 返回buffer 可以帮我们执行shell命令 或者跟软件交互
// exec('node -v',(err,stdout,stderr)=>{
//     if(err){
//         return err
//     }
//     console.log(stdout.toString());
// })
//2.execSync 同步的方法
// const nodeVersion = execSync('node -v')//buffer
// console.log(nodeVersion.toString());
//3.执行较小的shell命令 想要立马拿到结果的shell适合execSync exec字节上限200kb


//4.spawn 没有字节上限 返回的是个流
//5.spawnSync用的比较少
// const { stdout } = spawn('netstat')
// stdout.on('data', (msg) => {
//     console.log(msg.toString());
// })
// stdout.on('close', (msg) => {
//     console.log('结束');
// })

//6.execFile 执行可执行文件
//底层实现顺序
// exec -> execFile -> spawn

//7.fork 只能接受js模块 帮js创建子进程

const testProcess = fork('./test.js')
testProcess.send('我是主进程')

//Nodejs不适合CPU密集应用 
//所以可以把耗时代码塞到子进程

testProcess.on('message',(res)=>{
    console.log(res);
})
//底层是IPC通讯 IPC基于libuv