<template>
  <h1>Hello</h1>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getArticleList } from '@/apis/articles'


const page = ref(1)
const pageSize = ref(10)
onMounted(async () => {
  const res = await getArticleList({ page: page.value, pageSize: pageSize.value })
  console.log(res)
})
</script>
