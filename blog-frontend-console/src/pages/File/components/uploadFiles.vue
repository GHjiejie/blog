<template>
  <el-dialog v-model="visible" title="" width="500" :before-close="handleClose">
    <el-form>
      <el-form-item label="文章标签" class="tag">
        <el-select
          v-model="uploadForm.tag"
          placeholder="标签"
          style="width: 240px"
        >
          <el-option-group
            v-for="group in articleTags"
            :key="group.label"
            :label="group.label"
          >
            <el-option
              v-for="item in group.options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-option-group>
        </el-select>
      </el-form-item>
      <input
        type="file"
        style="display: none"
        ref="uploadFileRef"
        @change="filechange"
      />
      <el-form-item label="上传文件" class="file">
        <el-button type="primary" @click="handleUpload">上传文件</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm()"> 上传 </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
<script setup>
import { ref, defineExpose, reactive, defineEmits } from "vue";
import { ElMessage } from "element-plus";
import cache from "@/utils/cache";
import { articleTags } from "@/utils/articles";
import { uploadFile } from "@/apis/file";
const emits = defineEmits("uploadSuccess");
const uploadForm = reactive({
  tag: "",
  uploaderId: cache.sessionGet("userId"),
});

const uploadFileRef = ref(null);
const visible = ref(false);

// 提交表单
const submitForm = async () => {
  const { tag, uploaderId } = uploadForm; // 解构赋值

  if (!tag) {
    ElMessage.error("请选择标签");
    return;
  }

  // 检查是否有文件待上传
  const file = uploadFileRef.value.files[0];
  if (!file) {
    ElMessage.error("请上传文件");
    return;
  }

  // 创建 formData 对象
  const formData = new FormData();
  formData.append("tag", tag);
  formData.append("uploaderId", uploaderId);
  formData.append("file", file);

  try {
    console.log("formData", formData);
    await uploadFile(formData);
    visible.value = false; // 上传成功后关闭对话框或处理其他逻辑
    ElMessage.success("文件上传成功"); // 提示成功信息
    emits("uploadSuccess"); // 触发父组件的 uploadSuccess 事件
  } catch (error) {
    console.error(error); // 输出错误日志，帮助调试
  }
};

// 监听文件变化
// const filechange = () => {
//   console.log(uploadFileRef.value.files[0])
// }

// 点击文件上传按钮
const handleUpload = () => {
  uploadFileRef.value.click();
};

const changeVisible = (val) => {
  visible.value = val;
};
defineExpose({
  changeVisible,
});
</script>

<style scoped lang="scss"></style>
