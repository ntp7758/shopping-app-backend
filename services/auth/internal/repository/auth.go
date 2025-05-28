package repository

import (
	"context"
	"time"

	"github.com/ntp7758/shopping-app-backend/libs/databases"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	Insert(auth domain.Auth) (string, error)
	GetByID(id string) (*domain.Auth, error)
	GetByUsername(username string) (*domain.Auth, error)
}

const authCollection string = "auth"

type authRepository struct {
	ctx      context.Context
	dbClient databases.MongoDBClient
}

func NewAuthRepository(dbClient databases.MongoDBClient) (AuthRepository, error) {
	// err := dbClient.CreateCollection(authCollection)
	// if err != nil {
	// 	return nil, err
	// }

	return &authRepository{ctx: context.TODO(), dbClient: dbClient}, nil
}

func (r *authRepository) Insert(auth domain.Auth) (string, error) {

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	result, err := r.dbClient.Collection(authCollection).InsertOne(ctx, auth)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil
}

func (r *authRepository) GetByID(id string) (*domain.Auth, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	result := r.dbClient.Collection(authCollection).FindOne(ctx, bson.M{"_id": oID})
	err = result.Err()
	if err != nil {
		return nil, err
	}

	var auth *domain.Auth
	err = result.Decode(&auth)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *authRepository) GetByUsername(username string) (*domain.Auth, error) {

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	result := r.dbClient.Collection(authCollection).FindOne(ctx, bson.M{"username": username})
	err := result.Err()
	if err != nil {
		return nil, err
	}

	var auth *domain.Auth
	err = result.Decode(&auth)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return auth, nil
}
