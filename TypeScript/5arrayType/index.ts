// 数组类型
// 数组普通类型
let arr1: number[] = [1, 2, 3, 4, 5]
let arr2: Array<number> = [1, 2, 3, 4, 5]

// interface定义对象数组
interface X {
    name: string,
    age?: number
}
let a: X[] = [{ name: '老王' }, { name: '老屋' }]

// 二维数组
let b: number[][] = [[1], [2], [3]]

function c(...args:any[]){
    console.log(args)
}

c(1,2,3)