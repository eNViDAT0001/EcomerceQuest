package order

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *orderHandler) GetOrderReportByProviderID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		//newCtx := context.Background()
		//
		//providerID, err := strconv.Atoi(cc.Param("provider_id"))
		//if err != nil {
		//	cc.BadRequest(err)
		//	return
		//}

		//result, total, err := s.orderUC.GetOrderReportByProviderID(newCtx, uint(providerID))
		//if err != nil {
		//	if err == gorm.ErrRecordNotFound {
		//		cc.NoContent()
		//		return
		//	}
		//	cc.ResponseError(err)
		//	return
		//}

		cc.Ok("result")
	}
}
