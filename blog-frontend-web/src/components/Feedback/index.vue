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

    <div class="comment">
      <svg-icon iconClass="icon-comment" className="icon"></svg-icon>
      {{ props.articleInfo.commentCount }}
    </div>
  </div>

</template>

<script setup>
import { ref } from "vue";
const props = defineProps({
  articleInfo: Object,
});

const likeStatus = ref(false);
const handleClickLike = () => {
  // 首先要判断用户是否登录，如果没有登录，提示用户登录

  if (likeStatus.value) {
    likeStatus.value = false;
    props.articleInfo.likeCount--;
  } else {
    likeStatus.value = true;
    props.articleInfo.likeCount++;
  }

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