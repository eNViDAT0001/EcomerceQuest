package client

import (
	client "github.com/eNViDAT0001/Thesis/Ecommerce/internal/verification/domain/client"
	"github.com/gin-gonic/gin"
)

type clientHandler struct {
	clientUC client.UseCase
}

func (c clientHandler) CreateClient() func(*gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c clientHandler) UpdateClient() func(*gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewClientHandler(clientUC client.UseCase) client.HttpHandler {
	return &clientHandler{clientUC: clientUC}
}
