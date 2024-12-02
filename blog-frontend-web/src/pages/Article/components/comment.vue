<template>
  <div class="feedback">
    <div class="left">
      <div class="authorInfo">
        <div class="avatar">
          <el-avatar
            src="https://images.pexels.com/photos/237272/pexels-photo-237272.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
            alt="avatar"
          />
        </div>
        <div class="authorName">
          <span>{{ props.authorInfo.username }}</span>
        </div>
      </div>
    </div>

    <div class="right">
      <div class="feedbackInfo">
        <div class="view">
          <el-icon>
            <View />
          </el-icon>
          <span>{{ props.articleInfo.viewCount }}</span>
        </div>

        <div class="like">
          <el-icon>
            <Star />
          </el-icon>
          <!-- <span>{{ item.likeCount }}</span> -->
          <span>{{ props.articleInfo.likeCount }}</span>
        </div>

        <div class="comment">
          <el-icon>
            <ChatLineRound />
          </el-icon>

          {{ props.articleInfo.commentCount }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Star, ChatLineRound, View } from "@element-plus/icons-vue";

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

    .feedbackInfo {
      display: flex;

      .view,
      .like,
      .comment {
        display: flex;
        align-items: center; // 垂直居中
        margin-left: 20px; // 每个图标之间间距

        .el-icon {
          margin-right: 5px; // 图标右边距
          font-size: 16px; // 图标大小
          color: #999; // 图标颜色
        }

        span {
          font-size: 12px;
          color: #999;
        }
      }
    }
  }
}
</style>
