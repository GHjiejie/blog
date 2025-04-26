<template>
  <div class="layout">
    <el-container>
      <!-- <el-header>
        <Header />
      </el-header> -->
      <el-container>
        <el-aside>
          <Aside />
        </el-aside>
        <!-- <el-main>
          <router-view />
        </el-main> -->
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Aside from "./components/Aside.vue";
// import Header from "@/components/Header/index.vue";
</script>

<style scoped lang="scss">
.el-header,
.el-aside,
.el-main {
  padding: 0;
}

.layout {
  display: flex;
}

.el-header {
  height: 50px;
  border-bottom: 1px solid var(--el-border-color);
}

.el-container {
  flex: 1;

  .el-aside {
    width: 400px;
    // border-right: 1px solid var(--el-border-color);
  }
}
</style>
