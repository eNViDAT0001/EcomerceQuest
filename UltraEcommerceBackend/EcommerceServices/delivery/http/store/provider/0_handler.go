package provider

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/provider"
	"github.com/gin-gonic/gin"
	"net/http"
)

type providerHandler struct {
	providerUC provider.UseCase
}

func (s providerHandler) GetProviderFullDetailByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		cc.Response(http.StatusInternalServerError, "Hold on bruhh! We still not implement this function yet")
	}
}

func NewProviderHandler(providerUC provider.UseCase) provider.HttpHandler {
	return &providerHandler{providerUC: providerUC}
}
