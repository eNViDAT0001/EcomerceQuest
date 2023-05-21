package provider

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider"
)

type providerHandler struct {
	providerUC provider.UseCase
}

func NewProviderHandler(providerUC provider.UseCase) provider.HttpHandler {
	return &providerHandler{providerUC: providerUC}
}
