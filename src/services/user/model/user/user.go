package user_model

import (
	"database/sql"
	"time"
	pb "user-service/grpc_build/user"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserModel struct {
	Id         string
	Username   string
	FirstName  string
	SecondName string
	Patronymic *string
	Email      *string
	Phone      *string
	Birth      *time.Time
}

func (u_m UserModel) ConvertToGRPC() *pb.UserData {
	var birth *timestamppb.Timestamp
	if u_m.Birth != nil {
		birth = timestamppb.New(*u_m.Birth)
	}
	return &pb.UserData{Id: u_m.Id, Username: u_m.Username, FirstName: u_m.FirstName, SecondName: u_m.SecondName, Patronymic: u_m.Patronymic, Email: u_m.Email, Phone: u_m.Phone, Birth: birth}
}

func ConvertToUserModel(u_d *pb.UserData) *UserModel {
	var birth *time.Time
	if u_d.Birth != nil {
		birth = new(time.Time)
		*birth = u_d.Birth.AsTime()
	}

	return &UserModel{Id: u_d.Id, Username: u_d.Username, FirstName: u_d.FirstName, SecondName: u_d.SecondName, Patronymic: u_d.Patronymic, Email: u_d.Email, Phone: u_d.Phone, Birth: birth}
}

func ConvertToUserModelFromDBRow(row pgx.Rows) (*UserModel, error) {
	var u_m UserModel

	var patronymic sql.NullString
	var email sql.NullString
	var phone sql.NullString
	var birth sql.NullTime

	err := row.Scan(&u_m.Id, &u_m.Username, &u_m.FirstName, &u_m.SecondName, &patronymic, &email, &phone, &birth)

	if err != nil {
		return nil, err
	}

	if patronymic.Valid {
		u_m.Patronymic = &patronymic.String
	}
	if email.Valid {
		u_m.Email = &email.String
	}
	if phone.Valid {
		u_m.Phone = &phone.String
	}
	if birth.Valid {
		u_m.Birth = &birth.Time
	}
	return &u_m, nil
}
