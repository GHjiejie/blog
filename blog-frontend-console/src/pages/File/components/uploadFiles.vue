<template>
  <el-dialog
    v-model="visible"
    title="上传文件"
    width="500"
    :before-close="handleClose"
    class="upload-dialog"
  >
    <el-form label-width="80px">
      <el-form-item label="选择文件" class="file">
        <el-button type="primary" @click="handleUpload" :icon="Upload">
          选择文件
        </el-button>
        <span v-if="selectedFileName" class="file-name">{{
          selectedFileName
        }}</span>
      </el-form-item>
      <el-form-item>
        <el-button
          type="success"
          @click="submitForm"
          :icon="UploadFilled"
          :loading="uploading"
          class="submit-button"
        >
          {{ uploading ? "上传中..." : "上传" }}
        </el-button>
      </el-form-item>
    </el-form>

    <input
      type="file"
      style="display: none"
      ref="uploadFileRef"
      @change="filechange"
    />
  </el-dialog>
</template>

<script setup>
import { ref, defineExpose, reactive, defineEmits } from "vue";
import { ElMessage } from "element-plus";
import cache from "@/utils/cache";
//import { articleTags } from '@/utils/articles';  //  不再使用
import { uploadFile } from "@/apis/file";
import { Upload, UploadFilled } from "@element-plus/icons-vue"; // 引入 Element Plus 图标

const emits = defineEmits(["uploadSuccess"]); // 将数组改为数组表达式
const uploadForm = reactive({
  uploaderId: cache.sessionGet("userId"),
});

const uploadFileRef = ref(null);
const visible = ref(false);
const selectedFileName = ref(""); // 用于显示选中的文件名
const uploading = ref(false); // 用于表示正在上传状态

// 提交表单
const submitForm = async () => {
  const { uploaderId } = uploadForm;
  const file = uploadFileRef.value.files[0];

  if (!file) {
    ElMessage.error("请选择文件");
    return;
  }

  const formData = new FormData();
  formData.append("uploaderId", uploaderId);
  formData.append("file", file);

  uploading.value = true; // 设置上传状态为 true
  try {
    await uploadFile(formData);
    visible.value = false;
    ElMessage.success("文件上传成功");
    emits("uploadSuccess");
  } catch (error) {
    console.error(error);
    ElMessage.error("文件上传失败"); // 添加上传失败的提示
  } finally {
    uploading.value = false; // 无论成功或失败，都将上传状态设置为 false
    selectedFileName.value = ""; // 清空文件名，允许再次上传相同文件
    uploadFileRef.value.value = ""; // 清空 input file 的值，使文件change事件能被触发
  }
};

// 监听文件变化
const filechange = () => {
  selectedFileName.value = uploadFileRef.value.files[0]?.name || ""; // 显示文件名
};

// 点击文件上传按钮
const handleUpload = () => {
  uploadFileRef.value.click();
};

const changeVisible = (val) => {
  visible.value = val;
};

const handleClose = () => {
  // 关闭 Dialog 时，清空 input file的值和已选择的文件名
  uploadFileRef.value.value = "";
  selectedFileName.value = "";
  visible.value = false;
};

defineExpose({
  changeVisible,
});
</script>

<style scoped lang="scss">
.upload-dialog {
  .el-dialog__header {
    text-align: center;
    font-size: 1.4em;
    font-weight: bold;
    padding: 15px;
  }

  .el-form-item {
    margin-bottom: 22px;

    &.file {
      display: flex;
      align-items: center;

      .el-button {
        margin-right: 15px;
      }

      .file-name {
        font-style: italic;
        color: #999;
      }
    }
  }

  .submit-button {
    width: 100%;
    padding: 12px 20px;
    font-size: 1.1em;
    letter-spacing: 1px;
  }
}
</style>
