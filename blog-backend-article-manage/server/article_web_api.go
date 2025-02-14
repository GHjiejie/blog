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
	var articles []*articlepb.ArticleListInfo
	for _, article := range articleList {
		articles = append(articles, &articlepb.ArticleListInfo{
			ArticleId:    article.ID,
			Title:        article.Title,
			AuthorId:     article.AuthorId,
			Summary:      article.Summary,
			Tag:          article.Tag,
			ImageUrl:     article.ImageURL,
			Status:       article.Status,
			ViewCount:    int32(article.ViewCount),
			LikeCount:    int32(article.LikeCount),
			CommentCount: int32(article.CommentCount),
			CreatedAt:    timestamppb.New(article.CreatedAt),
			UpdatedAt:    timestamppb.New(article.UpdatedAt),
		})
	}

	return &articlepb.GetPublishedArticleListResponse{
		ArticleList: articles,
		Total:       int32(articleCount),
	}, nil
}

// 文章点赞增加
func (s *ArticleServer) LikeArticle(ctx context.Context, req *articlepb.LikeArticleRequest) (*articlepb.LikeArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "LikeArticle",
	})
	logger.Infof("LikeArticle request with articleId(%d) and userId(%d)", req.GetArticleId(), req.GetUserId())
	// 接下来要做的就是在数据库中创建一条点赞记录，同时更新文章的点赞次数
	articleId := req.GetArticleId()
	userId := req.GetUserId()

	// 查询文章是否存在()
	// article, err := s.DBEngine.GetArticleDetail(articleId)
	// if err != nil {
	// 	logger.Errorf("failed to get article by id(%d) with err(%s)", articleId, err.Error())
	// 	return nil, err
	// }
	// if article == nil {
	// 	logger.Errorf("article(%d) not found", articleId)
	// 	return nil, status.Errorf(codes.NotFound, "article(%d) not found", articleId)
	// }

	// 查询用户是否已经点赞
	_, err := s.DBEngine.GetArticleLike(articleId, userId)
	if err != nil {
		// 那么我们就需要在数据库中创建一条点赞记录，同时更新文章的点赞次数
		// 创建点赞记录
		err = s.DBEngine.CreateArticleLike(articleId, userId)
		if err != nil {
			logger.Errorf("failed to create article like with err(%s)", err.Error())
			return nil, err
		}

		// 更新文章点赞次数
		err = s.DBEngine.UpdateArticleLikeCount(articleId, 1)
		if err != nil {
			logger.Errorf("failed to update article like count with err(%s)", err.Error())
			return nil, err
		}
	} else {
		logger.Infof("user(%d) already like article(%d)", userId, articleId)
		return &articlepb.LikeArticleResponse{
			Message: "user already like article",
		}, nil
	}

	return &articlepb.LikeArticleResponse{
		Message: "like article success",
	}, nil
}

//	func (UnimplementedArticleManageServiceServer) CancelLikeArticle(context.Context, *LikeArticleRequest) (*LikeArticleResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method CancelLikeArticle not implemented")
//	}
//
// 文章点赞取消
func (s *ArticleServer) CancelLikeArticle(ctx context.Context, req *articlepb.CancelLikeArticleRequest) (*articlepb.CancelLikeArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "UnLikeArticle",
	})
	logger.Infof("UnLikeArticle request with articleId(%d) and userId(%d)", req.GetArticleId(), req.GetUserId())
	// 接下来要做的就是在数据库中删除一条点赞记录，同时更新文章的点赞次数
	articleId := req.GetArticleId()
	userId := req.GetUserId()

	// 查询用户是否已经点赞
	like, err := s.DBEngine.GetArticleLike(articleId, userId)
	if err != nil {
		logger.Errorf("failed to get article like with err(%s)", err.Error())
		return nil, err
	}
	if like.ID == 0 {
		logger.Infof("user(%d) not like article(%d)", userId, articleId)
		return &articlepb.CancelLikeArticleResponse{
			Message: "user not like article",
		}, nil
	}

	// 删除点赞记录
	err = s.DBEngine.DeleteArticleLike(articleId, userId)
	if err != nil {
		logger.Errorf("failed to delete article like with err(%s)", err.Error())
		return nil, err
	}

	// 更新文章点赞次数
	err = s.DBEngine.UpdateArticleLikeCount(articleId, -1)
	if err != nil {
		logger.Errorf("failed to update article like count with err(%s)", err.Error())
		return nil, err
	}

	return &articlepb.CancelLikeArticleResponse{
		Message: "cancel like article success",
	}, nil

}

// 文章浏览数增加
func (s *ArticleServer) ViewArticle(ctx context.Context, req *articlepb.ViewArticleRequest) (*articlepb.ViewArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "ViewArticle",
	})
	logger.Infof("ViewArticle request with articleId(%d)", req.GetArticleId())
	// 接下来要做的就是在数据库中更新文章的浏览次数
	articleId := req.GetArticleId()

	// 更新文章浏览次数
	err := s.DBEngine.UpdateArticleViewCount(articleId, 1)
	if err != nil {
		logger.Errorf("failed to update article view count with err(%s)", err.Error())
		return nil, err
	}

	return &articlepb.ViewArticleResponse{
		Message: "view article success",
	}, nil
}

// 查询用户是否已经点赞文章

func (s *ArticleServer) QueryUserLikeArticle(ctx context.Context, req *articlepb.QueryUserLikeArticleRequest) (*articlepb.QueryUserLikeArticleResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "QueryUserLikeArticle",
	})
	logger.Infof("QueryUserLikeArticle request with articleId(%d) and userId(%d)", req.GetArticleId(), req.GetUserId())
	// 接下来要做的就是在数据库中查询用户是否已经点赞文章
	articleId := req.GetArticleId()
	userId := req.GetUserId()

	// 查询用户是否已经点赞
	_, err := s.DBEngine.GetArticleLike(articleId, userId)
	if err != nil {
		logger.Errorf("failed to get article like with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get article like with err(%s)", err.Error())
	}
	return &articlepb.QueryUserLikeArticleResponse{
		Message: "user already like article",
	}, nil

	// return nil, status.Errorf(codes.Unimplemented, "method QueryUserLikeArticle not implemented")
}
