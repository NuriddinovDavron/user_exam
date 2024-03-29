package mongodb

import (
	"context"
	"time"
	pbu "user_exam/genproto/user_exam"
	"user_exam/pkg/logger"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	collection *mongo.Collection
	log        logger.Logger
}

func NewUserRepo(collection *mongo.Collection, log logger.Logger) *userRepo {
	return &userRepo{
		collection: collection,
		log:        log,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, req *pbu.CreateUserRequest) (*pbu.UserApi, error) {
	var user pbu.UserApi
	user.Id = uuid.NewString()
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUserById(ctx context.Context, req *pbu.GetUserByIdRequest) (*pbu.UserApi, error) {
	var response pbu.UserApi
	filter := bson.M{"id": req.UserId}
	err := r.collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *userRepo) GetAllUser(ctx context.Context, req *pbu.GetAllUserRequest) (*pbu.GetAllUserResponse, error) {
	var response pbu.GetAllUserResponse

	reqOptions := options.Find()

	reqOptions.SetSkip(int64((req.Page - 1) * req.Limit))
	reqOptions.SetLimit(int64(req.Limit))

	cursor, err := r.collection.Find(ctx, bson.M{}, reqOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user pbu.User
		err = cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		response.Users = append(response.Users, &user)
	}

	return &response, nil

}

func (r *userRepo) UpdateUser(ctx context.Context, req *pbu.User) (*pbu.UserApi, error) {
	var response pbu.UserApi

	filter := bson.M{"_id": req.Id}

	updateReq := bson.M{
		"$set": bson.M{
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"email":      req.Email,
			"password":   req.Password,
			"updated_at": time.Now(),
		},
	}

	err := r.collection.FindOneAndUpdate(ctx, filter, updateReq).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, req *pbu.GetUserByIdRequest) error {
	filter := bson.M{"id": req.UserId}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func(r *userRepo) CheckField(ctx context.Context, req *pbu.CheckUser) (*pbu.CheckRes, error){
	filter := bson.M{req.Field: req.Field}
	err := r.collection.FindOne(ctx, filter)
	if err != nil {
		return &pbu.CheckRes{Exists: false}, err.Err()
	}

	return &pbu.CheckRes{Exists: true}, nil
}

func(r *userRepo) GetUserByEmail(ctx context.Context, req *pbu.EmailRequest) (*pbu.UserApi, error){
	var response pbu.UserApi
	filter := bson.M{"email": req.Email}
	err := r.collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func(r *userRepo) GetUserByRefreshToken(ctx context.Context, req *pbu.UserToken) (*pbu.UserApi,error){
	var response pbu.UserApi
	filter := bson.M{"refresh_token": req.RefreshToken}
	err := r.collection.FindOne(ctx, filter).Decode(&response)
	if err != nil{
		return nil,err
	}
	return &response, nil
}
