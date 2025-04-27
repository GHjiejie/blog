// 设置文章的标签
export const articleTags = [
  {
    label: "前端",
    options: [
      {
        value: "html",
        label: "HTML",
      },
    ],
  },
  {
    label: "后端",
    options: [
      {
        value: "k8s",
        label: "K8S",
      },
    ],
  },
];

export const articleStatusMap = {
  DRAFT: "草稿",
  PUBLISHED: "已发布",
  DELETED: "取消发布",
  REVIEWING: "审核中",
  REVIEW_FAILED: "未通过",
  REVIEW_PASSED: "通过",
};

export const getArticleStatus = (status) => {
  return articleStatusMap[status];
};
