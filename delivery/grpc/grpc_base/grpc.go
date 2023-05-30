package grpc_base

var services *GRPCCollection

func GetServices() *GRPCCollection {
	if services == nil {
		services = initGRPCCollection()
	}

	return services
}
