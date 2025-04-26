<template>
  <div class="aside">
    <div class="menu">
      <div class="menu-item">
        <a href="/articleManage" class="menu-link">
          <i class="el-icon-document"></i>
          <span>文章管理</span>
        </a>
        <div class="sub-menu">
          <a href="/articleManage/articleList" class="menu-link">
            <i class="el-icon-list"></i>
            <span>文章列表</span>
          </a>
          <a href="/articleManage/articlePublish" class="menu-link">
            <i class="el-icon-upload-filled"></i>
            <span>发布文章</span>
          </a>
        </div>
      </div>

      <div class="menu-item" v-if="cache.sessionGet('userRole') === 'ADMIN'">
        <a href="/user" class="menu-link">
          <i class="el-icon-user"></i>
          <span>用户管理</span>
        </a>
      </div>

      <div class="menu-item" v-if="cache.sessionGet('userRole') === 'ADMIN'">
        <a href="/comment" class="menu-link">
          <i class="el-icon-chat-line-round"></i>
          <span>评论管理</span>
        </a>
      </div>

      <div class="menu-item">
        <a href="/tag" class="menu-link">
          <i class="el-icon-collection-tag"></i>
          <span>标签管理</span>
        </a>
      </div>

      <div class="menu-item">
        <a href="/file" class="menu-link">
          <i class="el-icon-folder"></i>
          <span>文件管理</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import cache from "@/utils/cache";
// Import icons directly from the Element Plus CSS
// Use the i tag and the correct class names (prefixed with el-icon)
</script>
