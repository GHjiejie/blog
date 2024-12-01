<template>
  <div class="search-container">
    <el-input v-model="keywords" style="width: 240px" placeholder="Crtl + K" :prefix-icon="Search" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
const keywords = ref('')
</script>

<style scoped lang="scss">
.el-input {
  // border-radius: 50%;
  // border: 1px solid #409EFF;
  --el-border-radius-base: 50px;
}
</style>