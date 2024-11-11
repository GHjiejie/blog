import { createApp } from "vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import App from "./App.vue";
import "@/styles/main.scss";
import router from "./router/index.js";
import plugins from "@/plugins";
import JsonViewer from "vue3-json-viewer";
// if you used v1.0.5 or latster ,you should add import "vue3-json-viewer/dist/index.css"
import "vue3-json-viewer/dist/index.css";
const app = createApp(App);
app.use(router);
app.use(plugins);
app.use(JsonViewer);
app.use(ElementPlus);
app.mount("#app");
