package user_model

import (
	pb "api-gateway/grpc_build/user"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UserModel struct {
	Id         string     `json:"id"`
	Username   string     `json:"username"`
	FirstName  string     `json:"first_name"`
	SecondName string     `json:"second_name"`
	Patronymic *string    `json:"patronymic"`
	Email      *string    `json:"email"`
	Phone      *string    `json:"phone"`
	Birth      *time.Time `json:"birth"`
}

func (u_m UserModel) ConvertToGRPC() *pb.UserData {
	var birth *timestamppb.Timestamp
	if u_m.Birth != nil {
		birth = timestamppb.New(*u_m.Birth)
	}
	return &pb.UserData{Id: u_m.Id, Username: u_m.Username, FirstName: u_m.FirstName, SecondName: u_m.SecondName, Patronymic: u_m.Patronymic, Email: u_m.Email, Phone: u_m.Phone, Birth: birth}
}

func ConvertToModel(u_d *pb.UserData) *UserModel {
	var birth *time.Time
	if u_d.Birth != nil {
		birth = new(time.Time)
		*birth = u_d.Birth.AsTime()
	}

	return &UserModel{Id: u_d.Id, Username: u_d.Username, FirstName: u_d.FirstName, SecondName: u_d.SecondName, Patronymic: u_d.Patronymic, Email: u_d.Email, Phone: u_d.Phone, Birth: birth}
}
