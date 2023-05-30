//go:build wireinject
// +build wireinject

package grpc_base

import "github.com/google/wire"

func initGRPCCollection() *GRPCCollection {

	wire.Build(IteratorCollection)

	return &GRPCCollection{}
}
