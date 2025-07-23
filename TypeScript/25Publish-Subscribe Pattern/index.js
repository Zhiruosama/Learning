"use strict";
// 发布订阅模式
// 发布订阅模式是一种常见的设计模式，在许多场景中都有应用。
// 如使用 addEventListener 方法来监听 DOM 事件、Vue 的事件总线机制等。
class Emitter {
    constructor() {
        this.events = new Map();
    }
    once(event, callback) {
        // 实现逻辑
        //1.创建一个自定义函数通过on触发 触发完以后立马通过off回收
        const cb = (...args) => {
            callback(...args);
            this.off(event, cb);
        };
        this.on(event, cb);
    }
    on(event, callback) {
        //证明存过了
        if (this.events.has(event)) {
            const callbackList = this.events.get(event);
            callbackList && callbackList.push(callback);
        }
        else {
            //否则就是第一次存
            this.events.set(event, [callback]);
        }
    }
    emit(event, ...args) {
        const callbackList = this.events.get(event);
        console.log(callbackList);
        if (callbackList) {
            callbackList.forEach(fn => {
                fn(...args);
            });
        }
    }
    off(event, callback) {
        const callbackList = this.events.get(event);
        console.log(callbackList);
        if (callbackList) {
            callbackList.splice(callbackList.indexOf(callback), 1);
        }
    }
}
const bus = new Emitter();
const fn = (b, n) => {
    console.log(1, b, n);
};
bus.on('message', fn);
// bus.off('message', fn)
// bus.once('message', fn)
bus.emit('message', false, 1);
bus.emit('message', false, 1);
bus.emit('message', false, 1);
