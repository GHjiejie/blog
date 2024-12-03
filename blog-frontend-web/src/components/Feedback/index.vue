<template>
  <div class="feedback">
    <div class="view">
      <svg-icon iconClass="icon-view" className="icon"></svg-icon>
      <span>{{ props.articleInfo.viewCount }}</span>
    </div>

    <div
      v-if="props.activePosition !== 'articleList'"
      class="like"
      @click="handleClickLike"
    >
      <template v-if="likeStatus">
        <svg-icon iconClass="icon-like-active" className="icon"></svg-icon>
      </template>
      <template v-else>
        <svg-icon iconClass="icon-like" className="icon"></svg-icon>
      </template>
      <span>{{ props.articleInfo.likeCount }}</span>
    </div>

    <div
      v-if="props.activePosition !== 'articleList'"
      class="comment"
      @click="handleClickComment"
    >
      <svg-icon iconClass="icon-comment" className="icon"></svg-icon>
      <span> {{ props.articleInfo.commentCount }}</span>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
import cache from "@/utils/cache";
import { ElMessage } from "element-plus";
import { cancelArticleLikeCount, addArticleLikeCount } from "@/apis/articles";
const props = defineProps({
  articleInfo: Object,
  activePosition: String,
  isLike: Boolean,
});

let likeStatus = props.isLike;

const userId = cache.sessionGet("userId");
// 用户点赞事件
const handleClickLike = async () => {
  // 首先要判断用户是否登录，如果没有登录，提示用户登录
  if (!cache.sessionGet("authorization")) {
    ElMessage({
      message: "请先登录哦",
      type: "warning",
    });
    return;
  }

  if (likeStatus) {
    const res = await cancelArticleLikeCount({
      articleId: props.articleInfo.articleId,
      userId: userId,
    });
    if (res.status == 200) {
      console.log(res);
      likeStatus = false;
      props.articleInfo.likeCount--;
    } else {
      ElMessage.error("取消点赞失败");
    }
  } else {
    const res = await addArticleLikeCount({
      articleId: props.articleInfo.articleId,
      userId: userId,
    });
    if (res.status == 200) {
      console.log(res);
      likeStatus = true;
      props.articleInfo.likeCount++;
    } else {
      ElMessage.error("点赞失败");
    }
  }
};

// 用户评论事件
const handleClickComment = () => {
  // 首先要判断用户是否登录，如果没有登录，提示用户登录
  if (!cache.sessionGet("authorization")) {
    ElMessage({
      message: "请先登录哦",
      type: "warning",
    });
    return;
  }
  // 跳转到文章详情页面
  // router.push(`/article/${props.articleInfo.articleId}`);
};

watch(
  () => props.isLike,
  (newVal) => {
    likeStatus = newVal;
  },
  {
    immediate: true,
  }
);
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
