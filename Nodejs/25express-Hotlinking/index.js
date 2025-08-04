import express from 'express';
import { log } from 'node:console';
import { ref } from 'node:process';

const app = express()
const whiteList = ['http://localhost:3000']
const preventHotLinking = (req, res, next) => {
    //referer
    //如果是直接打卡的资源是获取不到referer的
    const referer = req.get('referer')
    if (referer) {
        const { hostname } = new URL(referer)
        if (!whiteList.includes(hostname)) {
            res.status(403).send('403 Forbidden')
        }
    }
    next()
}
app.use(express.static('static'))

app.listen(3000, () => console.log('server started'))