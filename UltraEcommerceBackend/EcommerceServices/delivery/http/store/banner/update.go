package banner

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/banner/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/banner/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s bannerHandler) UpdateBanner() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.BannerUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		bannerID, err := strconv.Atoi(cc.Param("banner_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		inputSto, err := convert.UpdateReqToUpdateBannerInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.bannerUC.UpdateBanner(newCtx, uint(bannerID), inputSto, input.ProductIDsIN, input.ProductIDsOUT)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"BannerID": bannerID,
		}

		cc.Ok(result)
	}
}
