package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/controllers"
	"main.go/repositories"
)

func Init(addr, port string, db *mongo.Database) {
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)
	r := gin.Default()
	initRouter(r, userController)
	r.Run(fmt.Sprintf(addr + ":" + port))
}
