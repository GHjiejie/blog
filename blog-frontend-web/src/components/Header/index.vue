<template>
  <div class="headerContainer">
    <div class="left">
      <svg-icon iconClass="icon-blog-logo" className="icon"></svg-icon>
      <h2>Blog</h2>
    </div>
    <div class="navbar">
      <div class="menuList">
        <ul class="menuItem">
          <li v-for="(item, index) in menuList" :key="index">
            <router-link :to="item.path">{{ item.name }}</router-link>
          </li>
        </ul>
      </div>
    </div>
    <div class="right">
      <div class="blogsearch">
        <SearchBlog></SearchBlog>
      </div>
      <div class="useravatar">
        <template v-if="isLogin">
          <el-avatar :src="userAvatar" @click="shwoUserInfo" />
        </template>
        <template v-else>
          <el-avatar
            src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
            @click="userLogin"
          />
        </template>
      </div>
    </div>

    <el-dialog
      v-model="loginVisible"
      title=""
      width="500"
      center
      :show-close="false"
    >
      <Login @loginSuccess="handleLogin"></Login>
    </el-dialog>

    <el-dialog
      v-model="userInfoVisible"
      title=""
      width="500"
      center
      :show-close="false"
    >
      <CurrentUserInfo @logout-success="handleLogout"></CurrentUserInfo>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, ref, watch } from "vue";
import SearchBlog from "@/components/Search/index.vue";
import CurrentUserInfo from "@/components/UserInfo/index.vue";
import Login from "@/pages/Login/index.vue";
import cache from "@/utils/cache";

const loginVisible = ref(false);
const userInfoVisible = ref(false);
const userInfo = ref({});
const menuList = ref([
  { name: "首页", path: "/" },
  { name: "归档", path: "/archive" },
  { name: "关于", path: "/about" },
]);

let userAvatar = cache.sessionGet("avatar");

const isLogin = cache.sessionGet("isLogin")
  ? ref(cache.sessionGet("isLogin"))
  : ref(false);
isLogin.value = cache.sessionGet("isLogin");

const keywords = ref("");

const userLogin = () => {
  loginVisible.value = true;
};

const handleLogin = (user) => {
  isLogin.value = true;
  loginVisible.value = false;
  userInfo.value = user;
  cache.sessionSet("userInfo", user);
  userAvatar = user.avatar;
  // router.push("/");
  // location.reload();
};

const shwoUserInfo = () => {
  userInfoVisible.value = true;
};

const handleLogout = () => {
  isLogin.value = false;
  userInfoVisible.value = false;
};

watch(isLogin, (newVal) => {
  if (newVal) {
    userAvatar = cache.sessionGet("avatar");
  }
});
</script>

<style scoped lang="scss">
:deep(.el-dialog) {
  --el-dialog-bg-color: none;
  --el-dialog-box-shadow: none;
}

.headerContainer {
  position: fixed; // 设置为固定定位
  top: 0; // 固定在顶部
  left: 0; // 从左边开始
  width: 100vw; // 全宽
  overflow: hidden;
  z-index: 1000; // 确保在其他内容上方

  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
  padding: 0 2rem; // 左右内边距
  background-color: #282c34; // 深色背景
  color: white;

  .left {
    display: flex;
    align-items: center;

    .icon {
      width: 40px; // 设置图标大小
      height: 40px;
      margin-right: 0.5rem; // 图标与标题之间的间距
    }

    h2 {
      font-size: 1.5rem;
      margin: 0; // 去掉默认边距
    }
  }

  .navbar {
    .menuList {
      ul.menuItem {
        list-style: none; // 移除列表样式
        display: flex;
        gap: 1.5rem; // 菜单项之间的间距

        li {
          a {
            text-decoration: none; // 不显示下划线
            color: white; // 链接颜色
            padding: 0.5rem 1rem; // 内边距
            transition: background-color 0.3s; // 背景颜色过渡效果

            &:hover {
              background-color: rgba(255, 255, 255, 0.1); // 悬停时的背景颜色
              border-radius: 4px; // 圆角效果
            }
          }
        }
      }
    }
  }

  .right {
    display: flex;
    align-items: center;

    .blogsearch {
      margin-right: 1rem; // 搜索框与头像之间的间距
    }

    .useravatar {
      .el-avatar {
        width: 40px; // 头像大小
        height: 40px;

        &:hover {
          cursor: pointer; // 鼠标悬停时显示手型
        }
      }
    }
  }
}

body {
  padding-top: 70px; // 根据 headerContainer 的高度调整
  background-color: red;
}
</style>
