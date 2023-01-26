package v1

import (
	"context"
	"grpc-user-service/internal/domain/entity"

	proto_user_model "github.com/Alzgaymer/grpc-rest-server/gen/go/proto/model/v1"
	proto_user_service "github.com/Alzgaymer/grpc-rest-server/gen/go/proto/service/v1"
)

type UserServer struct {
	proto_user_service.UnimplementedUserServiceServer
}

func NewUserServer(unimplementedUserServiceServer proto_user_service.UnimplementedUserServiceServer) *UserServer {
	return &UserServer{UnimplementedUserServiceServer: unimplementedUserServiceServer}
}

func (s *UserServer) GetUsers(
	context.Context,
	*proto_user_service.GetUsersRequest) (*proto_user_service.GetUsersResponse, error) {
	U := entity.User{
		Id:    "id",
		Name:  "Slavik",
		Age:   "19",
		Email: "email@gmail.com",
	}
	return &proto_user_service.GetUsersResponse{
		Users: []*proto_user_model.User{
			U.ToProto(),
		},
	}, nil
}

func (s *UserServer) UpdateUser(
	context.Context,
	*proto_user_service.UpdateUserRequest) (*proto_user_service.UpdateUserResponse, error) {
	return nil, nil
}
