package oidc

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	vp = wrap_viper.GetViper()
)

func getGoogleSSoConfig() *oauth2.Config {
	clientID := vp.GetString("GOOGLE_OAUTH2.CLIENT_ID")
	clientSecret := vp.GetString("GOOGLE_OAUTH2.CLIENT_SECRET")
	redirectURL := vp.GetString("GOOGLE_OAUTH2.REDIRECT_URL")

	scopes := []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}
	googleConfig := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}
	return googleConfig
}
