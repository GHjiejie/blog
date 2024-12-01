import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import SvgIcon from "@/components/SvgIcon/index.vue";
import "virtual:svg-icons-register";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "@/styles/main.scss";
const app = createApp(App);
app.use(router);
app.component("svg-icon", SvgIcon);
app.use(ElementPlus);

app.mount("#app");
