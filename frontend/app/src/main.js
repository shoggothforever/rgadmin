import { createApp } from 'vue';
import { createPinia } from 'pinia';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import '@/styles/index.scss'
import './permission'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue';
import router from './router';
import echarts from './common/js/echart.js'

const app = createApp(App);

export const pinia = createPinia();

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.config.globalProperties.$echart = echarts;

app.use(pinia);
app.use(router);
app.use(ElementPlus);

app.mount('#app');
