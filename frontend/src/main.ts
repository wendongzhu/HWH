import {createApp} from 'vue'
import App from './App.vue'
import router from "./router/index";
import ElementPlus from "element-plus";
import * as ElIcons from '@element-plus/icons-vue'
import * as es from 'echarts';
import "element-plus/dist/index.css";
import "font-awesome/scss/font-awesome.scss";
import mitt from "mitt";

const app = createApp(App)
app.use(router)
app.use(ElementPlus)
for (const name in ElIcons){
    app.component(name,(ElIcons as any)[name])
}
app.config.globalProperties.$echarts = es
app.config.globalProperties.$mybus = mitt()

app.mount("#app")