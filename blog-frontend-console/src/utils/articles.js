// 设置文章的标签
export const articleTags = [
  {
    label: "前端",
    options: [
      {
        value: "html",
        label: "HTML",
      },
      {
        value: "css",
        label: "CSS",
      },
      {
        value: "javascript",
        label: "JavaScript",
      },
      {
        value: "vue",
        label: "Vue",
      },
      {
        value: "react",
        label: "React",
      },
      {
        value: "angular",
        label: "Angular",
      },
      {
        value: "node",
        label: "Node",
      },
      {
        value: "webpack",
        label: "Webpack",
      },
      {
        value: "typescript",
        label: "TypeScript",
      },
      {
        value: "es6",
        label: "ES6",
      },
    ],
  },
  {
    label: "后端",
    options: [
      {
        value: "golang",
        label: "Golang",
      },
      {
        value: "mysql",
        label: "MySQL",
      },
      {
        value: "linux",
        label: "Linux",
      },
      {
        value: "docker",
        label: "Docker",
      },
      {
        value: "k8s",
        label: "K8S",
      },
    ],
  },
  {
    label: "工具",
    options: [
      {
        value: "git",
        label: "Git",
      },
      {
        value: "vscode",
        label: "VSCode",
      },
      {
        value: "chrome",
        label: "Chrome",
      },
    ],
  },
  {
    label: "生活",
    options: [
      {
        value: "travel",
        label: "旅行",
      },
      {
        value: "movie",
        label: "电影",
      },
      {
        value: "music",
        label: "音乐",
      },
    ],
  },
  {
    label: "其他",
    options: [
      {
        value: "interview",
        label: "面试",
      },
      {
        value: "other",
        label: "其他",
      },
    ],
  },
];

export const articleStatusMap = {
  DRAFT: "草稿",
  PUBLISHED: "已发布",
  DELETED: "驳回",
  REVIEWING: "审核中",
  REVIEW_FAILED: "未通过",
  REVIEW_PASSED: "通过",
};

export const getArticleStatus = (status) => {
  return articleStatusMap[status];
};
