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

// 文章点赞
export function addArticleLikeCount(data) {
  return httpRequest({
    url: `/v1/articles/likeArticle`,
    method: "post",
    data: data,
  });
}

// 取消文章点赞
export function cancelArticleLikeCount(data) {
  return httpRequest({
    url: `/v1/articles/cancelLikeArticle`,
    method: "post",
    data: data,
  });
}

// 查询用户是否点赞过文章
export function getArticleLikeStatus(params) {
  console.log(params);
  return httpRequest({
    url: `/v1/articles/queryUserLikeArticle?article_id=${params.articleId}&user_id=${params.userId}`,
    method: "get",
  });
}

// 获取文章的所有评论
export function getArticleComments(params) {
  return httpRequest({
    url: `/v1/articles/getArticleCommentList?article_id=${params.articleId}&page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}

// 发表文章评论
export function publishArticleComment(data) {
  return httpRequest({
    url: `/v1/articles/publishArticleComment`,
    method: "post",
    data: data,
  });
}

// 查询文章
export function queryArticle(params) {
  console.log(params);
  return httpRequest({
    // url: `v1/articles/queryArticle?keyword=${params.keyword}&page=${params.page}&pageSize=${params.pageSize}`,
    url: `/v1/articles/queryArticle?keyword=${params.keyword}&page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}
