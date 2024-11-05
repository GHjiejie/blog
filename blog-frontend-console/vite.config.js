import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],

  // 设置代理服务器
  server: {
    proxy: {
      "/v1": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
  // 设置@别名
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
