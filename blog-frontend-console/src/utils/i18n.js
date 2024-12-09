// src/i18n.js
import { createI18n } from "vue-i18n";

// 语言包
const messages = {
  en: {
    welcome: "Welcome to our application!",
  },
  zh: {
    welcome: "欢迎使用我们的应用!",
  },
};

// 创建 i18n 实例
const i18n = createI18n({
  locale: "en", // 默认语言
  messages, // 语言包
});

export default i18n;
