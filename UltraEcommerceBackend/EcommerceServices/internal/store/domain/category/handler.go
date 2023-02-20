package category

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	CreateCategory() func(*gin.Context)
	UpdateCategory() func(*gin.Context)
	ListCategories() func(*gin.Context)
	GetCategoryChildrenTreeWithCategoryID() func(*gin.Context)
	GetCategoryParentsTreeWithCategoryID() func(*gin.Context)
	GetCategoryRoofList() func(*gin.Context)
	GetCategoryTree() func(*gin.Context)
}
