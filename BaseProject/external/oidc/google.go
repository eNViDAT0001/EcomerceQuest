package oidc

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getGoogleSSoConfig() *oauth2.Config {
	googleConfig := &oauth2.Config{
		ClientID:     "775502336982-tpq7oppduoni9f8reu1o9sbp8ckvus6e.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8082/api/v1/app/login/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}
	return googleConfig
}
