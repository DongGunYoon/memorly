package routes

import (
	"memorly/controllers"
	"memorly/middleware"

	"github.com/gin-gonic/gin"
)

func FolderRoute(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.POST("/folder", controllers.CreateFolder())
}
