<template>
  <div class="search">
    <el-input
      v-model="keywords"
      style="width: 90%"
      placeholder="输入关键字"
      :prefix-icon="Search"
      @keyup.enter="handleSearch"
      ref="searchRef"
    />
    <div class="articleList">
      <template v-if="searchResults.length == 0">
        <el-empty :image-size="200" />
      </template>
      <template v-else>
        <div
          class="ArticleItem"
          v-for="(item, index) in searchResults"
          :key="index"
        >
          <!-- 文章封面图片 -->
          <div class="left">
            <!-- <div class="articleCoverImg">
              <img :src="item.imageUrl" alt="coverImg" />
            </div> -->
          </div>
          <div class="right">
            <div class="title">
              {{ item.title }}
            </div>
            <div class="summary" @click="goArticleDetail(item.articleId)">
              <p>{{ item.summary }}</p>
            </div>
            <div class="articleInfo">
              <div class="creatAt">
                <span>{{ item.createdAt }}</span>
              </div>

              <div class="tag">
                <el-tag type="success">{{ item.tag }}</el-tag>
              </div>
            </div>

            <!-- 文章浏览消息包括点赞，评论，浏览量等 -->
            <div class="articleView">
              <FeedBack
                :articleInfo="item"
                :activePosition="position"
              ></FeedBack>
            </div>
          </div>
        </div>

        <div class="loadmore">
          <!-- <el-button type="primary" size="medium" @click="loadMore" text>{{ loadInfo }}</el-button> -->
          <el-pagination
            size="small"
            background
            layout="prev, pager, next"
            :total="total"
            class="mt-4"
            :page-size="5"
          />
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from "vue";
import { Search } from "@element-plus/icons-vue";
import { queryArticle } from "@/apis/articles";
import { useRouter } from "vue-router";
const router = useRouter();
const page = ref(1);
const pageSize = ref(100);
const keywords = ref("");
const searchResults = ref([]);

const total = ref(0);

const searchRef = ref(null);

const goArticleDetail = (articleId) => {
  router.push(`/article/${articleId}`);
};
const performSearch = async () => {
  const res = await queryArticle({
    keyword: keywords.value,
    page: page.value,
    pageSize: pageSize.value,
  });

  searchResults.value = res.data.articleList;
  total.value = res.data.total;
};
const handleSearch = async () => {
  await performSearch();
  console.log("Search results:", searchResults.value);
};

onMounted(async () => {
  await nextTick();
  searchRef.value.focus();
});
</script>

<style scoped lang="scss">
// 设置滚动条样式
::-webkit-scrollbar {
  width: 0px;
}

.search {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 85vh;
  background-color: #f9f9f9;
  border-radius: 20px;
  overflow: auto;

  .el-input {
    margin: 15px;
    height: 50px;
    --el-border-radius-base: 50px;
  }

  :deep(.el-input__wrapper) {
    background-color: #ebecf0;
  }
}

.articleList {
  margin-top: 150x;
  width: 100%;
  flex: 1;
  overflow: auto;

  .ArticleItem {
    display: flex;
    gap: 20px; // 增加项目之间的间距
    margin-bottom: 25px; // 增加底部间距
    padding: 20px; // 增加内边距
    background: #f9f9f9;
    border-radius: 12px; // 圆角更明显
    border: 1px solid #f0f0f0; // 增加边框
    transition: all 0.3s ease;

    &:hover {
      cursor: pointer;
      border-color: var(--el-color-primary-light-5); // 鼠标悬停时边框颜色变化
    }

    .left {
      flex-shrink: 0;

      .articleCoverImg {
        width: 150px; // 调整封面图像大小
        height: 150px; // 同上
        overflow: hidden;
        border-radius: 12px; // 与外容器一致的圆角风格

        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
          transition: transform 0.3s;

          &:hover {
            transform: scale(1.05); // 缩放效果
          }
        }
      }
    }

    .right {
      flex: 1;
      display: flex;
      flex-direction: column;
      justify-content: space-between;

      .title {
        font-size: 20px; // 加大标题字体
        font-weight: bold;
        color: #333;
        margin-bottom: 5px;
        transition: color 0.3s;
      }

      .summary {
        p {
          font-size: 16px; // 加大摘要字体
          color: #555;
          line-height: 1.6;
          margin-bottom: 10px;
          // 设置显示的行数
          overflow: hidden;
          text-overflow: ellipsis;
          display: -webkit-box;
          -webkit-line-clamp: 3;
          line-clamp: 3;
          -webkit-box-orient: vertical;
          word-wrap: break-word;

          &:hover {
            color: var(--el-color-primary); // 鼠标悬停时颜色变化
          }
        }
      }

      .articleInfo {
        display: flex;
        gap: 15px; // 调整信息项之间的间隔
        font-size: 13px;
        color: #888;

        .creatAt {
          display: flex;
          align-items: center;
          gap: 5px; // 调整图标与文字之间的间隔

          span {
            color: #555;
          }
        }
      }

      .articleView {
        display: flex;
        gap: 15px;
        justify-content: flex-end; // 调整位置到右侧

        div {
          display: flex;
          align-items: center;
          gap: 3px; // 稍微增加间隔
          font-size: 13px;
          color: #777;

          span {
            color: #555;
          }
        }
      }
    }
  }
}

.loadmore {
  display: flex;
  justify-content: center;
  margin-bottom: 10px; // 增加顶部间距
}
</style>
