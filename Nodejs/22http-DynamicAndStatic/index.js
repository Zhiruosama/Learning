//动态静态资源分离
import http from 'node:http'
import fs from 'node:fs'
import { log } from 'node:console'
import path from 'node:path'
import mime from 'mime'

const server = http.createServer((req, res) => {
    const { method, url } = req
    //静态资源
    if (method === 'GET' && url.startsWith('/static')) {
        const staticPath = path.join(process.cwd(), url)
        fs.readFile(staticPath, (err, data) => {
            if (err) {
                res.writeHead(404, {
                    'Content-Type': 'text/plain'
                })
                res.end('not found')
            } else {
                console.log('测试');
                const type = mime.getType(staticPath)
                res.writeHead(200, {
                    'content-type': type,
                    "cache-control": 'public,max-age=3600'
                })
                res.end(data)
            }
        })
    }
})

server.listen(80, () => {
    console.log('server is running at http://localhost');
})
