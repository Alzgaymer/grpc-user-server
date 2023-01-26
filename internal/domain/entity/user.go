package entity

import proto_user_model "github.com/Alzgaymer/grpc-rest-server/gen/go/proto/model/v1"

type User struct {
	Name  string
	Email string
	Id    string
	Age   string
}

func (u *User) ToProto() *proto_user_model.User {
	return &proto_user_model.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}
