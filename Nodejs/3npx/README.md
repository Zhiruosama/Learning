npx是个命令行工具

它的主要目的是**执行 Node.js 包，而无需先将其全局安装**

应用:
例如现在需要构建react项目

需要先安装create-react-app

但如果使用npx 那可以在没有安装该包的情况下通过

npx create-react-app my-app直接进行执行

npx会先查找当前目录下的node_modules 如果没有就检查全局

全局依旧没有则会通过npm下载安装 然后执行 后删除

且每次下载一定是最新版本 并且不会浪费任何空间
