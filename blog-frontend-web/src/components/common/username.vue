<template>
  <el-input
    v-bind="$attrs"
    v-model="bindValue"
    placeholder="请输入用户名"
    clearable
  ></el-input>
  <el-tag :type="checkUserNameLength ? 'success' : 'info'">5-20位</el-tag>
  <el-tag :type="checkUserName ? 'success' : 'info'">小写字母或数字</el-tag>
</template>

<script lang="ts" setup>
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
const { checkUserNameLength, checkUserName } = validate(bindValue);
defineExpose({
  checkUserNameLength,
  checkUserName,
});
</script>
