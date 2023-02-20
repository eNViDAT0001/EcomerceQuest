package convert

import (
	ioHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/user/io"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateReqToUpdateUserRepo[T ioHandler.UpdateUserReq](input *T) (*ioSto.UpdateUserInput, error) {
	var result ioSto.UpdateUserInput
	err := copier.Copy(&result, input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
