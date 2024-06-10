package routes

import (
	"github.com/gin-gonic/gin"
	"simple/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
