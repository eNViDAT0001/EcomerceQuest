package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

func (c categoryStorage) GetCategoryChildrenTreeWithCategoryID(ctx context.Context, categoryID uint) ([]entities.Category, error) {
	result := make([]entities.Category, 0)

	db := wrap_gorm.GetDB()

	err := db.Raw("WITH recursive cte (id, name, category_parent_id, image_path)"+
		" AS ( SELECT id, name, category_parent_id, image_path FROM Category"+
		" WHERE category_parent_id = ? UNION ALL SELECT p.id, p.name, p.category_parent_id, p.image_path"+
		" FROM Category p INNER JOIN cte ON p.category_parent_id = cte.id ) SELECT * FROM cte", categoryID).
		Find(&result).
		Error

	if err != nil {
		return result, err
	}

	if len(result) == 0 {
		return result, gorm.ErrRecordNotFound
	}

	return result, nil
}
func GetCategoryChildrenTree(roof *ioSto.CategoryChildrenTree, node ioSto.CategoryChildrenTree) {
	if node.CategoryParentID == roof.ID {
		roof.CategoryChildren = append(roof.CategoryChildren, node)
		return
	}

	for i := range roof.CategoryChildren {
		if roof.CategoryChildren[i].ID == node.CategoryParentID {
			roof.CategoryChildren[i].CategoryChildren = append(roof.CategoryChildren[i].CategoryChildren, node)
			return
		}
		GetCategoryChildrenTree(&roof.CategoryChildren[i], node)
	}

}
