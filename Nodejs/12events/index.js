const eventEmitter = require('events')

//发布订阅模式
//off on emit once核心api

const bus = new eventEmitter()
const fn = (name => {
    console.log(name);
})
//修改最大事件数
bus.setMaxListeners(20)

//订阅一个事件
//事件默认只能监听10个
bus.on('test', fn)
//发布
bus.emit('test', 'ceshi1')
//删除一个事件
bus.off('test', fn)
bus.emit('test', 'ceshi2')
//如果on改成once 就只能触发一次
