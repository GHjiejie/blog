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

// 删除文章
export function deleteArticle(params) {
  console.log(params);
  return httpRequest({
    url: `/v1/articles/deleteArticle/${params.articleId}`,
    method: "delete",
  });
}

// 根据id获取文章详情
export function getArticleById(params) {
  return httpRequest({
    url: `/v1/articles/getArticleDetail/${params.articleId}`,
    method: "get",
  });
}

//审核文章
export function auditArticle(data) {
  return httpRequest({
    url: `/v1/articles/reviewArticle`,
    method: "post",
    data,
  });
}

// 获取所有的Tag
export function getTags(params) {
  return httpRequest({
    url: `/v1/articles/getArticleTagList?page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}

// 新增Tag
export function addTag(data) {
  return httpRequest({
    url: `/v1/articles/addArticleTag`,
    method: "post",
    data,
  });
}

// 删除Tag
export function deleteTag(params) {
  return httpRequest({
    url: `/v1/articles/deleteArticleTag/${params.tagId}`,
    method: "delete",
  });
}
