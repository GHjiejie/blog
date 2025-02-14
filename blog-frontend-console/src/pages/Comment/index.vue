<template>
  <div class="comment-manage">
    <div class="search">
      <el-input v-model="articleId" style="width: 240px" placeholder="请输入文章ID" :prefix-icon="Search"
        @keyup.enter="searchComments" />
    </div>
    <template v-if="searchStatus">
      <div class="commentTable"></div>
      <el-table :data="articleCommentData" style="width: 100%" max-height="100%">

        <el-table-column prop="articleId" label="文章ID" width="120" />
        <el-table-column prop="userId" label="评论者ID" width="120" />
        <el-table-column prop="content" label="评论内容" property="address" width="240" show-overflow-tooltip>

        </el-table-column>
        <el-table-column prop="likeCount" label="点赞数" width="120" />
        <el-table-column prop="createdAt" label="评论发布时间" width="120" />
        <el-table-column fixed="right" label="操作" min-width="120">
          <template #default="scope">
            <el-button link type="primary" size="small" @click.prevent="deleteRow(scope.$index, scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
    <template v-else>
      <el-empty description="" />
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getArticleComments, deleteArticleComment } from "@/apis/articles";

import { Search } from '@element-plus/icons-vue'
import { ElMessageBox, ElMessage } from 'element-plus'

const searchStatus = ref(false);

const page = ref(1);

const pageSize = ref(10);
const articleId = ref("");

const articleCommentData = ref([]);

const deleteRow = async (index, rowData) => {
  console.log(index, rowData);
  // 弹窗提示是否删除
  ElMessageBox.confirm("此操作将永久删除该评论, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      // 删除操作
      const res = await deleteArticleComment({ commentId: rowData.id, articleId: articleId.value });
      // console.log('delete', res);
      if (res.status === 200) {
        articleCommentData.value.splice(index, 1);
      }
      // 删除成功提示
      ElMessage({
        type: "success",
        message: "删除成功!",
      });
    })
    .catch(() => {
      // 删除失败提示
      ElMessage({
        type: "info",
        message: "已取消删除",
      });
    });



};

const searchComments = async () => {
  try {
    const res = await getArticleComments({ articleId: articleId.value, page: page.value, pageSize: pageSize.value });
    articleCommentData.value = res.data.commentList;
    searchStatus.value = true;
    console.log(articleCommentData.value);
  } catch (error) {
    console.log(error);
  }
};


</script>

<style scoped lang="scss">
.comment-manage {
  padding: 20px;
  max-height: calc(100vh - 50px);


  .search {

    margin-bottom: 20px;
  }

  .commentTable {
    margin-bottom: 20px;
  }

  .el-table {
    max-height: 850px;
  }

  .el-table-column {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .el-button {
    margin-right: 10px;
  }
}
</style>
