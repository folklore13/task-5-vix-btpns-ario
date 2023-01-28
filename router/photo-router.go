package router

import (
	"github.com/gin-gonic/gin"
	"github.com/folklore13/task-5-vix-btpns-ario/controllers"
	"github.com/folklore13/task-5-vix-btpns-ario/helpers"
	"github.com/folklore13/task-5-vix-btpns-ario/middlewares"
	"github.com/folklore13/task-5-vix-btpns-ario/models"
)

var (
	photoRepo       models.PhotoRepo            = models.NewPhotoRepo(db)
	photoHelper     helpers.PhotoHelper         = helpers.NewPhotoHelper(photoRepo)
	jwtHelper       helpers.JWTHelper           = helpers.NewJWTHelper()
	photoController controllers.PhotoController = controllers.NewPhotoController(photoHelper, jwtHelper)
)

func PhotoRouter() {
	router := gin.Default()

	photoRoutes := router.Group("api/photo", middlewares.AuthJWT(jwtHelper))
	{
		photoRoutes.POST("/photos", photoController.Insert)
		photoRoutes.PUT("/:photoId", photoController.Update)
		photoRoutes.GET("/photos", photoController.GetByID)
		photoRoutes.DELETE("/:photoId", photoController.Delete)
	}
	router.Run()
}