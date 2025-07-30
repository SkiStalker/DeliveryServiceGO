package main

import (
	"context"
	"errors"

	pb "user-service/grpc_build/user"
	user_model "user-service/model/user"
	"user-service/repository/user"
	"user-service/util"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	user_rep *user_repository.UserRepository
}

func CreateUserServer() *UserServer {
	return &UserServer{user_rep: user_repository.CreateUserRepository()}
}

func (u_s *UserServer) Close() {
	u_s.user_rep.Close()
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := s.user_rep.GetUser(ctx, req.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &pb.GetUserResponse{}, status.Errorf(codes.NotFound, "User with specified id doesn't exist")
		} else {
			return &pb.GetUserResponse{}, status.Errorf(codes.Internal, "Internal error : %v", err)
		}
	}
	return &pb.GetUserResponse{UserData: res.ConvertToGRPC()}, nil
}

func (s *UserServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {

	firstName := ""
	if req.FirstName != nil {
		firstName = *req.FirstName
	}
	secondName := ""

	if req.SecondName != nil {
		secondName = *req.SecondName
	}

	res, err := s.user_rep.SearchUsers(ctx, req.Page, firstName, secondName)
	if err != nil {
		return &pb.SearchUsersResponse{}, status.Errorf(codes.Internal, "Internal error : %v", err)
	} else {

		return &pb.SearchUsersResponse{Users: &pb.BriefUserArray{
			Arr: util.Map(res, func(item user_model.BriefUserModel) *pb.BriefUserData {
				return item.ConvertToGRPC()
			})}}, nil
	}
}
