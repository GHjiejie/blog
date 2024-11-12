package server

import (
	"context"
	"strings"

	articlepb "articleManage/pb/articleManage"
	"articleManage/pkg/db"
	"articleManage/validate"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 发布文章
func (s *ArticleServer) PublishArticle(ctx context.Context, req *articlepb.PublishArticleRequest) (*articlepb.PublishArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "PublishArticle",
	})
	// 获取请求参数
	title := req.GetTitle()
	content := req.GetContent()
	authorId := req.GetAuthorId()
	summary := req.GetSummary()
	tag := req.GetTag()
	imageUrl := req.GetImageUrl()
	status := req.GetStatus()
	// logger.Infof("title: %s, content: %s, authorId: %d, summary: %s, imageUrl: %s, status: %s", title, content, authorId, summary, imageUrl, status)
	// 创建文章消息结构体
	articleInfo := db.Article{
		Title:     title,
		Content:   content,
		AuthorId:  authorId,
		Summary:   summary,
		Tag:       tag,
		ImageURL:  imageUrl,
		Status:    status,
		CreatedAt: timestamppb.Now().AsTime(),
		UpdatedAt: timestamppb.Now().AsTime(),
	}
	article, err := s.DBEngine.CreateArticle(articleInfo)
	if err != nil {
		logger.Errorf("failed to create article with err(%s)", err.Error())
		return nil, err
	}
	logger.Infof("article created: %v", article)

	// return nil, status.Errorf(codes.Unimplemented, "method PublishArticle not implemented")
	return &articlepb.PublishArticleResponse{
		Message: "article created successfully",
	}, nil
}

// 更新文章
func (s *ArticleServer) UpdateArticle(ctx context.Context, req *articlepb.UpdateArticleRequest) (*articlepb.UpdateArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "UpdateArticle",
	})
	// 获取请求参数
	articleId := req.GetArticleId()
	title := req.GetTitle()
	content := req.GetContent()
	summary := req.GetSummary()
	tag := req.GetTag()
	fileStatus := req.GetStatus()
	imageUrl := req.GetImageUrl()
	authorId := req.GetAuthorId()

	// 根据id去查询这个文章的消息
	article, err := s.DBEngine.GetArticleDetail(articleId)
	if err != nil {
		logger.Errorf("failed to get article detail with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article detail: %v", err)
	}
	// 首先校验这个参数是否合法(这个太多了,不处理了,前端处理就好了)
	fileInfo := db.Article{
		ID:        articleId,
		Title:     title,
		Content:   content,
		AuthorId:  authorId,
		Summary:   summary,
		Tag:       tag,
		ImageURL:  imageUrl,
		Status:    fileStatus,
		CreatedAt: article.CreatedAt,
		UpdatedAt: timestamppb.Now().AsTime(),
	}
	// 更新文章
	if err := s.DBEngine.UpdateArticle(fileInfo); err != nil {
		logger.Errorf("failed to update article with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}

	return &articlepb.UpdateArticleResponse{
		Message: "article updated successfully",
	}, nil
}

// 删除文章
func (s *ArticleServer) DeleteArticle(ctx context.Context, req *articlepb.DeleteArticleRequest) (*articlepb.DeleteArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "DeleteArticle",
	})
	// 获取请求参数
	articleId := req.GetArticleId()
	// 首先校验这个参数是否合法
	if err := validate.DeleteArticle(req); err != nil {
		logger.Errorf("failed to delete article with err(%s)", err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}
	// 然后获取所以文章的总数
	articleCount, err := s.DBEngine.GetArticleCount()
	if err != nil {
		logger.Errorf("failed to get article count with err(%s)", err.Error())
		return nil, err
	}
	// 判断文章ID是否在合法范围内
	if articleId > articleCount {
		logger.Errorf("articleId(%d) is out of range", articleId)
		return nil, status.Errorf(codes.InvalidArgument, "articleId is out of range")
	}
	// 删除文章
	if err := s.DBEngine.DeleteArticle(articleId); err != nil {
		logger.Errorf("failed to delete article with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to delete article: %v", err)
	}

	return &articlepb.DeleteArticleResponse{
		Message: "article deleted successfully",
	}, nil
}

