<template>
  <div class="comment-content">
    <div class="comment-content-left">
      <div class="avatar">
        <img :src="userInfo.avatar" alt="" />
      </div>
    </div>
    <div class="comment-content-right">
      <div class="comment-content-right-header">
        <div class="comment-content-right-header-left">
          <div class="comment-content-right-header-left-username">
            {{ userInfo.username }}
          </div>
          <div class="comment-content-right-header-left-time">
            {{ commentInfo.createdAt }}
          </div>
        </div>
      </div>
      <div class="comment-content-right-center">
        <p ref="commentRef">
          {{ commentInfo.content }}
        </p>
        <!-- 根据溢出状态和展开状态显示按钮 -->
        <el-button v-if="isOverflow && !isExpanded" type="text" @click="showMore">
          展开
        </el-button>
        <el-button v-if="isOverflow && isExpanded" type="text" @click="showLess">
          收起
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick, watch } from "vue";
import { getUserById } from "@/apis/user";

const props = defineProps({
  commentInfo: Object
});

const userInfo = ref({});

const commentRef = ref(null); // 绑定的文本元素
const isExpanded = ref(false); // 是否展开状态
const isOverflow = ref(false); // 是否溢出状态

// 展开功能
const showMore = () => {
  commentRef.value.style.webkitLineClamp = "none";
  isExpanded.value = true;
};

// 收起功能
const showLess = () => {
  commentRef.value.style.webkitLineClamp = "3";
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

const getUser = async (userId) => {
  try {
    const { data } = await getUserById({ userId });
    userInfo.value = data.user;
  } catch (error) {
    console.log(error);
  }
}

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
    await getUser(props.commentInfo.userId)
  }, {
  immediate: true
}
);
</script>



<style scoped lang="scss">
.comment-content {
  display: flex;
  // flex-direction: column;
  padding: 10px;
  border-bottom: 1px solid var(--el-color-primary-light-5);

  .comment-content-list {
    display: flex;
    margin-bottom: 10px;
    // flex-direction: column;
  }

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
      }

      .comment-content-right-header-right {
        .comment-content-right-header-right-like {
          font-size: 14px;
          cursor: pointer;
        }
      }
    }

    .comment-content-right-center {
      p {
        font-size: 14px;
        // 设置显示行数
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;
        transition: all 0.3s;
      }
    }
  }
}
</style>