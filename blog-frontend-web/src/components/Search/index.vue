<template>
  <div class="search-container">
    <el-input v-model="keywords" style="width: 240px" placeholder="输入关键字" :prefix-icon="Search"
      @keyup.enter="handleSearch" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { queryArticle } from '@/apis/articles'
const keywords = ref('')

// 监听用户按下的键盘事件，当用户按下 Ctrl + K 时，触发搜索事件：
const handleSearch = async () => {
  const res = await queryArticle({ keyword: keywords.value })
}
</script>

<style scoped lang="scss">
.el-input {
  // border-radius: 50%;
  // border: 1px solid #409EFF;
  --el-border-radius-base: 50px;
}
</style>