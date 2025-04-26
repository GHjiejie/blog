<template>
  <el-dialog v-model="visible" title="" width="500" :before-close="handleClose">
  </el-dialog>
</template>
<script setup>
import { ref, defineExpose, reactive, defineEmits } from "vue";
import { ElMessage } from "element-plus";
import cache from "@/utils/cache";

const emits = defineEmits("uploadSuccess");

const visible = ref(false);

const changeVisible = (val) => {
  visible.value = val;
};
defineExpose({
  changeVisible,
});
</script>

<style scoped lang="scss"></style>
