// 元组就是数组的一个变种
// 不同就不同在元组中的元素类型可以不同，而且数量固定

let arr = [1, false] // 联合类型
let arr1: [number, boolean] = [1, false] // 元组
// arr1.push(null) 报错 因为定义无这种类型

// 只有readonly才能让元组成为只读
let arr2: readonly [x: number, y?: boolean] = [1, false]

let excel:[string,string,number][]=[
    ['人','男',1],
    ['人','男',1],
    ['人','男',1],
]
