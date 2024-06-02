package repositories

import (
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/models"
	"main.go/shared"
)

type UserRepository interface {
	FindAll(page int64, size int64) (*shared.PagedUser, error)
	//FindUserById(id string) (*models.UserModel, error)
	SaveUser(user *models.UserModel) (*models.UserModel, error)
	//UptadeUser(id string, user *models.UserModel) (*models.UserModel, error)
	//DeleteUser(id string) error
}

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func NewUserRepository(db *mongo.Database) userRepositoryImpl {
	return userRepositoryImpl{Connection: db}
}

func (repo userRepositoryImpl) FindAll(page, size int64) (*shared.PagedUser, error) {
	//users := &[]models.UserModel{}
	//filter := bson.M{}
	//collection := repo.Connection.Collection("users")

	return nil, nil
}

func (repo userRepositoryImpl) SaveUser(user *models.UserModel) (*models.UserModel, error) {
	if err != nil {
		slog.Error("Deu ruim na 42")
	}
	user.Id = primitive.NewObjectID()
	_, err := repo.Connection.Collection("users").InsertOne(ctx, user)
	if err != nil {
		slog.Error(fmt.Sprintf("error: %v", err))
		return nil, err
	}

	return user, nil
}
