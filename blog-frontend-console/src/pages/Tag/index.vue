<template>
  <div class="tag-management">
    <el-collapse
      v-model="activeNames"
      @change="handleChange"
      class="tag-collapse"
    >
      <el-collapse-item title="前端" name="1">
        <div class="tag-container">
          <el-tag v-for="tag in frontEnd" :key="tag.tagId" class="tag"
            ><span @click="deletetag(tag.tagId)">{{
              tag.tagName
            }}</span></el-tag
          >

          <el-popover placement="right" :width="250" trigger="click">
            <template #reference>
              <el-button style="margin-right: 16px" text type="primary"
                >增加标签</el-button
              >
            </template>
            <el-form>
              <el-form-item label="">
                <el-input
                  v-model="tagForm.tagName"
                  placeholder="请输入标签名称"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleAddTag('1')"
                  >添加</el-button
                >
                <el-button @click="resetForm">重置</el-button>
              </el-form-item>
            </el-form>
          </el-popover>
        </div>
      </el-collapse-item>
      <el-collapse-item title="后端" name="2">
        <div class="tag-container">
          <el-tag v-for="tag in backEnd" :key="tag.tagId" class="tag"
            ><span @click="deletetag(tag.tagId)">{{
              tag.tagName
            }}</span></el-tag
          >

          <el-popover placement="right" :width="250" trigger="click">
            <template #reference>
              <el-button style="margin-right: 16px" text type="primary"
                >增加标签</el-button
              >
            </template>
            <el-form>
              <el-form-item label="">
                <el-input
                  v-model="tagForm.tagName"
                  placeholder="请输入标签名称"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleAddTag('2')"
                  >添加</el-button
                >
                <el-button @click="resetForm">重置</el-button>
              </el-form-item>
            </el-form>
          </el-popover>
        </div>
      </el-collapse-item>
      <el-collapse-item title="测试" name="3">
        <div class="tag-container">
          <el-tag v-for="tag in test" :key="tag.tagId" class="tag"
            ><span @click="deletetag(tag.tagId)">{{
              tag.tagName
            }}</span></el-tag
          >

          <el-popover placement="right" :width="250" trigger="click">
            <template #reference>
              <el-button style="margin-right: 16px" text type="primary"
                >增加标签</el-button
              >
            </template>
            <el-form>
              <el-form-item label="">
                <el-input
                  v-model="tagForm.tagName"
                  placeholder="请输入标签名称"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleAddTag('3')"
                  >添加</el-button
                >
                <el-button @click="resetForm">重置</el-button>
              </el-form-item>
            </el-form>
          </el-popover>
        </div>
      </el-collapse-item>
      <el-collapse-item title="运维" name="4">
        <div class="tag-container">
          <el-tag v-for="tag in operation" :key="tag.tagId" class="tag"
            ><span @click="deletetag(tag.tagId)">{{
              tag.tagName
            }}</span></el-tag
          >

          <el-popover placement="right" :width="250" trigger="click">
            <template #reference>
              <el-button style="margin-right: 16px" text type="primary"
                >增加标签</el-button
              >
            </template>
            <el-form>
              <el-form-item label="">
                <el-input
                  v-model="tagForm.tagName"
                  placeholder="请输入标签名称"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleAddTag('4')"
                  >添加</el-button
                >
                <el-button @click="resetForm">重置</el-button>
              </el-form-item>
            </el-form>
          </el-popover>
        </div>
      </el-collapse-item>
      <el-collapse-item title="其他" name="5">
        <div class="tag-container">
          <el-tag v-for="tag in other" :key="tag.tagId" class="tag"
            ><span @click="deletetag(tag.tagId)">{{
              tag.tagName
            }}</span></el-tag
          >

          <el-popover placement="right" :width="250" trigger="click">
            <template #reference>
              <el-button style="margin-right: 16px" text type="primary"
                >增加标签</el-button
              >
            </template>
            <el-form>
              <el-form-item label="">
                <el-input
                  v-model="tagForm.tagName"
                  placeholder="请输入标签名称"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleAddTag('5')"
                  >添加</el-button
                >
                <el-button @click="resetForm">重置</el-button>
              </el-form-item>
            </el-form>
          </el-popover>
        </div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { getTags } from "@/apis/articles";
import { Plus } from "@element-plus/icons-vue";
import { addTag, deleteTag } from "@/apis/articles";
import { ElMessage, ElMessageBox } from "element-plus";

const tagForm = ref({
  tagName: "",
});

const resetForm = () => {
  tagForm.value.tagName = "";
};

