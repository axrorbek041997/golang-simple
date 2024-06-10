package routes

import (
	"github.com/gin-gonic/gin"
	"simple/middleware"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.Group("/auth")
	{
		AuthRouter(router)
	}
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		UserRoutes(router)
	}
}
