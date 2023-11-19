package main

import (
	"finalprojec2-kel-go/models"
	"finalprojec2-kel-go/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := routers.StartServer()

	// Main router
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Railway API",
		})
	})
	
	err := r.Run()
	if err != nil {
		panic(err)
	}

}