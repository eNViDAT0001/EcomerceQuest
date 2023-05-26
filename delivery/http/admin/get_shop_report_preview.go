package admin

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	entities4 "github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	entities3 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *adminHandler) GetShopReportPreview() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)

		db := wrap_gorm.GetDB()

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		providers := make([]entities2.Provider, 0)
		err = db.Model(entities2.Provider{}).Where("user_id = ?", userID).Find(&providers).Error
		if err != nil {
			cc.ResponseError(err)
			return
		}
		providerIDs := make([]uint, 0)
		for _, provider := range providers {
			providerIDs = append(providerIDs, provider.ID)
		}
		var totalProduct int64 = 0
		var totalOrder int64 = 0
		var totalRevenue float64 = 0

		if len(providerIDs) > 0 {
			err = db.Model(entities3.Product{}).Where("provider_id IN ?", providerIDs).Count(&totalProduct).Error
			if err != nil {
				cc.ResponseError(err)
				return
			}

			err = db.Model(entities4.Order{}).Where("provider_id IN ?", providerIDs).Count(&totalOrder).Error
			if err != nil {
				cc.ResponseError(err)
				return
			}

			err = db.Model(entities4.Order{}).
				Select("IF(SUM(total) IS NULL, 0, SUM(total)) as total").
				Where("status = ?", entities4.DeliveredOrder).
				Scan(&totalRevenue).Error
		}

		result := map[string]interface{}{
			"total_product": totalProduct,
			"total_order":   totalOrder,
			"total_revenue": totalRevenue,
		}

		cc.Ok(result)
	}
}
