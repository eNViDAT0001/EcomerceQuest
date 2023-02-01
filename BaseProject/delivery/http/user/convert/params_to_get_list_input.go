package convert

import (
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/jinzhu/copier"
)

func ParamsToGetListInput(input *ioHandler.GetUserListParams) (*paging.ParamsInput, error) {
	var result paging.ParamsInput
	err := copier.Copy(&result, input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
