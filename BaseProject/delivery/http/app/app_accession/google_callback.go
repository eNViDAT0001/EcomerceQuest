package app_accession

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/app/app_accession/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/oidc"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
	ioJwtSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func (a *appAccessionHandler) CallbackGoogleSSO() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()
		state, _ := cc.GetQuery("state")
		code, _ := cc.GetQuery("code")
		googleProfile, err := getGoogleUserData(state, code)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		if !googleProfile.Verified {
			cc.ResponseError(errors.New("email not verified"))
			return
		}

		user, err := a.userUC.GetUserByEmail(newCtx, googleProfile.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			cc.ResponseError(err)
			return
		}

		var userID uint
		var salt string
		var username string

		if err == gorm.ErrRecordNotFound {
			newUser := &io2.CreateUserInput{
				Username: strconv.FormatInt(time.Now().Unix(), 10),
				Password: time.Now().String(),
				Salt:     time.Now().String(),
				Name:     fmt.Sprintf("%s %s", googleProfile.GivenName, googleProfile.FamilyName),
				Birthday: "2000/01/01",
				Gender:   false,
				Email:    googleProfile.Email,
				Phone:    "9999999999",
				Type:     entities.Buyer,
			}

			userID, err = a.userUC.CreateUser(newCtx, newUser)
			if err != nil {
				cc.ResponseError(err)
				return
			}

			username = newUser.Username
			salt = newUser.Salt
		}

		if userID == 0 {
			userID = user.ID
			salt = *user.Salt
			username = *user.Username
		}

		tokenForm := ioJwtSto.GenerateTokenInput{
			UserID:   strconv.Itoa(int(userID)),
			Username: username,
			Salt:     salt,
		}
		token, err := a.jwtUC.GenerateToken(tokenForm)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"UserID": user.ID,
			"Token":  token,
		}

		cc.Ok(result)
	}
}
func getGoogleUserData(state, code string) (io.GoogleProfile, error) {
	var result io.GoogleProfile

	if state != oidc.RandomString {
		return result, errors.New("invalid oauth state")
	}
	token, err := oidc.GetConfig().GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return result, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
