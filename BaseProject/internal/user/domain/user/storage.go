package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"

	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

type Storage interface {
	GetUserDetailByID(ctx context.Context, ID uint) (*entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	GetUserList(ctx context.Context, input *paging.ParamsInput) ([]*entities.User, error)
	CreateUser(ctx context.Context, input *io.CreateUserInput) (userID uint, err error)
	UpdateUser(ctx context.Context, userID uint32, input *io.UpdateUserInput) error
	UpdatePassword(ctx context.Context, userID uint32, password string) error
	ComparePassword(ctx context.Context, userID uint32, password string) (io.UserPassword, error)
	DeleteUserByIDs(ctx context.Context, IDs []uint) error
	GetUserWithIdentify(ctx context.Context, input io.UserIdentify) (*io.UserPreview, error)
	CountList(ctx context.Context, input *paging.ParamsInput) (int64, error)
}
