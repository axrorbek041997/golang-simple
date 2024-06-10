package routes

import (
	"github.com/gin-gonic/gin"
	"simple/controllers"
)

func AuthRouter(router *gin.RouterGroup) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
