<template>
  <div class="feedback">
    <div class="left">
      <div class="authorInfo">
        <div class="avatar">
          <el-avatar :src="src" alt="avatar" />
        </div>
        <div class="authorName">
          <span>{{ userInfo.username }}</span>
        </div>
      </div>
    </div>

    <div class="right">
      <div class="feedbackInfo">
        <FeedBack :articleInfo="props.articleInfo" :active-position="position" :is-like="isLike"></FeedBack>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import FeedBack from "@/components/Feedback/index.vue";
import cache from "@/utils/cache";
import { getArticleLikeStatus } from "@/apis/articles";
import { getFileUrl } from "@/utils/fileFilter";
import { useRoute } from "vue-router";
import { getUserById } from "@/apis/user.js";
const route = useRoute();
// import { ELMessage } from 'element-plus';
const props = defineProps({
  articleInfo: Object,
});
const src = ref('')
const isLike = ref(false);
const userInfo = ref({});

const userId = cache.sessionGet("userId");

onMounted(async () => {
  await getUser();
})
const getUser = async () => {
  const res = await getUserById({ userId });
  userInfo.value = res.data.user;
  src.value = getFileUrl(userInfo.value.avatar, 'image/jpeg')
}


watch(
  () => route.params.id,
  async (newVal) => {
    if (userId) {
      const res = await getArticleLikeStatus({
        articleId: newVal.articleId,
        userId: userId,
      });
      if (res.status == 200) {
        console.log("判断用户是否点赞该文章", res);
        isLike.value = true;
      }
    }
  }
);
// onMounted(async () => {
// 这个里面访问不到props，所以要用watch监听props的变化？？？但是这是为什么？？？？
//   console.log("props", props);
//   console.log("props.articleInfo", props.articleInfo);

//   if (userId) {
//     const res = await getArticleLikeStatus({
//       articleId: props.articleInfo.articleId,
//       userId: userId,
//     });
//     console.log("输出res", res);
//   }
// });

const position = ref("articleDetail");
</script>

<style scoped lang="scss">
.feedback {
  display: flex;
  align-items: center; // 垂直居中
  padding: 5px 20px;
  background-color: white;
  border-top: 1px solid var(--el-color-primary-light-5); // 上边框

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
