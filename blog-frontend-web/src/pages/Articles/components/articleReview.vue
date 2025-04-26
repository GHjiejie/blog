<template>
  <el-dialog
    v-model="visible"
    title=""
    width="60%"
    :before-close="handleClose"
    :show-close="false"
  >
    <div class="articleContainer">
      <div class="Top">
        <div class="articleTitle">{{ articleInfo.title }}</div>
        <div class="articleBaseInfo">
          <div class="author">
            <el-icon>
              <User />
            </el-icon>
            <span>{{ userInfo.username }}</span>
          </div>
          <span class="createAt">{{ articleInfo.updatedAt }}</span>
          <div class="viewCount">
            <el-icon>
              <View />
            </el-icon>
            <span>{{ articleInfo.viewCount }}</span>
          </div>
          <div class="tag">
            <el-tag size="small" round>{{ articleInfo.tag }}</el-tag>
          </div>
        </div>
      </div>
      <div class="Center">
        <v-md-preview :text="articleInfo.content"></v-md-preview>
      </div>
      <div class="Footer">
        <el-tag
          v-if="
            articleInfo.status === 'PUBLISHED' ||
            articleInfo.status === 'REVIEW_PASSED'
          "
          type="danger"
          @click="revoke"
          >撤回</el-tag
        >
        <el-tag v-else @click="publish">通过</el-tag>
      </div>
    </div>
  </el-dialog>
</template>
<script setup>
import { ref, defineExpose, watch } from "vue";
import { getArticleById, auditArticle } from "@/apis/articles";
import { getUserById } from "@/apis/user";
import { View, User } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";

const emits = defineEmits(["changeArticleStatus"]);

const props = defineProps({
  articleId: String,
});
const visible = ref(false);
const articleInfo = ref({});
const userInfo = ref({});
const userId = ref("");

// 改变弹窗的显示状态
const changeVisible = (val) => {
  visible.value = val;
};

// 发布文章
const publish = async () => {
  try {
    const res = await auditArticle({
      articleId: props.articleId,
      status: "PUBLISHED",
    });

    if (res.status === 200) {
      ElMessage.success("发布成功");
      changeVisible(false);
      emits("changeArticleStatus", {
        articleId: props.articleId,
        status: "PUBLISHED",
      });
    }
  } catch (error) {
    console.log("发布失败：", error);
  }
};

// 撤回文章
const revoke = async () => {
  try {
    const res = await auditArticle({
      articleId: props.articleId,
      status: "DELETED",
    });

    if (res.status === 200) {
      ElMessage.success("撤回成功");
      changeVisible(false);
      emits("changeArticleStatus", {
        articleId: props.articleId,
        status: "DELETED",
      });
    }
  } catch (error) {
    console.log("撤回失败：", error);
  }
};

// 根据文章id获取文章内容
const getArticleContentById = async (articleId) => {
  try {
    const { data } = await getArticleById({ articleId: articleId });
    articleInfo.value = data.articleInfo;
    userId.value = data.articleInfo.authorId;
    console.log("userId", userId.value);
  } catch (error) {
    console.log("获取文章内容失败：", error);
  }
};

// 获取用户信息
const getUserInfo = async (currentUserId) => {
  try {
    const { data } = await getUserById({ userId: currentUserId });
    userInfo.value = data.user;
  } catch (error) {
    console.log("获取用户信息失败：", error);
  }
};

// 监听articleId的变化，获取文章内容

watch(
  () => props.articleId,
  async (newVal) => {
    if (newVal) {
      await getArticleContentById(newVal);
    }
  },
  { immediate: true }
);
// 监听userId的变化，获取用户信息
watch(
  () => userId.value,
  async (newVal) => {
    if (newVal) {
      console.log("输出用户newVal：", newVal);
      await getUserInfo(newVal);
    }
  },
  { immediate: true }
);

defineExpose({
  changeVisible,
});
</script>

<style scoped lang="scss">
/* General Dialog Styling */

/* Article Container */
.articleContainer {
  display: flex;
  flex-direction: column;
  padding: 20px;
  height: 100%;
}

/* Top Section */
.Top {
  // margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f0f0f0;
  /* Light grey border for separation */
}

.articleTitle {
  font-size: 28px;
  font-weight: 600;
  color: $--color-text9;
  /* Dark grey for the title */
  margin-bottom: 10px;
}

.articleBaseInfo {
  display: flex;
  gap: 20px;
  align-items: center;
  font-size: 14px;
  color: $--color-text5;
  /* Light grey for the metadata */
}

.author {
  display: flex;
  align-items: center;
  gap: 8px;

  /* Green for author */
}

.author el-icon {
  font-size: 18px;
}

.createAt {
  color: $--color-text5;
  /* Lighter grey for the timestamp */
}

.viewCount {
  display: flex;
  align-items: center;
  gap: 8px;
}

.viewCount el-icon {
  font-size: 18px;
  color: #ff9800;
}

.viewCount span {
  font-weight: 500;
}

.tag {
  display: flex;
  align-items: center;
}

/* Center Section */
.Center {
  margin-top: 20px;
  overflow-y: auto;
}

/* Footer Section */
.Footer {
  margin-top: 30px;
  display: flex;
  justify-content: flex-end;

  .el-tag {
    cursor: pointer;
    font-size: 14px;
    padding: 6px 12px;
    border-radius: 20px;
    font-weight: 500;

    &:hover {
      opacity: 0.8;
    }
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .el-dialog {
    width: 90%;
  }

  .articleContainer {
    padding: 15px;
  }

  .articleTitle {
    font-size: 24px;
  }

  .articleBaseInfo {
    font-size: 13px;
    gap: 15px;
  }

  .author el-icon,
  .viewCount el-icon {
    font-size: 16px;
  }

  .tag {
    font-size: 13px;
    padding: 6px 12px;
  }
}
</style>
