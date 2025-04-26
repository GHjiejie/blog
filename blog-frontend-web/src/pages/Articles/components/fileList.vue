<template>
  <el-dialog
    v-model="visible"
    :title="fileName"
    width="70%"
    top="5vh"
    center
    :show-close="false"
  >
    <div class="container">
      <div class="Top"></div>
      <div class="Center">
        <el-table
          :data="filterTableData"
          size="large"
          :height="maxHeight"
          :highlight-current-row="false"
        >
          <el-table-column
            fixed="left"
            prop="fileName"
            label="文件名"
            width="200"
          />

          <el-table-column prop="tag" label="标签" width="120">
            <template #default="{ row }">
              <el-tag size="small" round>{{ row.tag }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="fileType" label="文件类型" width="200" />
          <el-table-column fixed="right" width="200">
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
                @click="selectFile(row.fileId)"
              >
                选择
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
  </el-dialog>
</template>
<script setup>
import { ref, defineExpose, computed, onMounted } from "vue";
import { getFileList } from "@/apis/file";
import { ElMessage } from "element-plus";

const emits = defineEmits(["selectOnlineFile"]);

const visible = ref(false);
const search = ref("");

const page = ref(1);
const pageSize = ref(10);
const fileList = ref([]);

const maxHeight = window.innerHeight - 200;
const changeVisible = (val) => {
  visible.value = val;
};
onMounted(async () => {
  await getListFiles();
});
// 表单数据过滤
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
  console.log("获取文件列表");

  try {
    const { data } = await getFileList({
      page: page.value,
      pageSize: pageSize.value,
    });
    // fileList.value = [...fileList.value, ...data.fileInfos];
    if (data.fileInfos.length === 0) {
      // ElMessage({
      //   message: "没有更多数据了",
      //   type: "info",
      // });
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
// 选择文件
const selectFile = (fileId) => {
  emits("selectOnlineFile", fileId);
  changeVisible(false);
};
defineExpose({
  changeVisible,
});
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
