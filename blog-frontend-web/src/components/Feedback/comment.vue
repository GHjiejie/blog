<template>
  <div class="comment-container">
    <div class="comment-header">
      <div class="comment-header-title">评论:{{ total }}</div>
    </div>
    <div class="comment-center">
      <div v-for="(item, index) in commentList" :key="index" class="comment-content-list">
        <CommentContent :commentInfo="item"></CommentContent>
      </div>
    </div>
    <div class="comment-footer">
      <div class="comment-footer-input">
        <input type="text" placeholder="请输入评论内容" />
      </div>
      <div class="comment-footer-button">
        <button>评论</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineExpose, ref, onMounted } from "vue";
import { getArticleComments, publishArticleComment } from "@/apis/articles";
import CommentContent from './commentContent.vue'

const props = defineProps({
  articleId: Number
})
const commentRef = ref(null);
const isExpanded = ref(false);

const page = ref(1);
const pageSize = ref(10);

const commentList = ref([]);
const total = ref(0);



const showMore = () => {
  commentRef.value.style.webkitLineClamp = "none";
  isExpanded.value = true;
};

const showLess = () => {
  commentRef.value.style.webkitLineClamp = "3";
  isExpanded.value = false;
};

onMounted(async () => {
  try {
    const { data } = await getArticleComments({
      articleId: props.articleId,
      page: page.value,
      pageSize: pageSize.value
    });
    commentList.value = data.commentList;
    total.value = data.total;
    console.log("评论列表", commentList.value);

  } catch (error) {

  }

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
    background-color: #f0f0f0;
    border-bottom: 1px solid #ccc;

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
