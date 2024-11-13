<template>
  <div class="aside">
    <el-menu router="true" default-active="/article/articleList">
      <!-- <el-menu-item index="/article">
        <el-icon><icon-menu /></el-icon>
        <span>文章管理</span>
      </el-menu-item> -->
      <el-sub-menu index="/article">
        <template #title>
          <el-icon>
            <Document />
          </el-icon>
          <span>文章管理</span>
        </template>

        <el-menu-item index="/article/articleList">
          <template #title>
            <el-icon>
              <List />
            </el-icon>
            <span>文章列表</span>
          </template>

        </el-menu-item>
        <el-menu-item index="/article/articlePublish">
          <template #title>
            <el-icon>
              <UploadFilled />
            </el-icon>
            <span>发布文章</span>
          </template>
        </el-menu-item>

      </el-sub-menu>
      <el-menu-item index="/user">
        <el-icon>
          <User />
        </el-icon>
        <span>用户管理</span>
      </el-menu-item>
      <el-menu-item index="/comment">
        <el-icon>
          <ChatLineRound />
        </el-icon>
        <span>评论管理</span>
      </el-menu-item>
      <el-menu-item index="/file">
        <el-icon>
          <Folder />
        </el-icon>
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
  Document,
  List,
  UploadFilled
} from "@element-plus/icons-vue";
</script>

<style scoped lang="scss">
.el-menu {
  border-right: none;
}

.aside {}
</style>
