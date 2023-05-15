package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s bannerHandler) GetBannerByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		bannerID, err := strconv.Atoi(cc.Param("banner_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.bannerUC.GetBannerByID(newCtx, uint(bannerID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}

func (s bannerHandler) GetBannerDetailByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		bannerID, err := strconv.Atoi(cc.Param("banner_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.bannerUC.GetBannerDetailByID(newCtx, uint(bannerID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
