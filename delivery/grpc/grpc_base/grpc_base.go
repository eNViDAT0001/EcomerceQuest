package grpc_base

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/services/recommender"
	"github.com/google/wire"
)

var IteratorCollection = wire.NewSet(

	recommender.NewRecommenderService,
	NewGRPCCollection,
)

type GRPCCollection struct {
	RecommenderService recommender.GRPCService
}

func NewGRPCCollection(
	recommenderService recommender.GRPCService,

) *GRPCCollection {
	return &GRPCCollection{
		RecommenderService: recommenderService,
	}
}
