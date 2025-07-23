"use strict";
// 泛型
// 就是动态类型
function test(a, b) {
    return [a, b];
}
// T->number
test(1, 2);
// T->boolean
test(false, true);
let a = true;
let data = {
    msg: "学习TS"
};
// 泛型也可以有多个
function add(a, b) {
    return [a, b];
}
add(1, false);
// T=boolean 的话就有了默认值 没有输入的时候就是默认值
const axios = {
    get(url) {
        return new Promise((resolve, reject) => {
            let xhr = new XMLHttpRequest();
            xhr.open('GET', url);
            xhr.onreadystatechange = () => {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    resolve(JSON.parse(xhr.responseText));
                }
            };
            xhr.send(null);
        });
    }
};
axios.get('./data.json').then(res => {
    console.log(res);
});
