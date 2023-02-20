package io

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_viper"
	"strconv"
	"time"
)

var wViper = wrap_viper.GetViper()

type JwtToken struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpiry  int64  `json:"access_token_expiry"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpiry int64  `json:"refresh_token_expiry"`
}

func GetMinute() (time.Time, error) {
	minute, err := strconv.Atoi(wViper.GetString("ACCESS_TOKEN_EXPIRE_IN_MINUTE"))
	//if minute < 30 {
	//	minute = 30 // default 30 minutes if not defined
	//}
	expiresAt := time.Now().Add(time.Minute * time.Duration(minute))
	return expiresAt, err
}
func GetDate() (time.Time, error) {
	days, err := strconv.Atoi(wViper.GetString("REFRESH_TOKEN_EXPIRE_IN_DAYS"))
	if days < 1 {
		days = 1 // default 2 days if not defined
	}
	expiresAt := time.Now().Add(24 * time.Hour * time.Duration(days))
	return expiresAt, err
}
