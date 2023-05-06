package user

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	GetAddressesWithUserID() func(*gin.Context)
	GetAddressDetailByID() func(*gin.Context)
	UpdateAddress() func(*gin.Context)
	DeleteAddressByIDs() func(*gin.Context)
	CreateAddress() func(*gin.Context)
}
