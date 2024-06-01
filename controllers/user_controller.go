package controllers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/dto"
	"main.go/models"
	"main.go/repositories"
	"main.go/shared"
)

type UserController interface {
	PostUser(context *gin.Context)
}

type userControllerImpl struct {
	userRepository repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) userControllerImpl {
	return userControllerImpl{userRepository: repo}
}

func (controller userControllerImpl) PostUser(context *gin.Context) {
	userDto := &dto.UserDto{}
	context.BindJSON(&userDto)

	user := &models.UserModel{
		Name:     userDto.Name,
		LastName: userDto.LastName,
		Cpf:      userDto.Cpf,
		City:     userDto.City,
		State:    userDto.State,
		Address:  userDto.Address,
		Password: userDto.Password,
		Mail:     userDto.Mail,
		Role:     userDto.Role,
	}
	post, err := controller.userRepository.SaveUser(user)
	if err != nil {
		slog.Error("error on post user")
		shared.ResponseError(context, http.StatusBadRequest, "error on post user")
		return
	}

	shared.ResponseSucess(context, "ok", &post)
}
