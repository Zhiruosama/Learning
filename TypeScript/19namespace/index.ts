// 命名空间
// namespace里所有的变量以及方法必须要导出才能访问
namespace Test {
    // 变量 方法
    export let a = 1;
    export const add = (a: number, b: number) => a + b
    //嵌套
    export namespace Test2 { }
}
// 跟interface一样 同名命名空间会自动合并

console.log(Test.add(1, 2));

// namespace还支持抽离 及以文件形式引用
// import{}from'./src'

// 一般用于跨端逻辑拆分
namespace ios {
    export const PushNotification = (msg: string, type: number) => { }
}
namespace android {
    export const PushNotification = (msg: string) => { }
    export const CallPhone = (phone: number) => { }
}