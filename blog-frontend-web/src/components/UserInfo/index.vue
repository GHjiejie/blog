<template>
  <div class="userInfo">
    <div class="Top">
      <div class="user-baseinfo">
        <div class="user-details">
          <el-avatar :size="60" :src="src" />
          <div class="user-name">{{ userInfo.username }}</div>
        </div>
        <div class="operation">
          <!-- <el-button type="primary" @click="handleLogout"> 登出 </el-button> -->
          <el-dropdown placement="bottom">
            <el-button> 操作 </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="handleLogout">登出</el-dropdown-item>
                <el-dropdown-item @click="editUser">编辑</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </div>

    <!-- <div class="media">

      <el-button type="text" icon="el-icon-link" @click="openLink(userInfo.github)">
        Github
      </el-button>

      <el-button type="text" icon="el-icon-link" @click="openLink(userInfo.blog)">
        Gitee
      </el-button>

      <el-button type="text" icon="el-icon-link" @click="openLink(userInfo.weibo)">
        WeChat
      </el-button>
    </div> -->

    <div class="Footer">
      <div class="user-id">ID: {{ userInfo.userId }}</div>
      <div class="user-role">Role:{{ userInfo.role }}</div>
      <div class="user-email">Email:{{ userInfo.email }}</div>
      <div class="user-phone">Phone:{{ userInfo.phone }}</div>
    </div>

    <el-dialog v-model="editUserdVisible" title="" width="40%" :show-close="false">
      <editUserPanel @updateSuccess="handleUpdateSuccess" />
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, defineEmits, watch, onMounted } from "vue";
import cache from "@/utils/cache";
import { logout, getUserById } from "@/apis/user.js";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { getFileUrl } from "@/utils/fileFilter";
import editUserPanel from "./component/editUser.vue";
const emits = defineEmits(["logoutSuccess"]);
const editUserdVisible = ref(false);
const userInfo = ref({});
// const userAvatar = getFileUrl(userInfo.avatar, "image/jpeg");


const userId = cache.sessionGet("userId");
const src = ref('')

onMounted(async () => {
  await getUser()
});

const getUser = async () => {
  const res = await getUserById({ userId });
  userInfo.value = res.data.user;
  src.value = getFileUrl(userInfo.value.avatar, 'image/jpeg')
}
// 登出操作
const handleLogout = async () => {
  try {
    const res = await logout({ userId });
    console.log(res);
    if (res.status == 200) {
      ElMessage.success("登出成功!");
      cache.removeToken();
      cache.sessionRemove("userInfo");
      cache.sessionRemove("userId");
      cache.sessionRemove("userRole");
      cache.sessionRemove("username");
      cache.sessionRemove("avatar");
      cache.sessionRemove("isLogin");
      emits("logoutSuccess");
    } else {
      ElMessage.error("登出失败");
    }
  } catch (error) {
    console.log(error);
  }
};

// 编辑用户操作
const editUser = () => {
  editUserdVisible.value = true;
};

const handleUpdateSuccess = async () => {
  console.log("更新成功");
  editUserdVisible.value = false;
  // location.reload();
  await getUser();
  location.reload()
};


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
