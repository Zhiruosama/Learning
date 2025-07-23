// 返回never的函数必须存在无法达到的终点
function err(): never {
    throw new Error('错啦')
}
function err1(): never {
    while (true) { }
}

// 联合类型中never会被忽略
type A = void | number | never

type B = '唱' | '跳' | 'rap'
function kun(value: B) {
    switch (value) {
        case "唱":
            break
        case "跳":
            break
        case "rap":
            break
        default:
            const error: never = value;// 进行兜底逻辑
            break
    }
}