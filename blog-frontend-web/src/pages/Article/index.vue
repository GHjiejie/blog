<template>
  <div class="articleContainer">
    <div class="articleList">
      <div class="ArticleItem" v-for="(item, index) in articleList" :key="index">
        <!-- 文章封面图片 -->
        <div class="left">
          <div class="articleCoverImg">
            <img :src="item.imageUrl" alt="coverImg" />
          </div>
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
            <FeedBack :articleInfo="item" :activePosition="position"></FeedBack>
          </div>
        </div>
      </div>
      <div class="loadmore">
        <el-button type="primary" size="medium" @click="loadMore" text>{{ loadInfo }}</el-button>
      </div>
    </div>
    <el-backtop :right="50" :bottom="100" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getArticleList } from "@/apis/articles";
import { useRouter } from "vue-router";
import FeedBack from "@/components/Feedback/index.vue";
const router = useRouter();
const page = ref(1);
const pageSize = ref(10);
const articleList = ref([]);
const position = ref("articleList");
const loadInfo = ref("加载更多");






const goArticleDetail = (articleId) => {
  router.push(`/article/${articleId}`);
};

const loadMore = async () => {
  page.value++;
  try {
    const { data } = await getArticleList({
      page: page.value,
      pageSize: pageSize.value,
    });
    console.log('data', data)
    if (data.articleList.length === 0) {
      loadInfo.value = "没有更多了";
      return;
    } else {
      articleList.value = articleList.value.concat(data.articleList);
    }

  } catch (error) {
    console.log(error);
  }
};

onMounted(async () => {
  try {
    const { data } = await getArticleList({
      page: page.value,
      pageSize: pageSize.value,
    });
    console.log(data);
    if (data.articleList.length === 0) {
      ElMessage({
        message: "暂无数据",
        type: "info",
      });
    }
    articleList.value = data.articleList;
  } catch (error) {
    console.log(error);
  }
});
</script>

<style scoped lang="scss">
.articleContainer {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;

  .articleList {
    width: 100%;
    max-width: 900px; // 增加最大宽度
    border-radius: 12px; // 增加边角圆度
    overflow: hidden;

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
    margin-top: 20px; // 增加顶部间距
  }
}
</style>

<!-- <style scoped lang="scss">
.articleContainer {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  background-color: #f9f9f9;

  .articleList {
    width: 100%;
    max-width: 800px;

    .ArticleItem {
      display: flex;
      gap: 15px;
      margin-bottom: 20px;
      padding: 15px;
      background: #ffffff;
      border-radius: 8px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      transition: transform 0.2s, box-shadow 0.2s;

      &:hover {
        transform: translateY(-5px);
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
      }

      .left {
        flex-shrink: 0;

        .articleCoverImg {
          width: 120px;
          height: 120px;
          overflow: hidden;
          border-radius: 8px;

          img {
            width: 100%;
            height: 100%;
            object-fit: cover;
          }
        }
      }

      .right {
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        .title {
          font-size: 18px;
          font-weight: bold;
          color: #333;
          margin-bottom: 8px;
        }

        .summary {
          font-size: 14px;
          color: #555;
          line-height: 1.6;
          margin-bottom: 10px;
        }

        .articleInfo {
          display: flex;
          gap: 10px;
          font-size: 12px;
          color: #888;

          span {
            display: inline-block;

            &:not(:last-child)::after {
              summary: '|';
              margin-left: 5px;
              color: #ccc;
            }
          }
        }

        .articleView {
          display: flex;
          gap: 15px;
          margin-top: 10px;

          div {
            display: flex;
            align-items: center;
            gap: 5px;
            font-size: 12px;
            color: #777;

            el-icon {
              font-size: 16px;
              color: #888;
            }

            span {
              color: #555;
            }
          }
        }
      }
    }
  }
}
</style> -->
