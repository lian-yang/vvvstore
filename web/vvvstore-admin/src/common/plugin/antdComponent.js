/**
 * 插件化加载Antd组件、简化main.js需要单独一个一个的去 app.use(Component)
 */

import {
    Button,
    message
} from 'ant-design-vue';

const components = {
    Button
}


export default {
    install: app => {
        app.config.globalProperties.$message = message;
        Object.keys(components).forEach(key => app.use(components[key]));
    }
}