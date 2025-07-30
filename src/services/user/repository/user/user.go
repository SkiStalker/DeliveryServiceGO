package user_repository

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"user-service/model/user"
	"user-service/util"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool     *pgxpool.Pool
	pageSize int64
}

func CreateUserRepository() *UserRepository {
	var user_db *UserRepository

	dsn := util.GetDBDSN()
	if dsn == "" {
		log.Fatal("DB_DSN not set in environment")
	}

	pageSizeStr := os.Getenv("DB_PAGE_SIZE")
	var pageSize int64 = 1000

	if pageSizeStr != "" {
		pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
		if err != nil || pageSize < 1 {
			log.Fatal("DB_PAGE_SIZE isn't natural number")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	user_db = &UserRepository{pageSize: pageSize}

	user_db.pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	err = user_db.pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return user_db
}

func (u_r UserRepository) Close() {
	u_r.pool.Close()
}

func (u_r UserRepository) GetUser(ctx context.Context, user_id string) (*user_model.UserModel, error) {

	args := pgx.NamedArgs{
		"userId": user_id,
	}

	rows, err := u_r.pool.Query(ctx, "SELECT id, username, first_name, second_name, patronymic, email, phone, birth FROM account WHERE id = @userId", args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[user_model.UserModel])

	return &res, err
}

func (u_r UserRepository) SearchUsers(ctx context.Context, page int32, first_name string, second_name string) ([]user_model.BriefUserModel, error) {
	args := pgx.NamedArgs{
		"firstName":  first_name,
		"secondName": second_name,
		"limit":      u_r.pageSize,
		"offset":     u_r.pageSize * int64(page),
	}

	rows, err := u_r.pool.Query(ctx, "SELECT account.id, account.username, account.first_name, account.second_name from account WHERE is_active = TRUE and (account.first_name LIKE @firstName and account.second_name LIKE @secondName) LIMIT @limit OFFSET @offset", args)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[user_model.BriefUserModel])
}
