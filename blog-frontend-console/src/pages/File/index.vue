<template>
  <div class="container">
    <div class="Top">
      <button class="addserBtn" @click="openAddUserDialog">
        <el-icon><Plus /></el-icon>
      </button>
    </div>
    <div class="Center">
      <el-table
        :data="filterTableData"
        size="large"
        :height="maxHeight"
        :highlight-current-row="true"
      >
        <el-table-column
          fixed="left"
          prop="fileId"
          label="文件ID"
          width="120"
        />
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

        <el-table-column prop="createdAt" label="创建时间" width="300" />
        <el-table-column fixed="right" min-width="150">
          <template #header>
            <el-input
              v-model="search"
              size="small"
              placeholder="Type to search"
            />
          </template>
          <template #default="{ row }">
            <el-button
              :disabled="row.role === 'ADMIN'"
              link
              type="success"
              size="small"
              @click="viewFile(row, row.userId)"
            >
              查看
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="Footer">
      <el-button type="primary" @click="handleNextPage()" size="small"
        >加载更多</el-button
      >
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted, computed } from "vue";
import { getFileList } from "@/apis/file";
import { ElMessage, ElMessageBox } from "element-plus";
import { Plus } from "@element-plus/icons-vue";
import { getFileSize } from "@/utils/calculation";

let page = ref(1);
let pageSize = ref(15);
const search = ref("");
const fileList = ref([]);
const maxHeight = window.innerHeight - 100;

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
  } catch (error) {}
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
