<template>
  <div class="search-container">
    <el-input v-model="keywords" style="width: 240px" placeholder="输入关键字" :prefix-icon="Search" @click="openPanel" />
    <el-dialog v-model="searchPanelStatus" title="" width="800" :show-close="false" top="0">
      <search-result></search-result>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'

import SearchResult from './component/searchResult.vue'
const keywords = ref('')

const searchPanelStatus = ref(false)



const openPanel = () => {
  searchPanelStatus.value = true
}




</script>

<style scoped lang="scss">
.el-input {
  // border-radius: 50%;
  // border: 1px solid #409EFF;
  --el-border-radius-base: 50px;
}
</style>