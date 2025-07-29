#!/usr/bin/env node
//告诉操作系统 执行自定义命令的时候 使用node去执行这个文件
import { program } from 'commander'
import inquirer from 'inquirer'
import fs from 'node:fs'
import { checkPath, downloadTemp } from './utils.js'
let JsonBuffer = fs.readFileSync('./package.json')
let JsonString = JsonBuffer.toString('utf8')
let json = JSON.parse(JsonString)
program.version(json.version)

program.command('create <projectName>').alias('c').description('创建项目').action((projectName) => {
    inquirer.prompt([{
        type: 'input',//输入input 确认框confirm 列表list
        name: 'projectName',//返回值的key
        message: '请输入项目名',//描述
        default: projectName//默认值
    },
    {
        type: 'confirm',
        name: 'isTs',
        message: '是否选用typescript',
    }
    ]).then(res => {
        if (checkPath(res.projectName)) {
            console.log('文件已存在');
            return
        }
        if (res.isTs) {
            downloadTemp('ts', res.projectName)
        } else {
            downloadTemp('js', res.projectName)
        }
    })
})






program.parse(process.argv)