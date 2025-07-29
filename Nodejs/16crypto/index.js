// crypto模块的目的是为了提供通用的加密和哈希算法

//对称加密
//双方协商定义一个密钥以及iv
const crypto = require('node:crypto');

//第一个参数 algorithm 接受一个算法 aes-256-cbc
//第二个参数 key 密钥 32位
//第三个参数 iv 初始化向量 支持16位 保证生成的密钥串每次都是不一样的
let key = crypto.randomBytes(32)
let iv = Buffer.from(crypto.randomBytes(16))
const cipher = crypto.createCipheriv("aes-256-cbc", key, iv)
cipher.update('测试', "utf-8", "hex")
const result = cipher.final("hex")//输出密文
console.log(result);
//解密 算法 key iv都要相同
const de = crypto.createDecipheriv("aes-256-cbc", key, iv)
de.update(result, "hex", "utf-8")
// console.log(de.final("utf-8"));


//非对称加密
//生成公钥和私钥
//私钥只是管理员拥有 不能对外公开
//公钥可以对外公开
const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
    modulusLength: 2048,//长度越长越安全
})
//公钥加密
const encrypted = crypto.publicEncrypt(publicKey, Buffer.from('测试'))
//私钥解密
const decrypted = crypto.privateDecrypt(privateKey, encrypted)


//哈希函数
//不能被解密
//具有唯一性
let hash = crypto.createHash('sha256')//或者md5
hash.update('测试')
console.log(hash.digest('hex'));
