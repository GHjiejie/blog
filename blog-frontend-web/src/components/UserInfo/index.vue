<template>
  <div class="userInfo">
    <div class="Top">
      <div class="user-baseinfo">
        <div class="user-details">
          <el-avatar :size="60" :src="userInfo.avatar" />
          <div class="user-name">{{ userInfo.username }}</div>
        </div>
        <div class="logout">
          <el-button type="primary" @click="logout"> 登出 </el-button>
        </div>
      </div>
    </div>

    <div class="media">
      <!-- github -->
      <el-button
        type="text"
        icon="el-icon-link"
        @click="openLink(userInfo.github)"
      >
        Github
      </el-button>

      <el-button
        type="text"
        icon="el-icon-link"
        @click="openLink(userInfo.blog)"
      >
        Gitee
      </el-button>

      <el-button
        type="text"
        icon="el-icon-link"
        @click="openLink(userInfo.weibo)"
      >
        WeChat
      </el-button>
    </div>

    <div class="Footer">
      <div class="user-id">ID: {{ userInfo.userId }}</div>
      <div class="user-role">Role:{{ userInfo.role }}</div>
      <div class="user-email">Email:3426571530@qq.com</div>
      <div class="user-phone">Phone:18196576670</div>
    </div>
  </div>
</template>
<script setup>
import cache from "@/utils/cache";
const userInfo = cache.sessionGet("userInfo");
console.log(userInfo);
</script>

<style scoped lang="scss">
.userInfo {
  background-color: var(--color-background);

  padding: 20px;
  max-width: 400px;

  .Top {
    display: flex;
    align-items: center;
    margin-bottom: 15px;

    .user-baseinfo {
      width: 100%;
      margin-right: 15px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 10px;
    }

    .user-details {
      display: flex;
      align-items: center;
      gap: 10px;
      .user-name {
        font-size: 18px;
        font-weight: bold;
        color: $--color-text3;
      }

      .user-email {
        font-size: 14px;
        color: $--color-text3;
      }
    }
  }

  .media {
    display: flex;
    justify-content: space-around;
    margin-bottom: 15px;

    .el-button {
      color: var(--el-color-primary);
      font-size: 14px;
      transition: color 0.3s;

      &:hover {
        color: #2980b9;
      }
    }
  }

  .Footer {
    border-top: 1px solid #eaeaea;
    padding-top: 10px;
    font-size: 14px;
    color: #555;

    .user-id,
    .user-role,
    .user-email,
    .user-phone {
      margin: 5px 0;

      &::before {
        content: "";
        display: inline-block;
        width: 12px;
        height: 12px;
        margin-right: 5px;
        vertical-align: middle;
        background-color: var(--el-color-primary);
        border-radius: 50%;
      }
    }

    .user-role {
      font-weight: bold;
    }
  }
}
</style>
