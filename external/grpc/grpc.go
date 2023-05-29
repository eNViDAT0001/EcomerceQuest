package grpc

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"google.golang.org/grpc"
	"log"
)

var (
	v = wrap_viper.GetViper()
)

func RecommenderClient() *grpc.ClientConn {
	connStr := fmt.Sprintf("%s:%s", v.GetString("RECOMMENDER.HOST"), v.GetString("RECOMMENDER.PORT"))
	conn, err := grpc.Dial(connStr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return conn
}
