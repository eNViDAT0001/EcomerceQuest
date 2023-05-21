package admin

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/admin"
)

type adminHandler struct {
}

func NewAdminHandler() admin.HttpHandler {
	return &adminHandler{}
}
