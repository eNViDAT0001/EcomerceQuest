package app_file

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *mediaHandler) DeleteMedia() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		publicID := cc.Param("public_id")
		result, err := s.mediaUC.DeleteMedia(newCtx, publicID)

		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
