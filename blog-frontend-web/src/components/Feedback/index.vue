<template>
  <div class="feedback">
    <div class="view">
      <svg-icon iconClass="icon-view" className="icon"></svg-icon>
      <span>{{ props.articleInfo.viewCount }}</span>
    </div>

    <div class="like" @click="handleClickLike">
      <template v-if="!likeStatus">
        <svg-icon iconClass="icon-like" className="icon"></svg-icon>
      </template>
      <template v-else>
        <svg-icon iconClass="icon-like-active" className="icon"></svg-icon>
      </template>
      <span>{{ props.articleInfo.likeCount }}</span>
    </div>

    <div class="comment" @click="handleClickComment">
      <svg-icon iconClass="icon-comment" className="icon"></svg-icon>
      {{ props.articleInfo.commentCount }}
    </div>
  </div>

</template>

<script setup>
import { ref } from "vue";
import cache from "@/utils/cache";
import { ElMessage } from "element-plus";
const props = defineProps({
  articleInfo: Object,
});

const likeStatus = ref(false);

// 用户点赞事件
const handleClickLike = () => {
  // 首先要判断用户是否登录，如果没有登录，提示用户登录
  if (!cache.sessionGet("authorization")) {
    ElMessage({
      message: "请先登录",
      type: "warning",
    });
    return;
  }

  if (likeStatus.value) {
    likeStatus.value = false;
    props.articleInfo.likeCount--;
  } else {
    likeStatus.value = true;
    props.articleInfo.likeCount++;
  }

};

// 用户评论事件
const handleClickComment = () => {
  // 首先要判断用户是否登录，如果没有登录，提示用户登录
  if (!cache.sessionGet("authorization")) {
    ElMessage({
      message: "请先登录",
      type: "warning",
    });
    return;
  }
  // 跳转到文章详情页面
  // router.push(`/article/${props.articleInfo.articleId}`);
};
</script>

<style scoped lang="scss">
.feedback {
  display: flex;

  .view,
  .like,
  .comment {
    display: flex;
    align-items: center; // 垂直居中
    margin-left: 20px; // 每个图标之间间距



    span {
      padding-left: 3px;
      font-size: 12px;
      color: #999;
    }
  }
}
</style>