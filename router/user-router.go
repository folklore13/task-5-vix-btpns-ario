package router

import (
	"github.com/gin-gonic/gin"
	"github.com/folklore13/task-5-vix-btpns-ario/controllers"
	"github.com/folklore13/task-5-vix-btpns-ario/helpers"
	"github.com/folklore13/task-5-vix-btpns-ario/middlewares"
)

var (
	userHelpers    helpers.UserHelper         = helpers.NewUserHelper(userRepo)
	userController controllers.UserController = controllers.NewUserController(userHelpers, jwtHelpers)
)

func UserRouter () {
	router := gin.Default()

	userRoutes := router.Group("api/users", middlewares.AuthJWT(jwtHelper))
	{
		userRoutes.GET("/:userId", userController.Profile)
		userRoutes.PUT("/:userId", userController.Update)
	}
	router.Run()
}