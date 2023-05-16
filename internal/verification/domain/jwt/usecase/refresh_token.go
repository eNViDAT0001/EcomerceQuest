package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"strconv"

	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase/io"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtUseCase) RefreshToken(ctx context.Context, refreshToken string) (*ioUC.JwtToken, error) {
	token, err := s.tokenSto.VerifyToken(nil, refreshToken)
	if err != nil {
		return nil, request.NewUnauthorizedError("refresh_token", refreshToken, "Invalid Refresh Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, request.NewUnauthorizedError("refresh_token", refreshToken, "Invalid Token")
	}

	userID, _ := strconv.Atoi(claims["user_id"].(string))
	user, err := s.userSto.GetUserDetailByID(ctx, uint(userID))

	if err != nil {
		return nil, request.NewUnauthorizedError("refresh_token", refreshToken, "Invalid Refresh Token")
	}

	baseToken := ioSto.GenerateTokenInput{
		UserID:   strconv.Itoa(int(user.ID)),
		Username: *user.Username,
		Salt:     *user.Salt,
	}

	minutes, _ := ioUC.GetMinute()
	days, _ := ioUC.GetDate()

	accessToken, err := s.tokenSto.GenerateToken(ctx, baseToken, minutes)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.tokenSto.GenerateToken(ctx, baseToken, days)
	if err != nil {
		return nil, err
	}

	output := ioUC.JwtToken{
		AccessToken:        accessToken.Token,
		AccessTokenExpiry:  accessToken.TokenExpiry,
		RefreshToken:       newRefreshToken.Token,
		RefreshTokenExpiry: newRefreshToken.TokenExpiry,
	}

	return &output, nil
}
