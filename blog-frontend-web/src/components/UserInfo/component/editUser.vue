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
          src="https://images.pexels.com/photos/920220/pexels-photo-920220.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
          @click="handleAvatar"
        />
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
import { ref, onMounted, defineEmits } from "vue";
import { getUserById, updateUser } from "@/apis/user";
import cache from "@/utils/cache";
import { getFileUrl } from "@/utils/fileFilter";
const emits = defineEmits(["updateSuccess"]);
const avatarRef = ref(null);
const src = ref("");
src.value = cache.sessionGet("userAvatar");

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
    formData.append(
      "avatar",
      avatarRef.value.files[0] !== undefined ? avatarRef.value.files[0] : null
    );
    formData.append("phone", editUserForm.value.phone);

    const res = await updateUser(formData);
    console.log("resjkfjs", res);
    emits("updateSuccess");
    console.log("用户修改信息", res);
    cache.sessionSet("userAvatar", getFileUrl(res.data.avatar, "image/jpeg"));
    console.log("用户头像", res);
    cache.sessionSet("username", res.data.username);
    cache.sessionSet("userEmail", res.data.email);
    cache.sessionSet("userPhone", res.data.phone);
  } catch (error) {}
};
</script>

<style scoped lang="scss">
.editUser {
  background-color: var(--color-background);
  display: flex;
  justify-content: center;
  height: 400px;

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
