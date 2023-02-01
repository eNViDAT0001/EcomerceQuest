package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
	"github.com/jinzhu/copier"
)

func UserEntityToUserRes(userEntity *entities.User) (*io.UserRes, error) {
	var result io.UserRes
	err := copier.Copy(&result, userEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
