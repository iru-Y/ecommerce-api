package routes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func initRouter(router *gin.Engine, u controllers.UserController) {
	slog.Info("Router initalizated")
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
        v1.POST("/user", u.PostUser)
	}
}
