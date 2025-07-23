//1.Record
//约束对象的Key以及Value
//所以接受两个泛型

//2.ReturnType
//作用是获取函数类型的返回值
const fn = () => [1, 2, 3]
type arrNum = ReturnType<typeof fn>

//对象的key只能是string number symbol
type ObjKey = keyof any;
type CustomRecord<K extends ObjKey, T> = {
    [P in K]: T
}

type Key = 'c' | 'x' | 'k'
type Value = '唱歌' | '跳舞' | 'rap'

//Key不可少 Value可随意
//且支持嵌套
let obj: Record<Key, Value> = {
    c: '唱歌',
    x: '跳舞',
    k: 'rap'
}