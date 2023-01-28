package router

import (
	"github.com/gin-gonic/gin"
	"github.com/folklore13/task-5-vix-btpns-ario/controllers"
	"github.com/folklore13/task-5-vix-btpns-ario/database"
	"github.com/folklore13/task-5-vix-btpns-ario/helpers"
	"github.com/folklore13/task-5-vix-btpns-ario/models"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = database.DatabaseConnection()
	userRepo       models.UserRepo            = models.NewUserRepo(db)
	authHelpers    helpers.AuthHelper         = helpers.NewAuthHelper(userRepo)
	jwtHelpers     helpers.JWTHelper          = helpers.NewJWTHelper()
	authController controllers.AuthController = controllers.NewAuthController(authHelpers, jwtHelpers)
)

func AuthRouter() {
	router := gin.Default()
	authRoutes := router.Group("api/users")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	router.Run()
}