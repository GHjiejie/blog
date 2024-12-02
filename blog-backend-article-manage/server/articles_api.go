package server

import (
	"context"
	"strings"

	articlepb "articleManage/pb/articleManage"
	"articleManage/pkg/auth"
	"articleManage/pkg/db"
	"articleManage/validate"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	if imageUrl == "" {
		imageUrl = "https://images.pexels.com/photos/3423896/pexels-photo-3423896.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
	}
	// logger.Infof("title: %s, content: %s, authorId: %d, summary: %s, tag: %s, imageUrl: %s, status: %s", title, content, authorId, summary, tag, imageUrl, status)
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
	_, err := s.DBEngine.CreateArticle(articleInfo)
	if err != nil {
		logger.Errorf("failed to create article with err(%s)", err.Error())
		return nil, err
	}
	// logger.Infof("article created: %v", article)

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

	// 根据id查询文章的详细信息
	article, err := s.DBEngine.GetArticleDetail(articleId)
	if err != nil {
		logger.Errorf("failed to get article detail with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article detail: %v", err)
	}

	// 创建一个更新的结构体，只更新有变动的字段
	updateFields := db.Article{
		ID:           article.ID,
		Title:        article.Title,        // 保留原标题
		Content:      article.Content,      // 保留原内容
		AuthorId:     article.AuthorId,     // 保留原作者
		Summary:      article.Summary,      // 保留原摘要
		Tag:          article.Tag,          // 保留原标签
		ViewCount:    article.ViewCount,    // 保留原浏览次数
		LikeCount:    article.LikeCount,    // 保留原点赞次数
		CommentCount: article.CommentCount, // 保留原评论次数
		Comments:     article.Comments,     // 保留原评论
		CategoryId:   article.CategoryId,   // 保留原分类ID
		ImageURL:     article.ImageURL,     // 保留原图片
		Status:       article.Status,       // 保留原状态
		CreatedAt:    article.CreatedAt,    // 保留原创建时间
		UpdatedAt:    timestamppb.Now().AsTime(),
		DeletedAt:    article.DeletedAt, // 保留原删除时间
	}

	// 检查哪些字段需要更新
	if title := req.GetTitle(); title != "" {
		updateFields.Title = title
	}
	if content := req.GetContent(); content != "" {
		updateFields.Content = content
	}
	if summary := req.GetSummary(); summary != "" {
		updateFields.Summary = summary
	}
	if tag := req.GetTag(); tag != "" {
		updateFields.Tag = tag
	}

	if imageUrl := req.GetImageUrl(); imageUrl != "" {
		updateFields.ImageURL = imageUrl
	}
	if authorId := req.GetAuthorId(); authorId != 0 {
		updateFields.AuthorId = authorId
	}
	logger.Infof("title: %s, content: %s, authorId: %d, summary: %s, imageUrl: %s, status: %s", updateFields.Title, updateFields.Content, updateFields.AuthorId, updateFields.Summary, updateFields.ImageURL, updateFields.Status)

	// 更新文章
	if err := s.DBEngine.UpdateArticle(updateFields); err != nil {
		logger.Errorf("failed to update article with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}

	return &articlepb.UpdateArticleResponse{
		Message: "article updated successfully",
	}, nil
}

// func (s *ArticleServer) UpdateArticle(ctx context.Context, req *articlepb.UpdateArticleRequest) (*articlepb.UpdateArticleResponse, error) {
// 	logger := log.WithFields(log.Fields{
// 		"api": "UpdateArticle",
// 	})
// 	// 获取请求参数
// 	articleId := req.GetArticleId()
// 	title := req.GetTitle()
// 	content := req.GetContent()
// 	summary := req.GetSummary()
// 	tag := req.GetTag()
// 	fileStatus := req.GetStatus()
// 	imageUrl := req.GetImageUrl()
// 	authorId := req.GetAuthorId()

// 	// 根据id去查询这个文章的消息
// 	article, err := s.DBEngine.GetArticleDetail(articleId)
// 	if err != nil {
// 		logger.Errorf("failed to get article detail with err(%s)", err.Error())
// 		return nil, status.Errorf(codes.Internal, "failed to get article detail: %v", err)
// 	}
// 	// 首先校验这个参数是否合法(这个太多了,不处理了,前端处理就好了)
// 	fileInfo := db.Article{
// 		ID:        articleId,
// 		Title:     title,
// 		Content:   content,
// 		AuthorId:  authorId,
// 		Summary:   summary,
// 		Tag:       tag,
// 		ImageURL:  imageUrl,
// 		Status:    fileStatus,
// 		CreatedAt: article.CreatedAt,
// 		UpdatedAt: timestamppb.Now().AsTime(),
// 	}
// 	// 更新文章
// 	if err := s.DBEngine.UpdateArticle(fileInfo); err != nil {
// 		logger.Errorf("failed to update article with err(%s)", err.Error())
// 		return nil, status.Errorf(codes.Internal, "failed to update article: %v", err)
// 	}

// 	return &articlepb.UpdateArticleResponse{
// 		Message: "article updated successfully",
// 	}, nil
// }

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
	// articleCount, err := s.DBEngine.GetArticleCount()
	// if err != nil {
	// 	logger.Errorf("failed to get article count with err(%s)", err.Error())
	// 	return nil, err
	// }
	// 判断文章ID是否在合法范围内
	// if articleId > articleCount {
	// 	logger.Errorf("articleId(%d) is out of range", articleId)
	// 	return nil, status.Errorf(codes.InvalidArgument, "articleId is out of range")
	// }
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
	// 获取请求头里面的auth
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Errorf("Failed to get metadata from context")
		return nil, status.Error(codes.Internal, "failed to get metadata from context")
	}
	authInfo := md.Get("authorization")
	if len(authInfo) == 0 {
		logger.Errorf("Failed to get authorization from context")
		return nil, status.Error(codes.Internal, "failed to get authorization from context")
	}
	// logger.Infof("Received authorization: %s", authInfo[0])

	// 解析token
	claims, err := auth.ParseUserToken(authInfo[0])
	if err != nil {
		logger.Errorf("Failed to parse token: %v", err)
		return nil, status.Error(codes.Internal, "failed to parse token")
	}
	logger.Infof("Received claims: %v", claims)

	// 判断是不是管理员，如果是管理员，就返回所有文章，如果不是管理员，就返回自己的文章
	var articleList []db.Article
	if claims.Role == "ADMIN" {
		articleList, err = s.DBEngine.GetArticleList(page, pageSize)
	} else {
		articleList, err = s.DBEngine.GetArticleListByAuthor(int32(claims.UserId), page, pageSize)
	}
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

// 审核文章
func (s *ArticleServer) ReviewArticle(ctx context.Context, req *articlepb.ReviewArticleRequest) (*articlepb.ReviewArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "ReviewArticle",
	})
	// 获取请求参数
	articleId := req.GetArticleId()
	articleStatus := req.GetStatus()

	// 查询文章
	article, err := s.DBEngine.GetArticleDetail(articleId)
	if err != nil {
		logger.Errorf("failed to get article detail with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article detail: %v", err)
	}
	// 更新文章状态
	article.Status = articleStatus
	if err := s.DBEngine.UpdateArticle(article); err != nil {
		logger.Errorf("failed to update article with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}

	return &articlepb.ReviewArticleResponse{
		Message: "article reviewed successfully",
	}, nil
}
