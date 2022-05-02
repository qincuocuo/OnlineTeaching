import { createApp } from "vue";
import App from "./App.vue";
import "normalize.css/normalize.css";
const app = createApp(App);

// 配置信息
import config from "@/config";
window.appConfig = config;
app.config.globalProperties.appConfig = config;

import router from "./router";
import store from "./store";
app.use(store).use(router);
import "@/permission";

import "element-plus/dist/index.css";
import locale from "element-plus/lib/locale/lang/zh-cn";
import ElementPlus from "element-plus";
app.use(ElementPlus, { locale });
// 注册 icons 全局组件
import * as Icons from "@element-plus/icons-vue";
Object.keys(Icons).forEach(key => {
  app.component(key, Icons[key]);
});
// 注册form-create
import formCreate from "@form-create/element-ui";
app.use(formCreate);

// 全局样式
import "@/styles/index.less";

import clickDebounce from "@/directives/clickDebounce";
app.directive("clickdebounce", clickDebounce);
import { has } from "@/directives/has";
app.directive("has", has);

import $eventBus from "@/utils/event";
app.config.globalProperties.$eventBus = $eventBus;

window.app = app;
app.mount("#app");
