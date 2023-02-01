package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Backend/internal/app/domain/app_accession"
	userPKG "github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase/io"
)

var wViper = wrap_viper.GetViper()

type appAccessionUseCase struct {
	userSto  userPKG.Storage
	tokenSto jwt.Storage
}

func (u *appAccessionUseCase) LoginByGoogle(ctx context.Context, client io.Client) (*ioUC.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (u *appAccessionUseCase) LoginByFacebook(ctx context.Context, client io.Client) (*ioUC.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func NewAppAccessionUseCase(
	userSto userPKG.Storage,
	tokenSto jwt.Storage,

) app_accession.UseCase {
	return &appAccessionUseCase{
		userSto:  userSto,
		tokenSto: tokenSto,
	}
}
