package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s orderStorage) GetOrdersReport(ctx context.Context, filter io.OrderReportFilter) (report io.OrderReport, err error) {
	db := wrap_gorm.GetDB()
	err = db.Model(entities.Order{}).Select("sdf").Find(&report).Error
	if err != nil {
		return report, err
	}
	return report, nil
}
