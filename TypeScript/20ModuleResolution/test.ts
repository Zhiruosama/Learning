// 使用方式import 如果是默认导出名字随便起
import * as api from './index'
console.log(api);
// * as api 输出api查看导出的所有内容

// 动态引入
// import只能在最上层使用
if (true) {
    //返回的是一个Promise
    import('./index').then(res => {
        console.log(res);
    })
}