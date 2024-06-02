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

	user := models.NewUserModel(
		userDto.Name,
		userDto.LastName,
		userDto.Cpf,
		userDto.City,
		userDto.State,
		userDto.Address,
		userDto.Mail,
		userDto.Role,
		userDto.Password,
	)
	post, err := controller.userRepository.SaveUser(user)
	if err != nil {
		slog.Error("error on post user")
		shared.ResponseError(context, http.StatusBadRequest, "error on post user")
		return
	}

	shared.ResponseSucess(context, "ok", &post)
}
