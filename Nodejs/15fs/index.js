const { log } = require('node:console')
const fs = require('node:fs')
const fs2 = require('node:fs/promises')
//1.异步 同步 promise
//不加Sync就是异步
// fs.readFile('./index.txt', {
//     encoding: 'utf-8',
//     flag: 'r'
// }, (err, data) => {
//     if (err) throw err
//     console.log(data);
// })

//同步会阻塞下面的代码
// let result = fs.readFileSync('./index.txt')

// fs2.readFile('./index.txt').then(result => {
//     console.log(result.toString("utf-8"));
// }).catch

// const readStream = fs.createReadStream('./index.txt')

// readStream.on('data', (chunk) => {
//     console.log(chunk.toString());
// })

//创建文件夹
// fs.mkdirSync('./cs')
// fs.mkdirSync('./cs/ts/fs', {
//     recursive: true
// })
//删除文件夹
// fs.rmSync('./cs', {
//     recursive: true
// })

//重命名 原名->新名
// fs.renameSync('./index.txt','./index2.txt')

//文件监听
// fs.watch('./index.txt', (event, filename) => {
//     console.log(event);
// })

//写入文件
//内容替换
// fs.writeFileSync('./index.txt','毁灭世界')
// fs.writeFileSync('./index.txt', '毁灭世界2.0', {
//     flag: 'a' //append
// })
// fs.appendFileSync('./index.txt', '追加内容')

//大量数据分批插入
// let verse = [
//     '123123123123',
//     '1231223123123',
//     '1231231231231',
//     'sdoaiwjdioawjid'
// ]

// let writeStream = fs.createWriteStream('./index.txt')

// verse.forEach(item => {
//     writeStream.write(item + '\n')
// })
// writeStream.end()
// writeStream.on('finish', () => {
//     console.log('写入完成');
// })
