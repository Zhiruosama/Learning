// 发布订阅模式
// 发布订阅模式是一种常见的设计模式，在许多场景中都有应用。
// 如使用 addEventListener 方法来监听 DOM 事件、Vue 的事件总线机制等。

// // 监听器
// const cb = () => {
//     console.log('触发了');
// }
// document.addEventListener('asd', cb, {
//     once: true // 只触发一次
// })
// document.removeEventListener('asd', cb) //支持删除监听器函数
// const e = new Event('asd')//订阅中心
// document.dispatchEvent(e)

// 实现 once on emit off 订阅中心Map<事件名称，[Function]订阅者集合>
interface I {
    events: Map<string, Function[]>
    once: (event: string, callback: Function) => void //触发一次订阅器
    on: (event: string, callback: Function) => void //订阅
    emit: (event: string, ...args: any[]) => void //派发
    off: (event: string, callback: Function) => void //删除监听器
}
class Emitter implements I {
    events: Map<string, Function[]>
    constructor() {
        this.events = new Map()
    }
    once(event: string, callback: Function) {
        // 实现逻辑
        //1.创建一个自定义函数通过on触发 触发完以后立马通过off回收
        const cb = (...args: any) => {
            callback(...args)
            this.off(event, cb)
        }
        this.on(event, cb)
    }
    on(event: string, callback: Function) {
        //证明存过了
        if (this.events.has(event)) {
            const callbackList = this.events.get(event)
            callbackList && callbackList.push(callback)
        } else {
            //否则就是第一次存
            this.events.set(event, [callback])
        }
    }
    emit(event: string, ...args: any[]) {
        const callbackList = this.events.get(event)
        console.log(callbackList);
        if (callbackList) {
            callbackList.forEach(fn => {
                fn(...args)
            })
        }
    }
    off(event: string, callback: Function) {
        const callbackList = this.events.get(event)
        console.log(callbackList);
        if (callbackList) {
            callbackList.splice(callbackList.indexOf(callback), 1)
        }
    }
}

const bus = new Emitter()
const fn = (b: boolean, n: number) => {
    console.log(1, b, n);
}
bus.on('message', fn)
// bus.off('message', fn)
// bus.once('message', fn)
bus.emit('message', false, 1)
bus.emit('message', false, 1)
bus.emit('message', false, 1)