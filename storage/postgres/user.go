package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	pb "user_exam/genproto/user_exam"
)

type UserRepo struct {
	db *sqlx.DB
}

func (u UserRepo) CheckField(req *pb.CheckUser) (*pb.CheckRes, error) {
	var existsClient int
	query := fmt.Sprintf("SELECT count(1) FROM users_exam WHERE %s = $1 AND deleted_at IS NULL", req.Field)
	if err := u.db.QueryRow(
		query,
		req.Value,
	).Scan(&existsClient); err != nil {
		return nil, err
	}
	if existsClient == 0 {
		return &pb.CheckRes{
			Exists: false,
		}, nil
	}
	return &pb.CheckRes{
		Exists: true,
	}, nil
}

func (u UserRepo) GetUserByEmail(req *pb.EmailRequest) (*pb.UserApi, error) {
	query := `
	SELECT 
		id,
		firstname,
		lastname,
		email,
		password,
		created_at,
		updated_at
	FROM 
		users_exam
	WHERE
		email = $1
	`
	var responseUser pb.UserApi
	if err := u.db.QueryRow(
		query,
		req.Email,
	).Scan(
		&responseUser.Id,
		&responseUser.FirstName,
		&responseUser.LastName,
		&responseUser.Email,
		&responseUser.Password,
		&responseUser.CreatedAt,
		&responseUser.UpdatedAt,
	); err != nil {
		log.Println("Error getting user by email")
		return nil, err
	}
	return &responseUser, nil
}

func (u UserRepo) GetUserByRefreshToken(req *pb.UserToken) (*pb.UserApi, error) {
	query := `
	SELECT 
		id,
		firstname,
		lastname,
		email,
		password,
		refresh_token,
		created_at,
		updated_at
	FROM 
		users_exam
	WHERE
		refresh_token = $1
	AND
		deleted_at IS NULL
	`
	var responseUser pb.UserApi
	if err := u.db.QueryRow(
		query,
		req.RefreshToken,
	).Scan(
		&responseUser.Id,
		&responseUser.FirstName,
		&responseUser.LastName,
		&responseUser.Email,
		&responseUser.Password,
		&responseUser.RefreshToken,
		&responseUser.CreatedAt,
		&responseUser.UpdatedAt,
	); err != nil {
		log.Println("Error getting user by token")
		return nil, err
	}
	return &responseUser, nil
}

func (u UserRepo) CreateUser(mailReq *pb.CreateUserRequest) (*pb.UserApi, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	Id := id.String()

	var res pb.UserApi
	query := `INSERT INTO users_exam (id, firstname, lastname, email, password, refresh_token) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, firstname, lastname, email, password, refresh_token, created_at, updated_at`
	err = u.db.QueryRow(query, Id, mailReq.FirstName, mailReq.LastName, mailReq.Email, mailReq.Password, mailReq.RefreshToken).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Password,
		&res.RefreshToken,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u UserRepo) GetUserById(isUnReq *pb.GetUserByIdRequest) (*pb.UserApi, error) {
	var res pb.UserApi
	query := `SELECT firstname, lastname, email, password FROM users_exam WHERE id=$1`
	err := u.db.QueryRow(query, isUnReq.UserId).Scan(
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Password)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u UserRepo) GetAllUser(crUsReq *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	var allUser pb.GetAllUserResponse
	query := `select id, firstname, lastname, email, password from users_exam limit $1 offset $2`
	offset := crUsReq.Limit * (crUsReq.Page - 1)
	rows, err := u.db.Query(query, crUsReq.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		allUser.Users = append(allUser.Users, &user)
	}
	return &allUser, nil
}

func (u UserRepo) UpdateUser(logInReq *pb.User) (*pb.UserApi, error) {
	var user pb.UserApi
	query := `update users_exam set firstname=$1, lastname=$2, email=$3, password=$4 where id=$5 returning id, firstname, lastname, email, password`
	err := u.db.QueryRow(query, logInReq.FirstName, logInReq.LastName, logInReq.Email, logInReq.Password, logInReq.Id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepo) DeleteUser(logInReq *pb.GetUserByIdRequest) error {
	query := `delete from users_exam where id=$1`
	err := u.db.QueryRow(query, logInReq.UserId).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}
