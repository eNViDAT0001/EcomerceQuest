package usecase

import "github.com/eNViDAT0001/Thesis/Backend/external_provider/recommender"

type RecommenderUseCase struct {
	recommendProvider recommender.RecommendProvider
}
