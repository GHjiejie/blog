<template>
  <div class="aside">
    <el-menu router="true">
      <el-menu-item index="/article">
        <el-icon><icon-menu /></el-icon>
        <span>文章管理</span>
      </el-menu-item>
      <el-menu-item index="/user">
        <el-icon><User /></el-icon>
        <span>用户管理</span>
      </el-menu-item>
      <el-menu-item index="/comment">
        <el-icon><ChatLineRound /></el-icon>
        <span>评论管理</span>
      </el-menu-item>
      <el-menu-item index="/file">
        <el-icon><Folder /></el-icon>
        <span>文件管理</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>
<script setup>
import { ref } from "vue";
import {
  Menu as IconMenu,
  ChatLineRound,
  Folder,
  User,
} from "@element-plus/icons-vue";
</script>

<style scoped lang="scss">
.el-menu {
  border-right: none;
}
.aside {
}
</style>
