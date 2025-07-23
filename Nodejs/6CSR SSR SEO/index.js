//CSR (Client-Side Rendering) 客户端渲染
//SSR (Server-Side Rendering) 服务端渲染
//SEO (Search Engine Optimization) 搜索引擎优化

const { JSDOM } = require('jsdom')
const fs = require('node:fs')
const root = new JSDOM(`<!DOCTYPE html><html lang="en"><head></head><body>
    <div id='app'></div>
    </body></html>`)

//请求一个接口 拿到数据然后填充到app
//fetch ndoe18版本之后才有 用法同浏览器
const window = root.window
const document = root.window.document
fetch('https://api.thecatapi.com/v1/images/search?limit=10&page=1').then(res => res.json()).then(data => {
    const app = document.getElementById('app');

    data.forEach(item => {
        const img = document.createElement('img')
        img.src = item.url
        img.style.width = '200px'
        img.style.height = '200px'
        app.appendChild(img)
    })
    //第一个参数是文件路径 第二个参数是内容
    //向文件内写入内容
    fs.writeFileSync('./index.html', root.serialize())
})
//SSR 服务端渲染
//CSR 客户端渲染 vue react SPA(单页面应用)

