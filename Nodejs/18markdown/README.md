# Node.js 学习笔记

## 简介

Node.js 是一个基于 Chrome V8 引擎的 JavaScript 运行环境，使开发者能够使用 JavaScript 来编写服务器端的应用程序。

## 主要特点

- **非阻塞 I/O** - 提高应用程序的性能和可扩展性
- **事件驱动** - 基于事件循环的编程模型
- **轻量且高效** - 适合构建网络应用
- **npm** - 世界上最大的开源库生态系统

## 基础示例

```javascript
// 创建一个简单的 HTTP 服务器
const http = require('http');

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('你好，世界！\n');
});

server.listen(3000, '127.0.0.1', () => {
  console.log('服务器运行在 http://127.0.0.1:3000/');
});
```

## 常用模块

| 模块名 | 描述                |
| ------ | ------------------- |
| fs     | 文件系统操作        |
| path   | 处理文件路径        |
| http   | HTTP 服务器和客户端 |
| events | 事件处理            |
| crypto | 加密功能            |

## 异步编程

### 回调函数

```javascript
fs.readFile('文件.txt', 'utf8', (err, data) => {
  if (err) throw err;
  console.log(data);
});
```

### Promise

```javascript
const fsPromises = require('fs').promises;

fsPromises.readFile('文件.txt', 'utf8')
  .then(data => console.log(data))
  .catch(err => console.error(err));
```

### Async/Await

```javascript
async function readFile() {
  try {
    const data = await fsPromises.readFile('文件.txt', 'utf8');
    console.log(data);
  } catch (err) {
    console.error(err);
  }
}

readFile();
```

## 学习资源

1. [Node.js 官方文档](https://nodejs.org/docs/)
2. [Node
