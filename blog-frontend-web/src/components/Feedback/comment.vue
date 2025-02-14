<template>
  <div class="comment-container">
    <div class="comment-header">
      <div class="comment-header-title">评论:{{ total }}</div>
    </div>
    <div class="comment-center">
      <div v-for="(item, index) in commentList" :key="index" class="comment-content-list">
        <CommentContent :commentInfo="item"></CommentContent>
      </div>
      <div class="showMore">
        <el-button v-if="hasMore" type="text" @click="showMore">查看更多</el-button>
        <el-button v-else type="text" disabled>没有更多了</el-button>
      </div>
    </div>
    <div class="comment-footer">
      <div class="comment-footer-input">
        <el-input v-model="inputText" placeholder="请输入评论内容" />
      </div>
      <div class="comment-footer-button">
        <el-button type="primary" @click="publishComment">发表</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineExpose, ref, onMounted } from "vue";
import { getArticleComments, publishArticleComment } from "@/apis/articles";
import CommentContent from "./commentContent.vue";
import caches from "@/utils/cache";

const props = defineProps({
  articleId: Number,
});
const hasMore = ref(true);

const page = ref(1);
const pageSize = ref(10);

const commentList = ref([]);
const total = ref(0);

const inputText = ref("");

const publishComment = async () => {
  try {
    const res = await publishArticleComment({
      articleId: props.articleId,
      content: inputText.value,
      userId: caches.sessionGet("userId"),
    });
    if (res.status === 200) {
      inputText.value = "";
      page.value = 1;
      const { data } = await getArticleComments({
        articleId: props.articleId,
        page: page.value,
        pageSize: pageSize.value,
      });
      commentList.value = data.commentList;
      total.value = data.total;
    }
  } catch (error) { }
};
const showMore = async () => {
  try {
    page.value++;
    const { data } = await getArticleComments({
      articleId: props.articleId,
      page: page.value,
      pageSize: pageSize.value,
    });
    if (data.commentList.length === 0) {
      hasMore.value = false;
      return;
    }
    commentList.value = [...commentList.value, ...data.commentList];
    total.value = data.total;
  } catch (error) { }
};

onMounted(async () => {
  try {
    const { data } = await getArticleComments({
      articleId: props.articleId,
      page: page.value,
      pageSize: pageSize.value,
    });
    if (data.commentList.length === 0) {
      hasMore.value = false;
    } else {
      commentList.value = data.commentList;
      total.value = data.total;
    }
  } catch (error) { }
});
</script>

<style scoped lang="scss">
.comment-container {
  width: 100%;
  height: 100vh;
  overflow-y: auto;
  margin: 0 auto;
  border: 1px solid #ccc;
  border-radius: 5px;
  overflow: hidden;
  position: relative;

  .comment-header {
    padding: 10px;
    height: 50px;
    border-bottom: 1px solid var(--el-color-primary-light-5);

    .comment-header-title {
      font-size: 16px;
      font-weight: bold;
    }
  }

  .comment-center {
    padding: 10px;
    box-sizing: border-box;
    overflow-y: auto;
    height: calc(100vh - 100px);
  }

  .showMore {
    text-align: center;
    margin-top: 10px;
  }

  .comment-footer {
    height: 50px;

    display: flex;
    padding: 10px;
    box-sizing: border-box;

    .comment-footer-input {
      flex: 1;
      margin-right: 10px;

      input {
        width: 100%;
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 5px;
      }
    }

    .comment-footer-button {
      button {
        padding: 8px 15px;
        background-color: #4caf50;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
      }
    }
  }
}
</style>
