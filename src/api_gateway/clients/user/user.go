package user_client

import (
	pb "api-gateway/grpc_build/user"
	"context"
	"fmt"
	"log"
	"os"

	"api-gateway/model/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	userConn pb.UserServiceClient
	grpcConn *grpc.ClientConn
}

func CreateUserServiceClient() *UserServiceClient {

	host := os.Getenv("USER_SERVICE_HOST")
	if host == "" {
		host = "user"
	}

	port := os.Getenv("USER_SERVICE_PORT")
	conn, err := grpc.NewClient(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to UserService: %v", err)
	}

	return &UserServiceClient{userConn: pb.NewUserServiceClient(conn), grpcConn: conn}
}

func (u_clt UserServiceClient) Close() {
	u_clt.grpcConn.Close()
}

func (u_clt UserServiceClient) GetUser(ctx context.Context, id string) (*user_model.UserModel, error) {
	resp, err := u_clt.userConn.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	} else {

		user_data := resp.GetUserData()
		if user_data == nil {
			return nil, fmt.Errorf("null reference to returned data")
		} else {
			return user_model.ConvertToModel(user_data), nil
		}
	}
}
