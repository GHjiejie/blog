<template>
  <div class="login-container">
    <div class="login-card">
      <h1>登录</h1>
      <input type="text" v-model="userForm.username" placeholder="Username" />
      <input type="password" v-model="userForm.password" placeholder="Password" />
      <button @click="handleLogin">登录</button>
      <el-button @click="handleRegister" link>注册</el-button>
    </div>
  </div>
</template>

<script setup>
import { reactive, defineEmits } from "vue";
import { login } from "@/apis/user.js";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import cache from "@/utils/cache";
import { getFileUrl } from "@/utils/fileFilter";
const emits = defineEmits(["loginSuccess"]);
const userForm = reactive({
  username: "admin",
  password: "12345",
});
const router = useRouter();
const handleLogin = async () => {
  // 首先验证字段是否为空
  if (!userForm.username || !userForm.password) {
    ElMessage.error("用户名或密码不能为空");
    return;
  }
  try {
    // console.log("输出userForm", userForm);
    const res = await login(userForm);
    console.log("输出res", res);
    if (res.status == 200) {
      const avatar = res.data.user.avatar;
      cache.setToken(res.data.token);
      cache.sessionSet("userId", res.data.user.userId);
      cache.sessionSet("username", userForm.username);
      cache.sessionSet("userRole", res.data.user.role);
      cache.sessionSet("userAvatar", getFileUrl(avatar, "image/jpeg"));
      cache.sessionSet("isLogin", true);
      ElMessage.success("登录成功");
      emits("loginSuccess", res.data.user);
    } else {
      ElMessage.error("账号或者密码错误");
    }
  } catch (error) { }
};
const handleRegister = () => {
  router.push("/register");
};
</script>
<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60vh;
  /* background-color: #f5f5f5; */
}

.login-card {
  background-color: #fff;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 350px;
}

h1 {
  margin-bottom: 20px;
  font-size: 2em;
  color: #333;
}

input {
  transition: all 0.3s;
}

input[type="text"],
input[type="password"] {
  width: 80%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 1em;
}

input[type="text"]:focus,
input[type="password"]:focus {
  border-color: var(--el-color-primary);
  outline: none;
  box-shadow: 0 0 8px var(--el-color-primary-light-5);
}

button {
  width: 80%;
  padding: 10px;
  margin-top: 20px;
  border: none;
  border-radius: 5px;
  background-color: var(--el-color-primary);
  color: white;
  font-size: 1em;
  cursor: pointer;
  transition: all 0.3s;
}

button:hover {
  background-color: var(--el-color-success);
}
</style>
