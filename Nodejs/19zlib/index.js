//zlib提供数据压缩和解压缩的功能
const zlib = require('node:zlib')
const fs = require('node:fs')

//压缩 Gzip
// const readStream = fs.createReadStream('index.txt')
// const writeStream = fs.createWriteStream('index.txt.gz')//压缩后后缀为.gz
// readStream.pipe(zlib.createGzip()).pipe(writeStream)

//解压缩
// const readStream = fs.createReadStream('index.txt.gz')
// const writeStream = fs.createWriteStream('index2.txt')
// readStream.pipe(zlib.createGunzip()).pipe(writeStream)

//deflate压缩
// const readStream = fs.createReadStream('index.txt')
// const writeStream = fs.createWriteStream('index.txt.deflate')//压缩后后缀为.deflate
// readStream.pipe(zlib.createDeflate()).pipe(writeStream)

//解压
const readStream = fs.createReadStream('index.txt.deflate')
const writeStream = fs.createWriteStream('index3.txt')
readStream.pipe(zlib.createInflate()).pipe(writeStream)
