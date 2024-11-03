<template>
  <div class="layout">
    <el-container>
      <el-header><Header /></el-header>
      <el-container>
        <el-aside><Aside /></el-aside>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Aside from "@/components/Aside/index.vue";
import Header from "@/components/Header/index.vue";
</script>

<style scoped lang="scss">
.el-header,
.el-aside,
.el-main {
  padding: 0;
}

.el-header {
  height: 5vh;
  border-bottom: 1px solid var(--el-border-color);
}
.el-aside {
  height: 95vh;
  width: 15vw;
}
.el-main {
  height: 95vh;
}
</style>
