<template>
  <el-dialog v-model="visible" :title="fileName" width="70%" top="5vh" center :show-close="false">
    <div class="viewContainer" v-loading="loading" element-loading-svg-view-box="-10, -10, 50, 50"
      :element-loading-spinner="svg" element-loading-background="rgba(122, 122, 122, 0.8)">
      <div class="imgContainer" v-if="
        fileType === 'jpg' ||
        fileType === 'png' ||
        fileType === 'jpeg' ||
        fileType === 'gif' ||
        fileType === 'bmp' ||
        fileType === 'webp'
      ">
        <img :src="src"></img>
      </div>
      <v-md-preview v-else-if="fileType === 'md'" :text="fileContent"></v-md-preview>
      <div v-else-if="fileType === 'pdf' || fileType === 'txt'" class="iframeContainer">
        <iframe :src="src"></iframe>
      </div>
      <JsonViewer v-else-if="fileType === 'json'" sort copyable boxed :value="fileContent" :expand-depth="999">
      </JsonViewer>
    </div>

  </el-dialog>
</template>
<script setup>
import { ref, defineExpose, watch } from "vue";
import { getFileById } from "@/apis/file";
import { decodeBase64, getFileType, getFileUrl } from "@/utils/fileFilter";

const props = defineProps({
  fileId: String,
});
const fileContent = ref(null);
const fileType = ref("");
const visible = ref(false);
const fileName = ref("");

let src = ref("");
const loading = ref(false);
const svg = `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `

const changeVisible = (val) => {
  visible.value = val;
};
const getFileContent = async () => {
  try {
    loading.value = true;
    const { data } = await getFileById({ fileId: props.fileId });
    fileName.value = data.fileInfo.fileName;
    // fileContent.value = decodeBase64(data.fileInfo.content);
    fileType.value = getFileType(data.fileInfo.fileType);
    if (
      fileType.value === "jpg" ||
      fileType.value === "png" ||
      fileType.value === "jpeg" ||
      fileType.value === "gif" ||
      fileType.value === "bmp" ||
      fileType.value === "webp" ||
      fileType.value === "pdf" ||
      fileType.value === "txt"
    ) {
      src.value = getFileUrl(data.fileInfo.content, data.fileInfo.fileType);
      console.log("输出返回的文件地址：", src.value);
    } else if (
      fileType.value === "json" ||
      fileType.value === "md"
    ) {
      fileContent.value = decodeBase64(data.fileInfo.content);
      console.log("输出返回的文件内容：", fileContent.value);
    }
    loading.value = false;
  } catch (error) { }
};

// 监听props.fileId的变化
watch(
  () => props.fileId,
  async (newVal) => {
    try {
      await getFileContent();
    } catch (error) { }
  }
);
defineExpose({
  changeVisible,
});
</script>
<style scoped lang="scss">
iframe {
  border: none;
  width: 100%;
  height: 80vh;
}

.imgContainer {
  display: flex;
  justify-content: center;
  align-items: center;

  img {
    width: 100%;
    max-height: 80vh;
    object-fit: cover
  }
}
</style>
