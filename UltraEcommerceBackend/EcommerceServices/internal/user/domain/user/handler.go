package user

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetUserDetailByID() func(*gin.Context)
	GetUserList() func(*gin.Context)
	CreateUser() func(*gin.Context)
	AdminUpdateUser() func(*gin.Context)
	UpdateUserIdentity() func(*gin.Context)
	UpdateUserInfo() func(*gin.Context)
	SetPassword() func(*gin.Context)
	DeleteUserByID() func(*gin.Context)
	DeleteUserByIDs() func(*gin.Context)
}
