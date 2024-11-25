import httpRequest from "@/request/index";

// 获取已经发布的文章列表
export function getArticleList(params) {
  return httpRequest({
    url: `/v1/web/articles/getPublishedArticleList`,
    method: "get",
    params,
  });
}

// 根据id获取文章详情
export function getArticleById(params) {
  return httpRequest({
    url: `/v1/articles/getArticleDetail/${params.articleId}`,
    method: "get",
  });
}
