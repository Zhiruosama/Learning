import express from 'express';
//express是个函数
import User from './src/user.js'
import List from './src/list.js'
import LoggerMiddleware from './middleware/logger.js';
const app = express();
app.use(express.json()) //支持post解析json
app.use(LoggerMiddleware)
app.use('/user',User)
app.use('/list',List)
//get
//第一个参数为api地址
//第二个参数是回调函数 req res
//req.query
app.get('/get', (req, res) => {
    res.send('get');
})

//post req.body
app.post('/post', (req, res) => {
    res.send('post')
})

//动态参数 req.params
app.get('/get/:id', (req, res) => {
    res.send('动态参数');
})



app.listen(3000, () => {
    console.log('http://localhost:3000');
})