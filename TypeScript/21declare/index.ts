// 声明文件declare
// npm init -y生成package.json
import axios from 'axios'// 有声明文件
import express from 'express' // 无声明文件

const app = express()

const router = express.Router()

app.use('/api', router)
router.get('./api', (req, res) => {
    res.json({
        code: 200
    })
})

app.listen(9001, () => {
    console.log('9001');
})