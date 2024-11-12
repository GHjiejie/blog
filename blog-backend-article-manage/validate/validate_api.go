package validate

import (
	articlepb "articleManage/pb/articleManage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DeleteArticle(req *articlepb.DeleteArticleRequest) error {
	articleId := req.GetArticleId()
	if articleId == 0 {
		return status.Errorf(codes.InvalidArgument, "articleId is required")
	}
	return nil
}
func GetArticleDetail(req *articlepb.GetArticleDetailRequest) error {
	articleId := req.GetArticleId()
	if articleId == 0 {
		return status.Errorf(codes.InvalidArgument, "articleId is required")
	}
	return nil
}
