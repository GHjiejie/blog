// import { VMdEditor, VMdPreview } from './mdEditor'
import VMdPreview from "./mdEditor";

const pluginList = [VMdPreview];
const plugins = {
  install(app) {
    // 批量注册插件  改用自动引入
    Object.entries(pluginList).forEach(([, v]) => {
      app.use(v);
    });
  },
};

export default plugins;