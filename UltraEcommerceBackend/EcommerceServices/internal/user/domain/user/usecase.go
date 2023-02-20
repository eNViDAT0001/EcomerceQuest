package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

type UseCase interface {
	GetUserDetailByID(ctx context.Context, ID uint) (*entities.User, error)
	GetUserList(ctx context.Context, input *paging.ParamsInput) ([]*entities.User, error)
	CreateUser(ctx context.Context, input *io.CreateUserInput) (userID uint, err error)
	UpdateUser(ctx context.Context, userID uint32, input *io.UpdateUserInput) error
	ComparePassword(ctx context.Context, userID uint32, password string) (io.UserPassword, error)
	SetPassword(ctx context.Context, userID uint32, password string, newPassword string) error
	DeleteUserByIDs(ctx context.Context, IDs []uint) error
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
}
