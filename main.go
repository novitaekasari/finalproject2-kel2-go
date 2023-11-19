package main

import (
	"finalprojec2-kel-go/models"
	"finalprojec2-kel-go/routers"
)

func main() {
	models.ConnectDatabase()
	routers.StartServer().Run(":8080")
}