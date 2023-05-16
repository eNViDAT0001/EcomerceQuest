package oidc

import (
	"golang.org/x/oauth2"
)

const RandomString = "SomeRandomStringAlgorithmForMoreSecurity"

type ssoConfig struct {
	GoogleOauthConfig *oauth2.Config
}

var config *ssoConfig

func GetConfig() *ssoConfig {

	if config != nil {
		return config
	}

	config = &ssoConfig{
		getGoogleSSoConfig(),
	}

	return config
}
