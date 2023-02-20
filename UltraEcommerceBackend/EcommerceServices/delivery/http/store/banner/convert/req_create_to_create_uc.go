package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/banner/io"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/banner/storage/io"
	"github.com/jinzhu/copier"
)

func CreateReqToCreateBannerInput(input *ioHttpHandler.BannerCreateReq) (ioSto.BannerCreateForm, error) {
	var result ioSto.BannerCreateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
