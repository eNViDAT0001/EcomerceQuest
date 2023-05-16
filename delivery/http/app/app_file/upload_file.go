package app_file

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/app/app_file/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_cloudinary"
	"github.com/gin-gonic/gin"
)

func (s *mediaHandler) UploadMedia() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var form io.FileForm
		err := c.ShouldBind(&form)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result, err := s.mediaUC.UploadMedia(newCtx, form.File, wrap_cloudinary.Product)

		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
