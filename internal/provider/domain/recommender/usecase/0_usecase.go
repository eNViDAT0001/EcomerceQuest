package usecase

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/recommender"
)

type RecommenderUseCase struct {
	recommendProvider recommender.RecommendProvider
}
