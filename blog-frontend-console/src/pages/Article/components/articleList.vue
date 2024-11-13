<template>
  <div class="container">
    <div class="Top">

    </div>
    <div class="Center">

    </div>
    <div class="Footer">
      <el-button type="primary" @click="handleNextPage()" size="small">加载更多</el-button>
    </div>
  </div>
  <fileView ref="fileViewRef" :file-id="viewFileId"></fileView>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { getArticleList } from '@/apis/articles'
const page = ref(1);
const pageSize = ref(10);
onMounted(async () => {
  await getArticleListData()
})

// 获取文章列表
const getArticleListData = async () => {
  const res = await getArticleList({ page: page.value, pageSize: pageSize.value })
  console.log(res)
}

</script>
<style scoped lang="scss"></style>