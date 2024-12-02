<template>
  <el-input
    v-bind="$attrs"
    v-model="bindValue"
    placeholder="请输入8-20位密码，包含大小写字母数字和特殊字符"
    type="password"
    clearable
    show-password
  ></el-input>
  <el-tag :type="checkLength ? 'success' : 'info'">8-20位</el-tag>
  <el-tag :type="checkUpperCase ? 'success' : 'info'">大写字母</el-tag>
  <el-tag :type="checkLowerCase ? 'success' : 'info'">小写字母</el-tag>
  <el-tag :type="checkNumber ? 'success' : 'info'">数字</el-tag>
  <el-tag :type="checkSpecialCharacter ? 'success' : 'info'">特殊字符</el-tag>
</template>

<script setup>
import { computed } from "vue";
import validate from "@/composables/common/validate";
const props = defineProps({
  value: {
    type: String,
    default: "",
  },
});
const emits = defineEmits(["update:value"]);
// 通过重写计算属性的set和get方法，将计算属性的结果绑定在输入框的model中
const bindValue = computed({
  get: () => props.value,
  set: (val) => {
    emits("update:value", val);
  },
});
const {
  checkLength,
  checkUpperCase,
  checkLowerCase,
  checkNumber,
  checkSpecialCharacter,
} = validate(bindValue);
defineExpose({
  checkLength,
  checkUpperCase,
  checkLowerCase,
  checkNumber,
  checkSpecialCharacter,
});
</script>
