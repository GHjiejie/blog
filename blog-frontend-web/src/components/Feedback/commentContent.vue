<template>
  <div class="comment-content">
    <div class="comment-content-left">
      <div class="avatar">
        <img src="../../assets/image/cat.jpg" alt="" />
      </div>
    </div>
    <div class="comment-content-right">
      <div class="comment-content-right-header">
        <div class="comment-content-right-header-left">
          <div class="comment-content-right-header-left-username">
            {{ userInfo.username }}
          </div>
          <div class="comment-content-right-header-left-time">
            {{ dayjs(commentInfo.createdAt).format("YYYY-MM-DD HH:mm:ss") }}
          </div>
          <span
            v-if="commentInfo.userId == loginUserId"
            class="deleteBtn"
            @click="deleteComment(commentInfo.id, commentInfo.articleId)"
            >删除评论
          </span>
        </div>
        <!-- 进行回复的评论操作，暂时不做 -->
        <!-- <div class="comment-content-right-header-right">
          <svg-icon iconClass="icon-comment" className="icon" @click="replyComment"></svg-icon>
        </div> -->
      </div>
      <div class="comment-content-right-center">
        <p ref="commentRef">
          {{ commentInfo.content }}
        </p>
        <!-- <textarea name="" id="" v-if="replyState"></textarea> -->
        <div class="replyContainer" v-if="replyState" contenteditable="true">
          <div class="publishCommentReply">
            <!-- <el-button type="primary" @click="publishReply" text
              >回复</el-button
            > -->
          </div>
        </div>

        <!-- 根据溢出状态和展开状态显示按钮 -->
        <el-button
          v-if="isOverflow && !isExpanded"
          type="text"
          @click="showMore"
        >
          展开
        </el-button>
        <el-button
          v-if="isOverflow && isExpanded"
          type="text"
          @click="showLess"
        >
          收起
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick, watch } from "vue";
import { getUserById } from "@/apis/user";
import { getFileUrl } from "@/utils/fileFilter";
import { dayjs } from "element-plus";
import caches from "@/utils/cache";
import { deleteArticleComment } from "@/apis/articles";
import { ElMessageBox, ElMessage } from "element-plus";

const props = defineProps({
  commentInfo: Object,
});

const emits = defineEmits(["deleteComment"]);

const loginUserId = ref("");

loginUserId.value = caches.sessionGet("userId");
console.log("loginUserId", loginUserId.value);
const src = ref("");

const userInfo = ref({});
const commentRef = ref(null); // 绑定的文本元素
const isExpanded = ref(false); // 是否展开状态
const isOverflow = ref(false); // 是否溢出状态

const replyState = ref(false); // 回复状态

const replyComment = () => {
  // console.log("回复评论");
  replyState.value = !replyState.value;
};

const deleteComment = async (commentId, articleId) => {
  //  弹窗提示是否删除
  ElMessageBox.confirm("此操作将永久删除该评论, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      try {
        const res = await deleteArticleComment({
          commentId: commentId,
          articleId: articleId,
        });
        if (res.status === 200) {
          ElMessage.success("删除成功");
          // 删除成功后重新获取评论列表
          emits("deleteComment", commentId);
        }
      } catch (error) {
        console.log(error);
      }
    })
    .catch(() => {
      ElMessage.info("已取消删除");
    });
};

// 展开功能
const showMore = () => {
  commentRef.value.style.webkitLineClamp = "none"; // 展开时不限制行数
  isExpanded.value = true;
};

// 收起功能
const showLess = () => {
  commentRef.value.style.webkitLineClamp = "3"; // 恢复限制为三行
  isExpanded.value = false;
};

// 检测文本是否溢出三行
const checkOverflow = () => {
  if (commentRef.value) {
    const lineHeight = parseFloat(
      getComputedStyle(commentRef.value).lineHeight
    );
    const maxLinesHeight = lineHeight * 3; // 三行的高度
    isOverflow.value = commentRef.value.scrollHeight > maxLinesHeight;
  }
};

// 获取评论的用户的消息
const getUser = async (userId) => {
  try {
    const { data } = await getUserById({ userId });
    userInfo.value = data.user;
    src.value = getFileUrl(userInfo.value.avatar, "image/jpeg");
  } catch (error) {
    console.log(error);
  }
};

// 初始加载时检测溢出状态
onMounted(() => {
  nextTick(() => {
    checkOverflow();
  });
});

// 当 `commentInfo.content` 发生变化时重新检测
watch(
  () => props.commentInfo.content,
  () => {
    nextTick(() => {
      checkOverflow();
    });
  }
);
watch(
  () => props.commentInfo,
  async () => {
    await getUser(props.commentInfo.userId);
  },
  {
    immediate: true,
  }
);
</script>

<style scoped lang="scss">
// 设置滚动条样式

.comment-content {
  display: flex;
  padding: 10px;

  .comment-content-left {
    width: 50px;
    margin-right: 10px;

    .avatar {
      width: 40px;
      height: 40px;
      border-radius: 50%;
      overflow: hidden;

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }

  .comment-content-right {
    flex: 1;

    .comment-content-right-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 5px;

      .comment-content-right-header-left {
        display: flex;
        align-items: center;

        .comment-content-right-header-left-username {
          font-size: 14px;
          font-weight: bold;
          margin-right: 10px;
        }

        .comment-content-right-header-left-time {
          font-size: 12px;
          color: #999;
        }

        .deleteBtn {
          padding-left: 5px;
          font-size: 12px;
          color: #999;
          cursor: pointer;

          &:hover {
            color: red;
          }
        }
      }
    }

    .comment-content-right-center {
      p {
        font-size: 14px;
        overflow: hidden;
        /* 隐藏溢出内容 */
        text-overflow: ellipsis;
        /* 超出内容显示省略号 */
        display: -webkit-box;
        /* 弹性盒子，结合 line-clamp 实现多行省略 */
        -webkit-box-orient: vertical;
        /* 垂直排列 */
        -webkit-line-clamp: 3;
        /* 限制显示的行数为3 */
        line-height: 1.5em;
        /* 每行的高度 */
        transition: all 0.3s;

        /* 新增属性 */
        word-break: break-word;
        /* 自动换行，支持超长单词 */
        white-space: normal;
        /* 文本超过宽度时换行 */
      }

      .replyContainer {
        width: 100%;
        height: 100px;
        border: 1px solid #ccc;
        border-radius: 5px;
        padding: 10px;
        margin-top: 10px;
        overflow-y: auto;
        word-break: break-all;
        white-space: normal;
      }
    }
  }
}
</style>
