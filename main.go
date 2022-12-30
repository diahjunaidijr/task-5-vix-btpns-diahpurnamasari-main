package main

import (
	"task-5-vix-btpns-Diah_Purnama_Sari/controllers"
	"task-5-vix-btpns-Diah_Purnama_Sari/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := controllers.Server{}
	server.Initialize()
	router.InitRoutes(&server)
	server.Run(8000)
}
