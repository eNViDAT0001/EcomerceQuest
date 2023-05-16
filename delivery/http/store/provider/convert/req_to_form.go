package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/provider/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"github.com/jinzhu/copier"
)

func ProviderCreateReqToProviderForm(inputEntity *io.CreateProviderReq) (ioSto.ProviderForm, error) {
	var result ioSto.ProviderForm
	err := copier.Copy(&result, &inputEntity)
	if err != nil {
		return result, err
	}
	return result, nil
}
func ProviderUpdateReqToProviderUpdateForm(inputEntity *io.UpdateProviderReq) (ioSto.ProviderUpdateForm, error) {
	var result ioSto.ProviderUpdateForm
	err := copier.Copy(&result, &inputEntity)
	if err != nil {
		return result, err
	}
	return result, nil
}
