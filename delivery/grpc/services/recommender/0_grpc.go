package recommender

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
	"google.golang.org/grpc"
)

var (
	v = wrap_viper.GetViper()
)

func GRPCClient() (*grpc.ClientConn, proto.RecommenderBaseCommentClient, error) {
	connStr := fmt.Sprintf("%s:%s", v.GetString("RECOMMENDER.HOST"), v.GetString("RECOMMENDER.PORT"))
	conn, err := grpc.Dial(connStr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := proto.NewRecommenderBaseCommentClient(conn)

	return conn, client, nil
}
