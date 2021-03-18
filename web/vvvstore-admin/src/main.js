import { createApp } from "vue";
import 'ant-design-vue/dist/antd.css';

import App from "./App.vue";
import router from "./router";
import store from "./store";

import {
  antdComponent,
  eventBus
} from './common/plugin';

const app = createApp(App)

app.use(store)
  .use(router)
  .use(antdComponent)
  .use(eventBus);

app.mount("#app");
