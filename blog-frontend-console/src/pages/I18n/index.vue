<template>
  <div>
    <h1>{{ $t("welcome") }}</h1>
    <button @click="switchLanguage('en')">English</button>
    <button @click="switchLanguage('zh')">中文</button>
  </div>
</template>

<script setup>
import { useI18n } from "vue-i18n";

// 使用 useI18n 获取 i18n 实例
const { locale } = useI18n();

// 切换语言的方法
const switchLanguage = (lang) => {
  locale.value = lang; // 设置当前语言
};
</script>
