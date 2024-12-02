<template>
  <div class="userInfo">
    <div class="Top">
      <div class="user-avatar">
        <el-avatar :size="60"
          src="https://images.pexels.com/photos/189349/pexels-photo-189349.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1" />
      </div>
      <div class="user-details">
        <div class="user-name">{{ userInfo.username }}</div>
        <div class="user-email">{{ userInfo.email }}</div>
      </div>
    </div>

    <div class="media">
      <!-- github -->
      <el-button type="text" icon="el-icon-link" @click="openLink(userInfo.github)">
        Github
      </el-button>

      <el-button type="text" icon="el-icon-link" @click="openLink(userInfo.blog)">
        Gitee
      </el-button>

      <el-button type="text" icon="el-icon-link" @click="openLink(userInfo.weibo)">
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
const props = defineProps({
  userInfo: Object,
});
const userInfo = props.userInfo;

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

    .user-avatar {
      margin-right: 15px;
    }

    .user-details {
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
        content: '';
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