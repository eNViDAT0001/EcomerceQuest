package admin

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	entities4 "github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	entities3 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
	"github.com/gin-gonic/gin"
)

// Don't ask why, I am so lazy to write the same thing again and again
// I will make a code generator for this later

func (s *adminHandler) GetAdminReport() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)

		db := wrap_gorm.GetDB()

		var totalUser int64
		err := db.Model(entities.User{}).Count(&totalUser).Error
		if err != nil {
			cc.ResponseError(err)
			return
		}
		var totalProvider int64
		err = db.Model(entities2.Provider{}).Count(&totalProvider).Error
		if err != nil {
			cc.ResponseError(err)
			return
		}
		var totalProduct int64
		err = db.Model(entities3.Product{}).Count(&totalProduct).Error
		if err != nil {
			cc.ResponseError(err)
			return
		}
		var totalOrder int64
		err = db.Model(entities4.Order{}).Count(&totalOrder).Error
		if err != nil {
			cc.ResponseError(err)
			return
		}
		var totalRevenue float64
		err = db.Model(entities4.Order{}).
			Select("SUM(total)").
			Where("status = ?", entities4.DeliveredOrder).
			Scan(&totalRevenue).Error
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"total_user":     totalUser,
			"total_provider": totalProvider,
			"total_product":  totalProduct,
			"total_order":    totalOrder,
			"total_revenue":  totalRevenue,
		}
		cc.Ok(result)
	}
}
