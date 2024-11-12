<template>
  <div class="container">
    <div class="Top">
      <button class="addserBtn" @click="openAddUserDialog">
        <el-icon>
          <Plus />
        </el-icon>
      </button>
    </div>
    <div class="Center">
      <el-table :data="filterTableData" size="large" :height="maxHeight" :highlight-current-row="false">
        <el-table-column fixed="left" prop="fileId" label="文件ID" width="120" />
        <el-table-column prop="fileName" label="文件名" width="200" />
        <el-table-column prop="bytes" label="文件大小" width="150">
          <template #default="{ row }">
            {{ getFileSize(row.bytes) }}
          </template>
        </el-table-column>
        <el-table-column prop="tag" label="标签" width="120">
          <template #default="{ row }">
            <el-tag size="small" round>{{ row.tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="fileType" label="文件类型" width="200" />

        <el-table-column prop="createdAt" label="创建时间" width="200" />
        <el-table-column fixed="right" width="200">
          <template #header>
            <el-input v-model="search" size="small" placeholder="Type to search" />
          </template>
          <template #default="{ row }">
            <el-button :disabled="row.role === 'ADMIN'" link type="success" size="small"
              @click="viewFile(row, row.fileId)">
              查看
            </el-button>
            <el-button :disabled="row.role === 'ADMIN'" link type="danger" size="small"
              @click="handelDelFile(row, row.fileId)">
              删除
            </el-button>
            <el-button :disabled="row.role === 'ADMIN'" link type="info" size="small" @click="handelDownload(row)">
              下载
            </el-button>

          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="Footer">
      <el-button type="primary" @click="handleNextPage()" size="small">加载更多</el-button>
    </div>
  </div>
  <fileView ref="fileViewRef" :file-id="viewFileId"></fileView>
</template>
<script setup>
import { ref, onMounted, computed } from "vue";
import {
  getFileList,
  deleteFile,
  getFileById,
} from "@/apis/file";
import { ElMessage, ElMessageBox } from "element-plus";
import { Plus } from "@element-plus/icons-vue";
import { getFileSize } from "@/utils/calculation";
import fileView from "./components/fileView.vue";
import { downloadFileContent } from "@/utils/fileFilter";

let page = ref(1);
let pageSize = ref(15);
const search = ref("");
const fileList = ref([]);
const maxHeight = window.innerHeight - 100;

const fileViewRef = ref(null);
const viewFileId = ref(null);
onMounted(async () => {
  await getListFiles();
});

const filterTableData = computed(() => {
  return fileList.value.filter((item) => {
    return (
      item.fileName.includes(search.value) ||
      item.tag.includes(search.value) ||
      item.fileType.includes(search.value)
    );
  });
});

// 下载文件
const handelDownload = async (row) => {
  try {
    const { data } = await getFileById({ fileId: row.fileId });
    downloadFileContent(row.fileType, row.fileName, data.fileInfo.content);
  } catch (error) {
  }
};

// 预览文件
const viewFile = async (row, fileId) => {
  viewFileId.value = fileId;
  fileViewRef.value.changeVisible(true);
};

// 删除文件
const handelDelFile = async (row, fileId) => {
  ElMessageBox.confirm("确定删除该文件吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      try {
        await deleteFile({ fileId });
        ElMessage({
          message: "删除成功",
          type: "success",
        });
        fileList.value = fileList.value.filter(
          (item) => item.fileId !== fileId
        );
      } catch (error) {
        ElMessage({
          message: "删除失败",
          type: "error",
        });
      }
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消删除",
      });
    });
};

// 获取文件列表
const getListFiles = async () => {
  try {
    const { data } = await getFileList({
      page: page.value,
      pageSize: pageSize.value,
    });
    // fileList.value = [...fileList.value, ...data.fileInfos];
    if (data.fileInfos.length === 0) {
      ElMessage({
        message: "没有更多数据了",
        type: "info",
      });
      return;
    } else {
      fileList.value = [...fileList.value, ...data.fileInfos];
    }
  } catch (error) { }
};

// 加载更多
const handleNextPage = async () => {
  page.value++;
  await getListFiles();
};
</script>
<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  height: 100%;

  .Top {
    position: fixed;
    top: 50%;
    right: 2%;
    z-index: 999;

    .addserBtn {
      background-color: $--color-primary-opacity7;
      color: $--color-text;
      border: none;
      padding: 10px;
      border-radius: 54% 46% 21% 79% / 52% 46% 54% 48%;

      box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
      transition: all 0.3s;
      cursor: pointer;

      &:hover {
        background-color: $--color-primary;
      }
    }
  }

  .Footer {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 50px;
    transition: all 0.3s;

    .el-button {
      &:hover {
        color: $--color-text3;
        background-color: $--color-primary;
      }
    }
  }
}
</style>
