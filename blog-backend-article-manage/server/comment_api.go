package server

import (
	articlepb "articleManage/pb/articleManage"
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 获取文章评论列表
func (s *ArticleServer) GetArticleCommentList(ctx context.Context, req *articlepb.GetArticleCommentListRequest) (*articlepb.GetArticleCommentListResponse, error) {

	logger := log.WithFields(log.Fields{
		"api": "GetArticleCommentList",
	})

	// 获取请求参数
	articleId := req.GetArticleId()
	page := req.GetPage()
	pageSize := req.GetPageSize()

	// 参数校验
	if articleId <= 0 || page <= 0 || pageSize <= 0 {
		logger.Errorf("articleId(%d) or page(%d) or pageSize(%d) is invalid", articleId, page, pageSize)
		return nil, status.Errorf(codes.InvalidArgument, "articleId or page or pageSize is invalid")

	}

	// 获取文章评论列表
	commentList, err := s.DBEngine.GetArticleCommentList(articleId, page, pageSize)
	if err != nil {
		logger.Errorf("failed to get article comment list with err(%s)", err.Error())
		return nil, err
	}

	// 获取文章评论总数
	total, err := s.DBEngine.GetArticleCommentCount(articleId)
	if err != nil {
		logger.Errorf("failed to get article comment count with err(%s)", err.Error())
		return nil, err
	}

	var comments []*articlepb.ArticleCommentInfo
	for _, comment := range commentList {
		comments = append(comments, &articlepb.ArticleCommentInfo{
			Id:        comment.ID,
			ArticleId: comment.ArticleId,
			UserId:    comment.UserId,
			Content:   comment.Content,
			Status:    comment.ReviewStatus,
			LikeCount: int32(comment.LikeCount),
			CreatedAt: timestamppb.New(comment.CreatedAt),
			UpdatedAt: timestamppb.New(comment.UpdatedAt),
		})
	}
	return &articlepb.GetArticleCommentListResponse{
		CommentList: comments,
		Total:       int32(total),
	}, nil

}

// 发布文章评论
func (s *ArticleServer) PublishArticleComment(ctx context.Context, req *articlepb.PublishArticleCommentRequest) (*articlepb.PublishArticleCommentResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "PublishArticleComment",
	})

	// 获取请求参数
	articleId := req.GetArticleId()
	userId := req.GetUserId()
	content := req.GetContent()

	// 参数校验
	if articleId <= 0 || userId <= 0 || content == "" {
		logger.Errorf("articleId(%d) or userId(%d) or content(%s) is invalid", articleId, userId, content)
		return nil, status.Errorf(codes.InvalidArgument, "articleId or userId or content is invalid")
	}

	// 发布文章评论
	_, err := s.DBEngine.CreateArticleComment(articleId, userId, content)
	if err != nil {
		logger.Errorf("failed to create article comment with err(%s)", err.Error())
		return nil, err
	}

	var addCommentCount int64 = 1
	// 更新文章评论数
	if err := s.DBEngine.UpdateArticleCommentCount(articleId, addCommentCount); err != nil {
		logger.Errorf("failed to update article comment count with err(%s)", err.Error())
		return nil, err
	}

	return &articlepb.PublishArticleCommentResponse{
		Message: "comment published successfully",
	}, nil
}

// 删除文章评论
func (s *ArticleServer) DeleteArticleComment(ctx context.Context, req *articlepb.DeleteArticleCommentRequest) (*articlepb.DeleteArticleCommentResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "DeleteArticleComment",
	})

	// 获取请求参数
	commentId := req.GetCommentId()
	articleId := req.GetArticleId()

	// 参数校验
	if commentId <= 0 {
		logger.Errorf("commentId(%d) is invalid", commentId)
		return nil, status.Errorf(codes.InvalidArgument, "commentId is invalid")
	}

	// 删除文章评论
	err := s.DBEngine.DeleteArticleComment(commentId)
	if err != nil {
		logger.Errorf("failed to delete article comment with err(%s)", err.Error())
		return nil, err
	}

	// 删除该评论的点赞记录
	err = s.DBEngine.DeleteCommentLike(articleId, commentId)
	if err != nil {
		logger.Errorf("failed to delete comment like with err(%s)", err.Error())
		return nil, err
	}

	// 更新文章评论数
	var addCommentCount int64 = -1
	err = s.DBEngine.UpdateArticleCommentCount(articleId, addCommentCount)
	if err != nil {
		logger.Errorf("failed to update article comment count with err(%s)", err.Error())
		return nil, err
	}

	return &articlepb.DeleteArticleCommentResponse{
		Message: "delete article comment successfully",
	}, nil
}

// func (UnimplementedArticleManageServiceServer) GetArticleCommentDetail(context.Context, *GetArticleCommentDetailRequest) (*GetArticleCommentDetailResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetArticleCommentDetail not implemented")
// }

// 点赞文章评论
func (s *ArticleServer) LikeArticleComment(ctx context.Context, req *articlepb.LikeArticleCommentRequest) (*articlepb.LikeArticleCommentResponse, error) {

	logger := log.WithFields(log.Fields{
		"api": "LikeArticleComment",
	})

	// 获取请求参数
	commentId := req.GetCommentId()
	userId := req.GetUserId()
	articleId := req.GetArticleId()

	// 参数校验(这里后续有时间再优化吧)
	if commentId <= 0 || userId <= 0 {
		logger.Errorf("commentId(%d) or userId(%d) is invalid", commentId, userId)
		return nil, status.Errorf(codes.InvalidArgument, "commentId or userId is invalid")
	}

	// 点赞文章评论
	err := s.DBEngine.UpdateCommentLikeCount(commentId)
	if err != nil {
		logger.Errorf("failed to like article comment with err(%s)", err.Error())
		return nil, err
	}
	// 更新点赞记录
	err = s.DBEngine.UpDateCommentLike(commentId, userId, articleId)
	if err != nil {
		logger.Errorf("failed to update comment like with err(%s)", err.Error())
		return nil, err
	}

	return &articlepb.LikeArticleCommentResponse{
		Message: "like article comment successfully",
	}, nil
}
