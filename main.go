package main

import (
	"github.com/gin-gonic/gin"
	"simple/config"
	"simple/routes"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	//// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
		c.Next()
	})

	api := r.Group("/api")
	routes.SetupRoutes(api)

	err1 := r.Run("127.0.0.1:8081")
	if err1 != nil {
		return
	}
}