// 获取文章列表
func (s *ArticleServer) GetArticleList(ctx context.Context, req *articlepb.GetArticleListRequest) (*articlepb.GetArticleListResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "GetArticleList",
	})
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
	// 获取文章列表
	articleList, err := s.DBEngine.GetArticleList(page, pageSize)
	if err != nil {
		logger.Errorf("failed to get article list with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article list: %v", err)
	}
	var articles []*articlepb.ArticleListInfo
	for _, article := range articleList {
		articles = append(articles, &articlepb.ArticleListInfo{
			ArticleId: article.ID,
			Title:     article.Title,
			AuthorId:  article.AuthorId,
			Summary:   article.Summary,
			Tag:       article.Tag,
			ImageUrl:  article.ImageURL,
			Status:    article.Status,
			CreatedAt: timestamppb.New(article.CreatedAt),
			UpdatedAt: timestamppb.New(article.UpdatedAt),
		})
	}

	return &articlepb.GetArticleListResponse{
		ArticleList: articles,
		Total:       int32(articleCount),
	}, nil
}

// 查询文章
func (s *ArticleServer) QueryArticle(ctx context.Context, req *articlepb.QueryArticleRequest) (*articlepb.QueryArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "QueryArticle",
	})
	///目前只支持tag查询
	keyword := req.GetKeyword()
	// 去除两端空格
	keyword = strings.TrimSpace(keyword)
	// 如果是英文,则全部转为小写
	keyword = strings.ToLower(keyword)
	logger.Infof("keyword: %s", keyword)
	if keyword == "" {
		logger.Errorf("keyword is required")
		return nil, status.Errorf(codes.InvalidArgument, "keyword is required")
	}

	// 查询文章
	articleList, err := s.DBEngine.GetArticleListByTag(keyword)
	if err != nil {
		logger.Errorf("failed to get article list with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article list: %v", err)
	}
	var articles []*articlepb.ArticleListInfo
	for _, article := range articleList {
		articles = append(articles, &articlepb.ArticleListInfo{
			ArticleId: article.ID,
			Title:     article.Title,
			AuthorId:  article.AuthorId,
			Summary:   article.Summary,
			Tag:       article.Tag,
			ImageUrl:  article.ImageURL,
			Status:    article.Status,
			CreatedAt: timestamppb.New(article.CreatedAt),
			UpdatedAt: timestamppb.New(article.UpdatedAt),
		})
	}

	return &articlepb.QueryArticleResponse{
		ArticleList: articles,
		Total:       int32(len(articleList)),
	}, nil
}

// 获取文章详情（fileId）
func (s *ArticleServer) GetArticleDetail(ctx context.Context, req *articlepb.GetArticleDetailRequest) (*articlepb.GetArticleDetailResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "GetArticleDetail",
	})
	// 获取请求参数
	articleId := req.GetArticleId()
	// 首先校验这个参数是否合法
	if err := validate.GetArticleDetail(req); err != nil {
		logger.Errorf("failed to get article detail with err(%s)", err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}
	// 查询文章
	article, err := s.DBEngine.GetArticleDetail(articleId)
	if err != nil {
		logger.Errorf("failed to get article detail with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article detail: %v", err)
	}

	return &articlepb.GetArticleDetailResponse{
		ArticleInfo: &articlepb.ArticleInfo{
			ArticleId:    article.ID,
			Title:        article.Title,
			Content:      article.Content,
			AuthorId:     article.AuthorId,
			Summary:      article.Summary,
			Tag:          article.Tag,
			ImageUrl:     article.ImageURL,
			Status:       article.Status,
			ViewCount:    int32(article.ViewCount),
			LikeCount:    int32(article.LikeCount),
			CommentCount: int32(article.CommentCount),
			CategoryId:   article.CategoryId,
			CreatedAt:    timestamppb.New(article.CreatedAt),
			UpdatedAt:    timestamppb.New(article.UpdatedAt),
		},
	}, nil
}
