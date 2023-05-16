package app_accession

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/oidc"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *appAccessionHandler) GoogleSSO() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		url := oidc.GetConfig().GoogleOauthConfig.AuthCodeURL(oidc.RandomString)

		cc.Redirect(http.StatusTemporaryRedirect, url)
	}
}
