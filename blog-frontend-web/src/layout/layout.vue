<template>
  <el-container>
    <el-header>
      <Header />
    </el-header>
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>
<script setup>
import Header from "@/components/Header/index.vue";
</script>

<style scoped lang="scss">
.el-header,
.el-main {
  padding: 0;
  margin: 0;
}

.el-header {
  height: 50px;
}

.el-main {
}
</style>
