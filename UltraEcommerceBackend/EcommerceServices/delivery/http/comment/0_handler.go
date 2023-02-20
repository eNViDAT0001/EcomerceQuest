package comment

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/comment"
)

type commentHandler struct {
	commentUC comment.UseCase
}

func NewCommentHandler(commentUC comment.UseCase) comment.HttpHandler {
	return &commentHandler{commentUC: commentUC}
}
