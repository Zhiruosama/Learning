// Set Map weakSet weakMap
let set: Set<number> = new Set([1, 2, 3, 4, 5, 6, 6, 6])// set会天然去重 引用类型除外

//map的key可以是引用类型
let obj = { name: "小满" }
let map: Map<object, any> = new Map()
// map.get clear keys forof
map.set(obj, "测试")

//weakmap weakset 弱项 弱引用 不会被计入垃圾回收策略
//跟map的区别在weakmap的key只能是引用类型
let obj1: any = { name: "测试" }//引用次数1
let aahph: any = obj//引用次数2
let weakmap: WeakMap<object, any> = new WeakMap()
weakmap.set(obj, 'sad')//引用次数还是2
obj1 = null;

console.log(aahph);

