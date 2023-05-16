package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
	"github.com/jinzhu/copier"
)

func ArrUserEntityToArrUserRes(userEntity []*entities.User) ([]*io.UserRes, error) {
	result := make([]*io.UserRes, 0)

	err := copier.Copy(&result, userEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}
