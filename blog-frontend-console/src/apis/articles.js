import httpRequest from "@/request/index";

// 获取文件列表
export function getArticleList(params) {
  return httpRequest({
    url: `/v1/articles/getArticleList?page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}

// 发布文章
export function publishArticle(data) {
  return httpRequest({
    url: "/v1/articles/publishArticle",
    method: "post",
    data,
  });
}
