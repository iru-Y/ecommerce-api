package routes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func initUserRouter(router *gin.Engine, u controllers.UserController) {
	slog.Info("Router initalizated")
	v1 := router.Group(basePath)
	{

		v1.POST("/user", u.PostUser)
	}
}
