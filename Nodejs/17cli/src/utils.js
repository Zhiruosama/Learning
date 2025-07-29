import fs from 'node:fs'
import download from 'download-git-repo'
import ora from 'ora'
const spinner = ora('下载中')
//检查路径
export const checkPath = (path) => {
    if (fs.existsSync(path)) {
        return true
    } else {
        return false
    }
}

export const downloadTemp = (branch,project) => {
    return new Promise((resolve, reject) => {
        spinner.start()
        download(`direct:https://gitee.com/chinafaker/vue-template.git#${branch}`, project, { clone: true }, function (err) {
            if (err) reject(err)
            resolve()
            spinner.succeed('下载完成')
        })
    })
}