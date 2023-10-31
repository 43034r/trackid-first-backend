package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"monitoriong.wiki/trackid-first-backend/controllers"
	"monitoriong.wiki/trackid-first-backend/database"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/trackids/:id", controllers.ReadTrackid)
	r.GET("/trackids", controllers.ReadTrackids)
	r.POST("/trackids", controllers.CreateTrackid)
	r.PUT("/trackids/:id", controllers.UpdateTrackid)
	r.DELETE("/trackids/:id", controllers.DeleteTrackid)
	r.Run(":5000")
}
