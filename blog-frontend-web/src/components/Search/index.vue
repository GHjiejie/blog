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
const page = ref(1)
const pageSize = ref(10)
const searchResults = ref([])

// 监听用户按下的键盘事件，当用户按下 Ctrl + K 时，触发搜索事件：

const performSearch = async () => {
  const res = await queryArticle({
    keyword: keywords.value,
    page: page.value,
    pageSize: pageSize.value
  })
  searchResults.value = res.data
}



const handleSearch = async () => {
  await performSearch()
  console.log('Search results:', searchResults.value)
}


</script>

<style scoped lang="scss">
.el-input {
  // border-radius: 50%;
  // border: 1px solid #409EFF;
  --el-border-radius-base: 50px;
}
</style>