<template>
  <div class="feedback">
    <div class="left">
      <div class="authorInfo">
        <div class="avatar">
          <el-avatar
            src="https://images.pexels.com/photos/237272/pexels-photo-237272.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
            alt="avatar" />
        </div>
        <div class="authorName">
          <span>{{ props.authorInfo.username }}</span>
        </div>
      </div>
    </div>

    <div class="right">
      <div class="feedbackInfo">
        <FeedBack :articleInfo="props.articleInfo"></FeedBack>
        <!-- <div class="view">
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
</div> -->
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import FeedBack from "@/components/Feedback/index.vue";
// import { ELMessage } from 'element-plus';
const props = defineProps({
  authorInfo: Object,
  articleInfo: Object,
});


</script>

<style scoped lang="scss">
.feedback {
  display: flex;
  align-items: center; // 垂直居中
  padding: 5px 20px;
  background-color: white;
  border-top: 1px solid var(--el-color-primary-light-5); // 底部边框

  .left {
    display: flex;
    align-items: center; // 垂直居中

    .authorInfo {
      display: flex;
      align-items: center; // 垂直居中

      .avatar {
        margin-right: 10px; // 头像右边距

        .el-avatar {
          width: 40px; // 头像宽度
          height: 40px; // 头像高度
          border-radius: 50%; // 圆形头像
        }
      }

      .authorName {
        font-size: 14px;
        color: #333;

        span {
          margin-right: 5px; // 作者名右边距
        }
      }
    }
  }

  .right {
    margin-left: auto; // 右侧内容靠右
    display: flex;


  }
}
</style>
