package repository

import (
	"context"
	"time"

	"github.com/ntp7758/shopping-app-backend/libs/databases"
	"github.com/ntp7758/shopping-app-backend/services/user/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Insert(user domain.User) (*mongo.InsertOneResult, error)
	GetByID(id string) (*domain.User, error)
	GetByAuthId(authId string) (*domain.User, error)
}

const userCollection string = "user"

type userRepository struct {
	ctx      context.Context
	dbClient databases.MongoDBClient
}

func NewUserRepository(dbClient databases.MongoDBClient) (UserRepository, error) {
	// err := dbClient.CreateCollection(userCollection)
	// if err != nil {
	// 	return nil, err
	// }

	return &userRepository{ctx: context.TODO(), dbClient: dbClient}, nil
}

func (r *userRepository) Insert(user domain.User) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	result, err := r.dbClient.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	result := r.dbClient.Collection(userCollection).FindOne(ctx, bson.M{"_id": oID})
	err = result.Err()
	if err != nil {
		return nil, err
	}

	var user *domain.User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByAuthId(authId string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	result := r.dbClient.Collection(userCollection).FindOne(ctx, bson.M{"authId": authId})
	err := result.Err()
	if err != nil {
		return nil, err
	}

	var user *domain.User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}
