<template>
  <div class="container">
    <div class="Top">
      <button class="addserBtn" @click="openAddUserDialog">
        <el-icon><Plus /></el-icon>
      </button>
    </div>
    <div class="Center">
      <el-table
        :data="filterTableData"
        size="large"
        :height="maxHeight"
        :highlight-current-row="true"
      >
        <el-table-column
          fixed="left"
          prop="userId"
          label="用户ID"
          width="120"
        />
        <el-table-column prop="username" label="用户名" width="200" />
        <el-table-column prop="role" label="角色" width="150">
          <template #default="{ row }">
            <el-tag size="small" round>{{ row.role }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="status" label="状态" width="120">
          <template #default="{ row }">
            <el-tag v-if="row.status === 'NORMAL'" type="success" round
              >正常</el-tag
            >
            <el-tag v-else type="danger" round>禁用</el-tag>
          </template>
        </el-table-column>
        <!-- <el-table-column prop="email" label="邮箱" width="120" />
        <el-table-column prop="phone" label="手机号" width="120" /> -->

        <el-table-column prop="createdAt" label="创建时间" width="300" />
        <el-table-column fixed="right" min-width="150">
          <template #header>
            <el-input
              v-model="search"
              size="small"
              placeholder="Type to search"
            />
          </template>
          <template #default="{ row }">
            <el-button
              :disabled="row.role === 'ADMIN'"
              link
              type="danger"
              size="small"
              @click="handelDelUser(row, row.userId)"
            >
              删除
            </el-button>
            <el-button
              :disabled="row.role === 'ADMIN'"
              link
              type="warning"
              size="small"
              @click="handlePWDReset(row.userId)"
              >密码重置</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="Footer">
      <el-button type="primary" @click="handleNextPage()" size="small"
        >加载更多</el-button
      >
    </div>
  </div>
  <AddUser
    ref="addUserRef"
    :user-id="userId"
    @success="handleAddUser"
  ></AddUser>
</template>

<script setup>
import AddUser from "./components/addUser.vue";
import { onMounted, computed, ref } from "vue";
import {
  getUserList,
  deleteUser,
  resetPassword,
  registerUser,
} from "@/apis/user";
import { ElMessage, ElMessageBox } from "element-plus";
import { Plus } from "@element-plus/icons-vue";

let page = ref(1);
let pageSize = ref(20);
let UserListData = ref([]);
const search = ref("");
let addUserRef = ref(null);
const userId = ref("");
const userCount = ref(0);
const maxHeight = window.innerHeight - 100;

onMounted(async () => {
  await getUserListData();
});

// 表单过滤操作
const filterTableData = computed(() =>
  UserListData.value.filter((item) => {
    return (
      item.username.toLowerCase().includes(search.value) ||
      item.role.toLowerCase().includes(search.value) ||
      item.userId.toString().includes(search.value)
    );
  })
);

// 获取用户列表
const getUserListData = async () => {
  try {
    const res = await getUserList({
      page: page.value,
      pageSize: pageSize.value,
    });
    if (res.data.users.length === 0) {
      ElMessage.info("没有更多数据了");
      return;
    } else {
      UserListData.value = [...UserListData.value, ...res.data.users];
      userCount.value = res.data.total;
    }
  } catch (error) {}
};
// 用户密码重置
const handlePWDReset = async (userId) => {
  ElMessageBox.confirm("此操作将重置该用户密码, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      // 重置密码
      try {
        const res = await resetPassword({
          userId,
        });
        if (res.status === 200) {
          ElMessage.success("重置成功");
        }
      } catch (error) {}
    })
    .catch(() => {
      ElMessage.info("已取消重置");
    });
};

// 删除用户
const handelDelUser = async (rowInfo, userId) => {
  ElMessageBox.confirm("此操作将永久删除该用户, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      // 删除用户
      try {
        const res = await deleteUser({
          userId,
        });
        if (res.status === 200) {
          ElMessage.success("删除成功");
          // 删除表里面ID为userId的数据
          UserListData.value = UserListData.value.filter(
            (item) => item.userId !== userId
          );
        }
      } catch (error) {}
    })
    .catch(() => {
      ElMessage.info("已取消删除");
    });
};

// 添加用户
const handleAddUser = async (formData) => {
  await registerUser({
    username: formData.username,
    role: formData.role,
    password: formData.password,
  });
  ElMessage.success("添加成功");
  addUserRef.value.changeVisible(false);
  await getUserListData();
};

// 打开添加用户弹窗
const openAddUserDialog = async () => {
  addUserRef.value.changeVisible(true);
};

// 加载更多
const handleNextPage = async () => {
  page.value++;
  await getUserListData();
};
</script>
<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100%;
  .Top {
    position: fixed;
    top: 50%;
    right: 2%;
    z-index: 999;
    .addserBtn {
      background-color: $--color-primary-opacity7;
      color: $--color-text;
      border: none;
      padding: 10px;
      border-radius: 54% 46% 21% 79% / 52% 46% 54% 48%;

      box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
      transition: all 0.3s;
      cursor: pointer;
      &:hover {
        background-color: $--color-primary;
      }
    }
  }

  .Footer {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 50px;
    transition: all 0.3s;
    .el-button {
      &:hover {
        color: $--color-text3;
        background-color: $--color-primary;
      }
    }
  }
}
</style>
