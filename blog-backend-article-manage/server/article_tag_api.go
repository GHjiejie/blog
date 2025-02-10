package server

import (
	articlepb "articleManage/pb/articleManage"
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ArticleServer) GetArticleTagList(ctx context.Context, req *articlepb.GetArticleTagListRequest) (*articlepb.GetArticleTagListResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "GetArticleTagList",
	})
	page := req.GetPage()
	pageSize := req.GetPageSize()
	if page <= 0 || pageSize <= 0 {
		logger.Errorf("invalid page(%d) or pageSize(%d)", page, pageSize)
		return nil, status.Errorf(codes.InvalidArgument, "invalid page or pageSize")
	}

	// 获取tag列表
	tags, err := s.DBEngine.GetTagList(page, pageSize)
	if err != nil {
		logger.Errorf("failed to get tag list with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get tag list")
	}
	logger.Info("get tag list success, tags:", tags)

	var tagList []*articlepb.ArticleTagInfo
	for _, tag := range tags {
		tagList = append(tagList, &articlepb.ArticleTagInfo{
			TagId:         tag.ID,
			TagName:       tag.Name,
			TagCategoryId: tag.CategoryId,
			CreatedAt:     timestamppb.New(tag.CreatedAt),
			UpdatedAt:     timestamppb.New(tag.UpdatedAt),
		})
	}

	// 获取tag总数
	total, err := s.DBEngine.GetTagCount()

	return &articlepb.GetArticleTagListResponse{
		Total:   int32(total),
		TagList: tagList,
	}, nil
}

// 新增tag
func (s *ArticleServer) AddArticleTag(ctx context.Context, req *articlepb.AddArticleTagRequest) (*articlepb.AddArticleTagResponse, error) {

	logger := log.WithFields(log.Fields{
		"api": "AddArticleTag",
	})
	// 从请求中获取tag
	name := req.GetTagName()
	categoryId := req.GetTagCategoryId()

	// 查看是否已经存在该tag
	err := s.DBEngine.GetTagByName(name)
	if err == nil {
		logger.Errorf("tag(%s) already exists", name)
		return nil, status.Errorf(codes.AlreadyExists, "tag already exists")

	}

	// 创建tag
	err = s.DBEngine.CreateTag(name, categoryId)
	if err != nil {
		logger.Errorf("failed to create tag with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create tag")
	}

	return &articlepb.AddArticleTagResponse{
		Message: "tag add success",
	}, nil
}

func (s *ArticleServer) DeleteArticleTag(ctx context.Context, req *articlepb.DeleteArticleTagRequest) (*articlepb.DeleteArticleTagResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "DeleteArticleTag",
	})
	tagId := req.GetTagId()

	// 判断tag是否存在
	_, err := s.DBEngine.GetTagById(tagId)
	if err != nil {
		logger.Errorf("tag(%d) not exists", tagId)
		return nil, status.Errorf(codes.NotFound, "tag not exists")
	}

	// 删除tag
	err = s.DBEngine.DeleteTag(tagId)
	if err != nil {
		logger.Errorf("failed to delete tag with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to delete tag")
	}

	return &articlepb.DeleteArticleTagResponse{
		Message: "tag delete success ",
	}, nil

}
