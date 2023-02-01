package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"github.com/jinzhu/copier"
)

func ProviderFormToEntityProvider(input io.ProviderForm) (entities.Provider, error) {
	var result entities.Provider
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
