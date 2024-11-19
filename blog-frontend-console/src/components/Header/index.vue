<template>
  <div class="header">
    <div class="left">
      <img src="../../assets/svg/logo.svg" alt="" height="30px" />
      <span>控制台</span>
    </div>
    <!-- <div class="center">博客管理系统</div> -->
    <div class="right">
      <div class="username">{{ username }}</div>
      <div class="logout">
        <el-button type="primary" link size="small" @click="handleLogout"><span>登出</span></el-button>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import cache from "@/utils/cache";
import { logout } from "@/apis/user";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";
const router = useRouter();

const username = ref("");
username.value = cache.sessionGet("user");
const userId = cache.sessionGet("userId");
console.log("userId", userId);
const handleLogout = async () => {
  try {
    const res = await logout(userId);
    if (res.status == 200) {
      ElMessage.success("登出成功, 即将跳转登录页面...");
      setTimeout(() => {
        cache.removeToken();
        cache.sessionRemove("userId");
        cache.sessionRemove("user");
        router.push("/login");
      }, 500);
    } else {
      ElMessage.error("登出失败");
    }
  } catch (error) { }
};
</script>
<style scoped lang="scss">
.header {
  height: 100%; //需要继承父元素的高度
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: $--bg-color;
  border-bottom: 2px solid $--color-border;


  .left {
    flex: 1;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 20px;

    span {
      font-size: 20px;
      font-weight: bold;
      color: $--color-text2;
    }
  }

  .right {
    flex: 1;
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding-right: 20px;

    .username {
      font-weight: bold;
      margin-right: 10px;
      color: $--color-text2;

      &:hover {
        cursor: pointer;
      }
    }

    .logout {
      .el-button {
        border: none;
        padding: 5px 10px;
        border-radius: 5px;
        color: $--color-primary;

        span {
          font-weight: bold;
        }

        &:hover {
          background-color: $--color-primary-opacity7;
        }
      }
    }
  }
}
</style>
