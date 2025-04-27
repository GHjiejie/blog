<template>
  <div class="Article">
    <el-form :model="form" label-width="auto" style="max-width: 80%">

      <div class="BaseInfo">
        <div class="left">
          <el-form-item label="文章标题" class="title">
            <el-input v-model="articleFrom.title" class="title" />
          </el-form-item>
          <el-form-item label="文章摘要" class="summary">
            <el-input v-model="articleFrom.summary" class="summary" />
          </el-form-item>
          <el-form-item label="文章标签" class="tag">
            <el-select v-model="articleFrom.tag" placeholder="标签" style="width: 240px">
              <el-option-group v-for="group in articleTags" :key="group.label" :label="group.label">
                <el-option v-for="item in group.options" :key="item.value" :label="item.label" :value="item.label" />
              </el-option-group>
            </el-select>
          </el-form-item>
        </div>
        <div class="right">
          <!-- <input type="file" style="display: none;" ref="uploadImgRef" @change="fileChange"></input>
          <el-button plain class="uploadFileBtn" @click="$refs.uploadImgRef.click()">
            <el-icon :size="20">
              <Plus />
            </el-icon>
          </el-button> -->
        </div>
      </div>
      <!-- 文件的主要内容，可以在线编辑上传，也可以选择文件 -->
      <div class="mainInfo">
        <div class="mainInfoTop">
          <div class="left">
            <div style="max-width: 600px">
              <el-alert title="提示消息" type="warning" :closable="false">
                <template #default>
                  <p>1.文章内容支持markdown语法，可以直接在线编辑上传</p>
                  <p>2.文章内容支持在线文件上传，可以选择在线文件进行编辑</p>
                  <p>3.暂不支持本地图片上传</p>
                </template>
              </el-alert>
            </div>
          </div>
          <div class="right">
            <el-dropdown>
              <el-button type="success"> 上传 </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="uploadLocalFile">本地文件</el-dropdown-item>
                  <el-dropdown-item @click="uploadOriginFile">在线文件</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
        <div class="mainInfoCenter">
          <v-md-editor v-model="articleFrom.content" height="600px"></v-md-editor>
        </div>
      </div>
      <!-- 下面的消息是确认发布者的ID是否正确 -->
      <div class="confirmInfo">
        <el-tag type="success">作者ID:{{ articleFrom.authorId }}</el-tag>
        <el-button @click="handlePublish">发布</el-button>
      </div>
    </el-form>
    <!-- 虚拟按钮 -->
    <input type="file" style="display: none;" ref="uploadImgRef" @change="imgChange"></input>
    <input type="file" style="display: none;" ref="uploadLocalFileRef" @change="fileChange" accept=".md">
  </div>
  <fileList ref="fileListRef" :file-id="viewFileId" @select-online-file="handleOnlineFile"></fileList>