const deletetag = (tagId) => {
  // 弹窗提示是否删除
  ElMessageBox.confirm("此操作将永久删除该标签, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      // 确认删除
      try {
        const res = await deleteTag({
          tagId,
        });
        if (res.status === 200) {
          ElMessage.success("删除成功");
          // 页面刷新
          window.location.reload();
        }
      } catch (error) {
        console.error("删除标签失败", error);
      }
    })
    .catch(() => {
      // 取消删除
    });
};

const handleAddTag = async (type) => {
  if (!tagForm.value.tagName) {
    ElMessage.error("请输入标签名称");
    return;
  }
  try {
    const res = await addTag({
      tagName: tagForm.value.tagName,
      tagCategoryId: type,
    });
    if (res.status === 200) {
      ElMessage.success("添加成功");
      // 页面刷新
      window.location.reload();
    }
  } catch (error) {
    console.error("添加标签失败", error);
  }
};

// 将不同分类的标签数组合并为一个响应式对象
const tagsByCategory = ref({
  1: [], // 前端
  2: [], // 后端
  3: [], // 测试
  4: [], // 运维
  default: [], // 其他
});

// 动态计算分类后的标签数组
const frontEnd = computed(() => tagsByCategory.value["1"]);
const backEnd = computed(() => tagsByCategory.value["2"]);
const test = computed(() => tagsByCategory.value["3"]);
const operation = computed(() => tagsByCategory.value["4"]);
const other = computed(() => tagsByCategory.value["default"]);

const activeNames = ref(["1"]);
const handleChange = (val) => {
  console.log(val);
};

// 不需要单独的 page 和 pageSize ref，直接在 API 请求中配置
onMounted(async () => {
  try {
    const { data } = await getTags({ page: 1, pageSize: 100 }); // 一次性获取所有标签
    const tagsList = data.tagList;

    // 使用 forEach 处理标签分类
    tagsList.forEach((tag) => {
      const category = tag.tagCategoryId || "default"; // 如果没有 tagCategoryId，则归为 default
      if (!tagsByCategory.value[category]) {
        tagsByCategory.value[category] = []; // 如果分类不存在，则初始化一个新数组
      }
      tagsByCategory.value[category].push(tag); // 将标签添加到对应的分类数组中
    });
  } catch (error) {
    console.error("获取标签失败", error);
  }
});
</script>

<style scoped lang="scss">
.tag-management {
  padding: 20px;
  font-family: Arial, sans-serif;

  // .Top {
  //   position: fixed;
  //   top: 50%;
  //   right: 2%;
  //   z-index: 999;

  //   .addFileBtn {
  //     background-color: $--color-primary-opacity7;
  //     color: $--color-text;
  //     border: none;
  //     padding: 10px;
  //     border-radius: 54% 46% 21% 79% / 52% 46% 54% 48%;

  //     box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  //     transition: all 0.3s;
  //     cursor: pointer;

  //     &:hover {
  //       background-color: $--color-primary;
  //     }
  //   }
  // }

  h1 {
    font-size: 24px;
    margin-bottom: 20px;
    color: #333;
  }

  .tag-collapse {
    border: none; // 移除 collapse 容器的边框
    background: transparent; // 移除 collapse 容器的背景色

    .el-collapse-item {
      border: none; // 移除 item 的边框
      margin-bottom: 10px; // 增加 item 之间的间距
      // background-color: #f9f9f9; // 设置 item 的背景色
      border-radius: 4px; // 设置边框的圆角
      // box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1); // 添加/轻微阴影

      .el-collapse-item__header {
        font-size: 16px;
        color: #555;
        padding: 12px 20px; // 增加 header 的 padding
        transition: background-color 0.3s ease; // 添加过渡效果

        &:hover {
          background-color: #eee; // hover 时改变背景色
          cursor: pointer; // 修改鼠标样式
        }
      }

      .el-collapse-item__content {
        padding: 10px 20px 20px; // 增加 content 的 padding
      }
    }
  }

  .tag-container {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    justify-content: flex-start;

    gap: 8px;

    &:hover {
      cursor: pointer;
    }
  }

  .tag {
    font-size: 14px;
    padding: 6px 12px;
    border-radius: 20px;
    background-color: #e6f7ff; // 浅蓝色背景
    color: #1890ff; // 主题色文字
    border: 1px solid #91d5ff; // 加一个淡淡的边框
    margin-bottom: 5px;
    transition: all 0.2s ease-in-out;

    &:hover {
      background-color: #bae7ff;
      color: #096dd9;
      transform: scale(1.05);
    }
  }
}
</style>
