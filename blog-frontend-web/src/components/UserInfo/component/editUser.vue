<template>
  <div class="editUser">
    <el-form
      :model="editUserForm"
      :rules="editUserRules"
      ref="editUserFormRef"
      label-width="80px"
    >
      <!-- 用户头像 -->
      <div class="avatar">
        <input
          type="file"
          name=""
          id=""
          @change="avatarChange"
          accept="image/*"
          ref="avatarRef"
          style="display: none"
        />
        <el-avatar
          :size="100"
          :src="editUserForm.avatar"
          @click="handleAvatar"
        />
      </div>
      <el-form-item label="用户名" prop="username">
        <el-input v-model="editUserForm.username" disabled />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="editUserForm.email" />
      </el-form-item>

      <el-form-item label="Github" prop="github">
        <el-input v-model="editUserForm.github" />
      </el-form-item>
      <el-form-item label="Gitee" prop="blog">
        <el-input v-model="editUserForm.blog" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="editUser">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getUserById } from "@/apis/user";
import { ElMessage } from "element-plus";
import cache from "@/utils/cache";
import { useRouter } from "vue-router";

const avatarRef = ref(null);

const editUserForm = ref({});
const editUserRules = ref({
  email: [
    { required: true, message: "请输入邮箱", trigger: "blur" },
    {
      type: "email",
      message: "请输入正确的邮箱地址",
      trigger: ["blur", "change"],
    },
  ],
  github: [{ required: false, message: "请输入Github地址", trigger: "blur" }],
  blog: [{ required: false, message: "请输入Gitee地址", trigger: "blur" }],
});

const handleAvatar = () => {
  avatarRef.value.click();
};

onMounted(async () => {
  try {
    const { data } = await getUserById({
      userId: cache.sessionGet("userId"),
    });
    editUserForm.value = data.user;
  } catch (error) {
    console.error("获取用户信息失败", error);
  }
});

const editUser = async () => {};
</script>

<style scoped lang="scss">
.editUser {
  background-color: var(--color-background);
  display: flex;
  justify-content: center;
  .avatar {
    display: flex;
    justify-content: center;
    margin-bottom: 20px;

    .el-avatar {
      cursor: pointer;
      border: 2px solid #dcdfe6;
      transition: border-color 0.3s;

      &:hover {
        border-color: #409eff;
      }
    }
  }

  .el-form-item {
    margin-bottom: 20px;
    width: 100%;

    .el-input {
      width: 100%;
    }
  }

  .el-button {
    width: 100%;
  }
}
</style>
