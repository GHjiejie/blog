import httpRequest from "@/request/index";

// 获取已经发布的文章列表
export function getArticleList(params) {
  return httpRequest({
    url: `/v1/web/articles/getPublishedArticleList?page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}

// 根据id获取文章详情
export function getArticleById(params) {
  return httpRequest({
    url: `/v1/articles/getArticleDetail/${params.articleId}`,
    method: "get",
  });
}

// 文章浏览数增加
export function addArticleViewCount(data) {
  return httpRequest({
    url: `/v1/articles/viewArticle`,
    method: "post",
    data: data,
  });
}
