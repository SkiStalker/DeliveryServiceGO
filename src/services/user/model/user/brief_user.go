package user_model

import (
	pb "user-service/grpc_build/user"

	"github.com/jackc/pgx/v5"
)

type BriefUserModel struct {
	Id         string
	Username   string
	FirstName  string
	SecondName string
}

func (b_u_m BriefUserModel) ConvertToGRPC() *pb.BriefUserData {
	return &pb.BriefUserData{Id: b_u_m.Id, Username: b_u_m.Username, FirstName: b_u_m.FirstName, SecondName: b_u_m.SecondName}
}

func ConvertToBriefUserModel(u_d *pb.BriefUserData) *BriefUserModel {

	return &BriefUserModel{Id: u_d.Id, Username: u_d.Username, FirstName: u_d.FirstName, SecondName: u_d.SecondName}
}

func ConvertToBriefUserModelFromDBRow(row pgx.Rows) (*BriefUserModel, error) {
	var u_m BriefUserModel

	err := row.Scan(&u_m.Id, &u_m.Username, &u_m.FirstName, &u_m.SecondName)

	if err != nil {
		return nil, err
	}
	return &u_m, nil
}
