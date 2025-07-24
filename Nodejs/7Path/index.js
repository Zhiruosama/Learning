const path = require('node:path')
//1. basename 返回给定路径的最后一部分
//windows系统默认是是反斜杠\
//windows同时也兼容正斜杠/
//posix处理不了windows的路径
//如果要在macos上处理windows路径可以用path.win32.xxx
// console.log(path.basename('/foo/bar/baz/sda/cs.html'));

//2. dirname 返回路径的目录名
//与上面正好相反
// console.log(path.dirname('/foo/bar/baz/sda/cs.html'));

//3.extname返回路径的扩展名 .html
//返回值带.
//如果没有. 返回空字符串
//如果有多个点 返回最后一个点后面的内容
// console.log(path.extname('index.js'));

//4.path.join()拼接路径
// console.log(path.join('/a', '/b', '/c'));

//5.path.resolve() 解析路径 返回的是绝对路径
//都是绝对路径 返回最后一个
//如果只有一个相对路径，返回当前工作目录的绝对路径
// console.log(path.resolve('/a', '/b', '/c'));
// console.log(path.resolve('./index.js'));

//6.path.parse() 解析路径 返回一个对象
// console.log(path.parse('/home/user/dir/file.txt'));
// {
//   root: '/', 根目录
//   dir: '/home/user/dir', 文件所在目录
//   base: 'file.txt', 文件名+后缀名
//   ext: '.txt', 后缀名
//   name: 'file' 文件名
// }

//7.path.format()就是parse的逆过程

//8.path.sep
//windwos返回的是\
//posix返回的是/