</template>
<script setup>
import { reactive, ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus'
import { Plus, ArrowDown } from '@element-plus/icons-vue'
import fileList from './fileList.vue';
import { decodeBase64 } from "@/utils/fileFilter";
import { getFileById } from '@/apis/file'
import { publishArticle, getTags } from '@/apis/articles'
import cache from '@/utils/cache'
import router from '@/router';

const articleFrom = reactive({
  title: '',
  tag: '',
  content: '',
  summary: '',
  authorId: cache.sessionGet('userId'),
})
const uploadImgRef = ref(null);
const uploadLocalFileRef = ref(null);
const fileListRef = ref(null);

const page = ref(1)
const pageSize = ref(100)

const articleTags = ref([{
  label: '前端',
  id: 1,
  options: []
}, {
  label: '后端',
  id: 2,
  options: []
}, {
  label: '测试',
  id: 3,
  options: []
}, {
  label: '运维',
  id: 4,
  options: []
},
{
  label: '其他',
  id: 5,
  options: []
}])
// 处理本地文件上传
const fileChange = (e) => {
  const file = e.target.files[0];
  const reader = new FileReader();
  reader.onload = (e) => {
    articleFrom.content = e.target.result;
  }
  reader.readAsText(file);
}

// 上传本地文件
const uploadLocalFile = () => {
  uploadLocalFileRef.value.click();
}

// 上传在线文件
const uploadOriginFile = () => {
  fileListRef.value.changeVisible(true);
}

// 处理在线文件上传
const handleOnlineFile = async (fileId) => {
  await getFileContent(fileId)
}
// 根据ID获取文件的内容
const getFileContent = async (fileId) => {
  try {
    const { data } = await getFileById({ fileId });
    articleFrom.content = decodeBase64(data.fileInfo.content);


  } catch (error) {
    console.log('getFileContent error', error)
  }
}

const handlePublish = async () => {
  try {
    await publishArticle(articleFrom);
    ElMessage.success('初稿发布成功，请审核通过后发布');
    router.push({
      name:'ArticleList'
    })
  } catch (error) {
    console.log('发布失败', error)
  }

}

const initTagList = (tags) => {
  // 根据tagCategoryId将标签放在对应的分类中
  tags.forEach(tag => {
    articleTags.value[tag.tagCategoryId - 1].options.push({
      label: tag.tagName,
      value: tag.tagId
    })
  })


}

onMounted(async () => {
  try {
    const { data } = await getTags({ page: 1, pageSize: 100 });
    initTagList(data.tagList)
  } catch (error) {
    console.log('获取标签失败', error)
  }

})
</script>

<style scoped lang="scss">
/* General styling */
.Article {
  background: --color-background;
  padding: 30px;
  max-height: calc(100vh - 50px);
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

/* Base Info Section */
.BaseInfo {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.left,
.right {
  width: 48%;
}

.title,
.summary,
.tag {
  margin-bottom: 15px;
}

/* Input styling */
.el-input,
.el-select {
  border-radius: 8px;
  border: 1px solid #e1e4e8;
  transition: all 0.3s ease-in-out;
}

/* Upload button styling */
// .uploadFileBtn {
//   background-color: #409eff;
//   border-radius: 8px;
//   color: $--color-text;
//   padding: 10px 15px;
//   cursor: pointer;
//   transition: background-color 0.3s;
// }

// .uploadFileBtn:hover {
//   background-color: #66b1ff;
// }

/* Dropdown button styling */
// .el-button.success {
//   background-color: #67c23a;
//   border-color: #67c23a;
//   color: white;
//   border-radius: 8px;
//   padding: 10px 20px;
// }

// .el-button.success:hover {
//   background-color: #85ce61;
//   border-color: #85ce61;
// }

/* Dropdown menu item */




/* Alert styling */
.el-alert {
  border-radius: 8px;
  margin-bottom: 20px;
}



.mainInfoTop {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;


}

.mainInfoCenter {
  margin-top: 20px;
}

/* Confirm section styling */
.confirmInfo {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 30px;
}

.el-tag {
  font-size: 14px;
  padding: 10px;
  border-radius: 12px;
  background-color: #e0f7ea;
  color: $--color-text2;
  transition: all 0.3s ease;
}

.el-tag:hover {
  background-color: #c8e6c9;
  transform: scale(1.05);
  cursor: pointer;
}

.el-button {
  background-color: $--color-primary;
  border-color: $--color-border;
  color: white;
  padding: 10px 20px;
  border-radius: 8px;
  transition: background-color 0.3s ease;
}

// .el-button:hover {
//   background-color: $--bg-color-hover;
//   border-color: $--color-border3;
// }

/* File list component */
.fileListRef {
  margin-top: 30px;
  border-top: 2px solid #f0f0f0;
  padding-top: 20px;
}

.fileListRef .file-list-item {
  border: 1px solid #e1e4e8;
  border-radius: 8px;
  padding: 10px;
  margin-bottom: 10px;
  transition: all 0.3s ease;
}

.fileListRef .file-list-item:hover {
  background-color: #f5f5f5;
  transform: translateX(5px);
}

/* Styling for overall content */
.mainInfo {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.mainInfoTop,
.mainInfoCenter {
  width: 100%;
}

/* Adding subtle "梅花" (Plum Blossom) feel to elements */
.el-input,
.el-select,
.el-button,
.el-tag {
  // box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.1);
  /* subtle shadow */
  border-radius: 8px;
}



.el-button {
  transition: all 0.3s ease;
}

.el-button:hover {
  transform: translateY(-2px);
  /* Slight lift effect on hover */
}

.el-alert {
  border-left: 4px solid $--color-primary;
  background-color: #f1f9ff;
  /* Soft background */
}

/* Responsive Design */
@media (max-width: 768px) {
  .BaseInfo {
    flex-direction: column;
  }

  .left,
  .right {
    width: 100%;
  }

  .el-input,
  .el-select {
    width: 100%;
  }
}
</style>



<!-- <style scoped lang="scss">
.Article {

  width: 100%;
  margin: 20px auto;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #f9f9f9;

  .BaseInfo {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;

    .left {
      flex: 1;
      margin-right: 20px;
      display: flex;
      flex-direction: column;





      input,
      select {
        width: auto;
        padding: 10px;
        margin-bottom: 10px;
        border: none;

        background: none;
        border-bottom: 1px solid var(--el-disabled-text-color);

        &:focus {
          border-color: #007bff;
          outline: none;
        }
      }


    }

    .right {
      img {
        width: 250px;
        height: 250px;
        border-radius: 4px;
        object-fit: cover;
      }

      .uploadFileBtn {
        width: 200px;
        height: 200px;
        border-radius: 4px;
        object-fit: cover;
        background-color: white;
        border: 1px dashed var(--el-disabled-text-color);
        transition: all 0.3s;



        &:hover {
          border: 1px solid var(--el-color-primary);
        }
      }
    }
  }

  .mainInfo {
    textarea {
      width: 100%;
      height: 200px;
      padding: 10px;
      border: 1px solid #ccc;
      border-radius: 4px;
      font-size: 16px;
      resize: none;

      &:focus {
        border-color: #007bff;
        outline: none;
      }
    }
  }

  .confirmInfo {
    margin-top: 20px;
    font-size: 16px;


    button {
      padding: 10px 20px;
      border: none;
      border-radius: 4px;
      background-color: #28a745;
      color: white;
      font-size: 16px;
      cursor: pointer;

      &:hover {
        background-color: #218838;
      }
    }
  }
}
</style> -->