package client

import (
	"fmt"

	"github.com/Shemetov-Sergey/GoCensor-service/pkg/config"
	"github.com/Shemetov-Sergey/GoCensor-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CensorServiceClient
}

func InitServiceClient(c *config.Config) pb.CensorServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCensorServiceClient(cc)
}
