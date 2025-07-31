const { log } = require('node:console')
const http = require('node:http')
const path = require('node:path')
const url = require('node:url')

//req接受前端信息
//res给前端返回的信息
http.createServer((req, res) => {
    const { pathname, query } = url.parse(req.url, true)
    if (req.method === 'POST') {
        if (pathname === '/login') {
            let data = ''
            req.on('data', (chunk) => {
                data += chunk
            })
            req.on('end', () => {
                res.setHeader('Content-Type', 'application/json')
                res.statusCode = 200
                res.end(data)
            })
        } else {
            res.statusCode = 404
            res.end('404')
        }
        res.end('POST')
    } else if (req.method === 'GET') {
        if (pathname === '/get') {
            console.log(query);
            res.end('GET')
        } else {
            res.statusCode = 404
            res.end('404')
        }
    }
}).listen(98, () => {
    console.log('98端口启动成功');
})