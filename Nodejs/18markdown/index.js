//EJS JS模板引擎 用于在HTML中嵌入动态内容
//Marked md转html
//BrowserSync 实时预览和同步网站更改
const ejs = require('ejs')
const fs = require('fs')
const marked = require('marked')
let browser;
const browserSync = require('browser-sync')

const server = () => {
    browser = browserSync.create()
    browser.init({
        server: {
            baseDir: './',
            index: 'index.html'
        }
    })
}

const init = (callback) => {
    //读取md内容
    const md = fs.readFileSync('README.md', 'utf-8')
    ejs.renderFile('template.ejs', {
        content: marked.parse(md),
        title: "markdown to html",
    }, (err, data) => {
        if (err) throw err
        fs.writeFileSync('index.html', data)
        callback && callback()
    })
}
fs.watchFile('README.md', (curr, prev) => {
    if (curr.mtime !== prev.mtime) {
        init(() => {
            browser.reload()
        })
    }
})
init(() => {
    server()
})
