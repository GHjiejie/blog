<template>
  <div class="aside-menu">
    <div class="menu-group">
      <h3 class="menu-title">操作列表</h3>
      <div class="menu-item">
        <router-link to="/articleManage/articleList" class="menu-link">
          <i class="el-icon-document"></i>
          <span>文章列表</span>
        </router-link>
      </div>
      <div class="menu-item">
        <router-link to="/articleManage/articlePublish" class="menu-link">
          <i class="el-icon-upload"></i>
          <span>发布文章</span>
        </router-link>
      </div>
      <div class="menu-item">
        <router-link to="/tag" class="menu-link">
          <i class="el-icon-collection-tag"></i>
          <span>标签管理</span>
        </router-link>
      </div>
      <div class="menu-item">
        <router-link to="/file" class="menu-link">
          <i class="el-icon-folder"></i>
          <span>文件管理</span>
        </router-link>
      </div>
    </div>

    <div class="menu-group" v-if="cache.sessionGet('userRole') === 'ADMIN'">
      <h3 class="menu-title">系统管理</h3>
      <div class="menu-item">
        <router-link to="/user" class="menu-link">
          <i class="el-icon-user"></i>
          <span>用户管理</span>
        </router-link>
      </div>
      <div class="menu-item">
        <router-link to="/comment" class="menu-link">
          <i class="el-icon-chat-dot-round"></i>
          <span>评论管理</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import cache from "@/utils/cache";
</script>
