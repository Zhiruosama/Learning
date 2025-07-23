// 泛型
// 就是动态类型
function test<T>(a: T, b: T): Array<T> {
    return [a, b]
}
// T->number
test(1, 2)
// T->boolean
test(false, true)

type A<T> = string | number | T
let a: A<boolean> = true

interface Data1<T> {
    msg: T
}

let data: Data1<string> = {
    msg: "学习TS"
}
// 泛型也可以有多个
function add<T, K>(a: T, b: K): Array<T | K> {
    return [a, b]
}
add(1, false)
// T=boolean 的话就有了默认值 没有输入的时候就是默认值

const axios = {
    get<T>(url: string): Promise<T> {
        return new Promise((resolve, reject) => {
            let xhr: XMLHttpRequest = new XMLHttpRequest()
            xhr.open('GET', url)
            xhr.onreadystatechange = () => {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    resolve(JSON.parse(xhr.responseText))
                }
            }
            xhr.send(null)
        })
    }
}

interface Data {
    message: string
    code: number
}

axios.get<Data>('./data.json').then(res => {
    console.log(res);
})