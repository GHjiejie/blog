<template>
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
          prop="articleId"
          label="文章ID"
          width="120"
        />
        <el-table-column
          property="title"
          label="文章标题"
          width="200"
          show-overflow-tooltip
        />
        <!-- <el-table-column prop="title" label="文章标题" width="200" /> -->
        <el-table-column prop="tag" label="文件标签" width="150">
          <template #default="{ row }">
            <el-tag size="small">{{ row.tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="authorId" label="文章作者" width="120">
        </el-table-column>
        <el-table-column prop="status" label="文章状态" width="120">
          <template #default="{ row }">
            <span :class="getStatusClass(row.status)" id="statusText">{{
              getArticleStatus(row.status)
            }}</span>
          </template>
        </el-table-column>

        <el-table-column prop="updatedAt" label="更新时间" width="200">
          <template #default="{ row }">
            {{ dayjs(row.createdAt).format("YYYY-MM-DD HH:mm:ss") }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" width="200">
          <template #header>
            <el-input
              v-model="search"
              size="small"
              placeholder="输入文章作者ID"
            />
          </template>
          <template #default="{ row }">
            <el-button
              :disabled="row.role === 'ADMIN'"
              link
              type="success"
              size="small"
              @click="reviewArticle(row.articleId)"
            >
              审核
            </el-button>
            <el-button
              :disabled="row.role === 'ADMIN'"
              link
              type="danger"
              size="small"
              @click="handelDelArticle(row)"
            >
              删除
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
  <articleReview
    ref="articleReviewRef"
    :article-id="articleReviewId"
    @changeArticleStatus="updateArticleStatus"
  >
  </articleReview>
</template>
<script setup>
import { ref, onMounted, computed } from "vue";
import { getArticleList, deleteArticle } from "@/apis/articles";
import { ElMessageBox, ElMessage, dayjs } from "element-plus";
import articleReview from "./articleReview.vue";
import { getArticleStatus } from "@/utils/articles";
const page = ref(1);
const pageSize = ref(10);
const articleList = ref([]);
const search = ref("");
const maxHeight = window.innerHeight - 100;

const articleReviewId = ref("");
const articleReviewRef = ref(null);

onMounted(async () => {
  await getArticleListData();
});

const getStatusClass = (status) => {
  return {
    published: status === "PUBLISHED",
    deleted: status === "DELETED",
    "review-passed": status === "REVIEW_PASSED",
    "review-failed": status === "REVIEW_FAILED",
    reviewing: status === "REVIEWING",
    draft: status === "DRAFT",
    default: ![
      "PUBLISHED",
      "DELETED",
      "REVIEW_PASSED",
      "REVIEW_FAILED",
      "REVIEWING",
      "DRAFT",
    ].includes(status),
  };
};

const filterTableData = computed(() => {
  return articleList.value.filter((item) => {
    return item.authorId.includes(search.value);
  });
});

// 审核文章
const reviewArticle = (articleId) => {
  articleReviewId.value = articleId;
  articleReviewRef.value.changeVisible(true);
};

// 更新文章状态
const updateArticleStatus = (data) => {
  const { articleId, status } = data;
  const index = articleList.value.findIndex(
    (item) => item.articleId === articleId
  );
  articleList.value[index].status = status;
};

// 获取文章列表
const getArticleListData = async () => {
  try {
    const { data } = await getArticleList({
      page: page.value,
      pageSize: pageSize.value,
    });
    if (data.articleList.length === 0) {
      ElMessage.info("没有更多数据了");
      return;
    }
    articleList.value = [...articleList.value, ...data.articleList];
  } catch (error) {
    console.log(error);
  }
};

// 删除文章
const handelDelArticle = (row) => {
  // console.log(row, articleId)
  ElMessageBox.confirm("此操作将永久删除该文章, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      const res = await deleteArticle({ articleId: row.articleId });
      if (res.status === 200) {
        ElMessage.success("删除成功");
        // 过滤掉删除的文章
        articleList.value = articleList.value.filter(
          (item) => item.articleId !== row.articleId
        );
      }
    })
    .catch(() => {});
};

// 加载更多
const handleNextPage = async () => {
  page.value++;
  await getArticleListData();
};
</script>
<style scoped lang="scss">
#statusText {
  cursor: pointer;
  font-weight: bold;
  font-size: 12px;
}

.published {
  color: $--color-status-published;
}

.deleted {
  color: $--color-status-deleted;
}

.review-passed {
  color: $--bg-color-status-review-passed;
}

.review-failed {
  color: $--bg-color-status-review-failed;
}

.reviewing {
  color: $--bg-color-status-reviewing;
}

.draft {
  color: gray;
}

.default {
  color: black;
}

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
