<template>
  <el-dialog
    v-model="dialogVisible"
    :title="title"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    width="550"
  >
    <el-form
      ref="formDataRef"
      class="user-form"
      label-position="right"
      label-width="80px"
      :model="formData"
      :rules="rules"
    >
      <el-form-item label="角色">
        <el-select v-model="formData.role" disabled class="role-select">
          <el-option
            v-for="item in roleList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="用户名" prop="username">
        <username
          ref="usernameRef"
          v-model:value="formData.username"
        ></username>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <password
          ref="passwordRef"
          v-model:value="formData.password"
        ></password>
      </el-form-item>
      <el-form-item label="确认密码" prop="checkPassword">
        <el-input
          v-model="formData.checkPassword"
          placeholder="请输入确认密码"
          type="password"
          clearable
          show-password
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="changeVisible(false)">取消</el-button>
        <el-button type="primary" @click="submitData()">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup>
import { computed, ref, reactive } from "vue";
import password from "@/components/common/password.vue";
import username from "@/components/common/username.vue";
import { ElMessage } from "element-plus";

const props = defineProps({
  userId: String,
});

// const props = defineProps<{
//   userId: string
// }>()
const emits = defineEmits(["success"]);
const defaultFormData = {
  username: "",
  role: "USER",
  password: "",
  checkPassword: "",
};
const roleList = [
  {
    label: "管理员",
    value: "ADMIN",
  },
  {
    label: "用户",
    value: "USER",
  },
];
const passwordRef = ref(null);
const usernameRef = ref(null);
const formDataRef = ref();
let formData = ref({
  ...defaultFormData,
});
let validatePassword = async (rule, value) => {
  if (value === "") {
    return Promise.reject("请输入密码");
  } else {
    if (formData.value.checkPassword !== "") {
      formDataRef.value.validateField("checkPassword");
    }
    return Promise.resolve();
  }
};
let checkPass = async (rule, value) => {
  if (value === "") {
    return Promise.reject("请再次输入密码");
  } else if (value !== formData.value.password) {
    return Promise.reject("两次输入的密码不一致！");
  } else {
    return Promise.resolve();
  }
};
const rules = reactive({
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [
    { required: true, message: "密码不能为空", trigger: "blur" },
    { validator: validatePassword, trigger: "change" },
  ],
  checkPassword: [
    { required: true, message: "密码不能为空", trigger: "blur" },
    { validator: checkPass, trigger: "change" },
  ],
});
const pwdDisabled = computed(() => {
  if (passwordRef.value) {
    const {
      checkLength,
      checkUpperCase,
      checkLowerCase,
      checkNumber,
      checkSpecialCharacter,
    } = passwordRef.value;
    return !(
      checkLength &&
      checkUpperCase &&
      checkLowerCase &&
      checkNumber &&
      checkSpecialCharacter
    );
  } else {
    return true;
  }
});
const nameDisabled = computed(() => {
  if (usernameRef.value) {
    const { checkUserNameLength, checkUserName } = usernameRef.value;
    return !(checkUserNameLength && checkUserName);
  } else {
    return true;
  }
});
const dialogVisible = ref(false);
const title = computed(() => {
  return "新建用户";
});
const changeVisible = (val) => {
  dialogVisible.value = val;
  if (!val) {
    formData.value = {
      ...defaultFormData,
    };
    formDataRef.value?.resetFields();
  }
};
const submitData = () => {
  if (!formDataRef.value) return;
  formDataRef.value.validate((valid) => {
    if (valid && !pwdDisabled.value && !nameDisabled.value) {
      // let params: any = {
      //   ...formData.value
      // }
      // if (props.userId) {
      //   params.user_id = props.userId
      // }
      emits("success", formData.value);
    } else {
      ElMessage.error("请检查输入格式是否正确");
      return false;
    }
  });
};
defineExpose({
  changeVisible,
});
</script>
<style lang="scss" scoped>
.user-form {
  padding: 30px 40px;
}
.role-select {
  width: 100%;
}
</style>
