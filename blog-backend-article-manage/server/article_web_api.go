package server

import (
	"context"

	articlepb "articleManage/pb/articleManage"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 获取已经发布的文章列表
func (s *ArticleServer) GetPublishedArticleList(ctx context.Context, req *articlepb.GetPublishedArticleListRequest) (*articlepb.GetPublishedArticleListResponse, error) {
	// 查询数据库获取文章列表
	logger := log.WithFields(log.Fields{
		"api": "GetPublishedArticleList",
	})
	logger.Infof("GetPublishedArticleList request with page(%d) and pageSize(%d)", req.GetPage(), req.GetPageSize())

	// 获取请求参数
	page := req.GetPage()
	pageSize := req.GetPageSize()

	// 获取文章的总数
	articleCount, err := s.DBEngine.GetArticleCount()
	if err != nil {
		logger.Errorf("failed to get article count with err(%s)", err.Error())
		return nil, err
	}
	// 判断分页参数是否合法
	if page <= 0 || pageSize <= 0 {
		logger.Errorf("page(%d) or pageSize(%d) is invalid", page, pageSize)
		return nil, status.Errorf(codes.InvalidArgument, "page or pageSize is invalid")
	}
	// 获取发布的的文章列表
	articleList, err := s.DBEngine.GetPublishedArticleList(page, pageSize)
	if err != nil {
		logger.Errorf("failed to get published article list with err(%s)", err.Error())
		return nil, err
	}
	var articles []*articlepb.ArticleInfo
	for _, article := range articleList {
		articles = append(articles, &articlepb.ArticleInfo{
			ArticleId: article.ID,
			Title:     article.Title,
			AuthorId:  article.AuthorId,
			Summary:   article.Summary,
			Content:   article.Content,
			Tag:       article.Tag,
			ImageUrl:  article.ImageURL,
			Status:    article.Status,
			CreatedAt: timestamppb.New(article.CreatedAt),
			UpdatedAt: timestamppb.New(article.UpdatedAt),
		})
	}

	return &articlepb.GetPublishedArticleListResponse{
		ArticleList: articles,
		Total:       int32(articleCount),
	}, nil
}
