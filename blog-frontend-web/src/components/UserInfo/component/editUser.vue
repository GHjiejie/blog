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
        <el-avatar :size="100" :src="src" @click="handleAvatar" />
      </div>
      <el-form-item label="用户名" prop="username">
        <el-input v-model="editUserForm.username" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="editUserForm.email" />
      </el-form-item>

      <el-form-item label="手机号" prop="phone">
        <el-input v-model="editUserForm.phone" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="editUser">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getUserById, updateUser } from "@/apis/user";
import cache from "@/utils/cache";

const avatarRef = ref(null);
const src = ref(cache.sessionGet("avatar"));

const editUserForm = ref({
  username: "",
  email: "",
  userId: "",
  phone: "",
});
const editUserRules = ref({
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  email: [{ required: true, message: "请输入邮箱", trigger: "blur" }],
  github: [{ required: true, message: "请输入手机号", trigger: "blur" }],
});

const handleAvatar = () => {
  avatarRef.value.click();
};

const avatarChange = (e) => {
  const file = e.target.files[0];
  const reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = (e) => {
    src.value = e.target.result;
  };
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

const editUser = async () => {
  try {
    // 创建formdata对象
    const formData = new FormData();
    // 添加数据
    formData.append("userId", cache.sessionGet("userId"));
    formData.append("username", editUserForm.value.username);
    formData.append("email", editUserForm.value.email);
    formData.append("avatar", avatarRef.value.files[0]);
    formData.append("phone", editUserForm.value.phone);

    const res = await updateUser(formData);
    console.log("用户修改信息", res);
    if (res.status == 200) {
      // cache.sessionSet("avatar", res.data.user.avatar);
      // cache.sessionSet("username", res.data.user.username);
      // cache.sessionSet("email", res.data.user.email);
      ElMessage.success("修改成功");
    } else {
      ElMessage.error("修改失败");
    }
  } catch (error) {}
};
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
