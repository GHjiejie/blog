<template>
  <div class="articleDetail">
    <div class="content">
      <div class="left">
        <div class="anchors">
          <div v-for="anchor in titles" @click="handleAnchorClick(anchor)">
            <a>{{ anchor.title }}</a>
          </div>
        </div>
      </div>
      <div class="right">
        <v-md-preview :text="articleDetail.content" ref="preview" />
        <div class="footer">
          <ArticleDetailFooter :author-info="authorInfo" :article-info="articleDetail"></ArticleDetailFooter>
        </div>
      </div>
    </div>
    <el-backtop :right="50" :bottom="100" />
  </div>
</template>
<script setup>
import { ref, onMounted, Comment, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getArticleById, addArticleViewCount } from "@/apis/articles";
import { getUserById } from "@/apis/user";
import ArticleDetailFooter from "@/pages/Article/components/articleDetailFooter.vue";
const route = useRoute();
// const router = useRouter();
const articleDetail = ref({});
const titles = ref([]);
const position = ref("articleDetail");

const preview = ref(null);

const authorId = ref(0);
const authorInfo = ref({});

onMounted(async () => {
  await getArticleDetail(route.params.id);
  const anchors = preview.value.$el.querySelectorAll("h1,h2,h3,h4,h5,h6");
  const titlesArray = Array.from(anchors).filter(
    (title) => !!title.innerText.trim()
  );

  if (!titlesArray.length) {
    titles.value = [];
    return;
  }

  const hTags = Array.from(
    new Set(titlesArray.map((title) => title.tagName))
  ).sort();

  titles.value = titlesArray.map((el) => ({
    title: el.innerText,
    lineIndex: el.getAttribute("data-v-md-line"),
    indent: hTags.indexOf(el.tagName),
  }));
});

const handleAnchorClick = (anchor) => {
  const { lineIndex } = anchor;

  const heading = preview.value.$el.querySelector(
    `[data-v-md-line="${lineIndex}"]`
  );

  if (heading) {
    // 注意：如果你使用的是编辑组件的预览模式,则这里的方法名改为 previewScrollToTarget
    preview.value.scrollToTarget({
      target: heading,
      scrollContainer: window,
      top: 60,
    });
  }
};
const getArticleDetail = async () => {
  const params = {
    articleId: route.params.id,
  };
  try {
    const { data } = await getArticleById(params);
    articleDetail.value = data.articleInfo;
    authorId.value = data.articleInfo.authorId;
    const res = await getUserById({ userId: data.articleInfo.authorId });
    authorInfo.value = res.data.user;
  } catch (error) { }
};
// 监听路由变化，文章浏览量+1
watch(
  () => route.params.id,
  async (newVal, oldVal) => {
    console.log("newVal", newVal);
    if (newVal !== oldVal) {
      await addArticleViewCount({ articleId: newVal });
    }
  },
  {
    immediate: true,
  }
);
</script>
<style scoped lang="scss">
.articleDetail {
  display: flex;
  flex-direction: column;

  width: 100%;

  .content {
    .left {
      position: fixed;
      width: 20%;
      max-height: 100vh;
      overflow-y: auto;
      padding: 20px;

      // 隐藏滚动条
      &::-webkit-scrollbar {
        display: none;
      }

      .anchors {
        a {
          display: block;
          margin-bottom: 10px;
          cursor: pointer;

          &:hover {
            color: var(--el-color-primary);
          }
        }
      }
    }

    .right {
      margin-left: 20%;
      border-left: 1px solid var(--el-color-primary-light-5);

      .footer {
        position: fixed;
        bottom: 0;
        width: 80%;
        z-index: 888;
      }
    }
  }
}
</style>
