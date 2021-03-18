/**
 * 全局挂载事件总线
 * 
 */
import mitt from 'mitt'

const emitter = mitt();

export const $event = emitter;

export default {
    install: app => {
        app.config.globalProperties.$event = emitter;
    }
}




/**
// 监听事件
this.$event.on('foo', msg => console.log(msg));

// 发送事件
this.$event.emit('foo', 'send foo event');

// 移除事件
this.$event.off('foo', () => console.log('remove foo event'));
 */