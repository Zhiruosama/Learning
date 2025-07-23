// npm init -y 生成package.json
// npm i @types/node --save-dev （node环境支持的依赖必装）
let a: number = 123
console.log(a)
// ts-node index.ts直接运行

//any 任意类型 unknown 未知类型
//1.top type 顶级类型 any unknown
//2.Object
//3.Number String Boolean
//4.number string boolean
//5.   1  'ceshistr'   false
//6.never

// any类型可以随意被赋值跟赋值
// unknown只能赋值给自身或者别的unknown类型
// unknown没办法读任何属性 方法也无法调用

let b: any = 1;
let c = b;
b = 'string';

let d: unknown = 'string';
// let e: string = d;

// 所以说unknown相对于any其实更加安全
// 当一个类型未知的时候优先unknown