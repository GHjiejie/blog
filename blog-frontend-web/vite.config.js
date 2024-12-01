import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    createSvgIconsPlugin({
      iconDirs: [path.resolve(process.cwd(), "src/assets/icons")], // icon存放的目录
      symbolId: "icon-[name]", // symbol的id,简单的说就是svg文件命名的格式
      inject: "body-last", // 插入的位置
      customDomId: "__svg__icons__dom__", // svg的id
    }),
  ],

  server: {
    proxy: {
      "/v1": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },

  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
  // 处理css资源
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "@/styles/themes.scss";`,
      },
    },
  },
});
