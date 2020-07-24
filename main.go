package main

import (
	"fmt"
	"log"
	"os"

	"github.com/apiraport/config"
	"github.com/apiraport/controller"
	"github.com/apiraport/middleware"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	config.InitDB()
	defer config.DB.Close()

	fmt.Printf("Welcome API \n")
	router := gin.Default()
	url := router.Group("/appraport/")
	{
		url.POST("/auth", controller.Login)
		url.GET("/raport", middleware.CheckJWT(), controller.Getraport)
		url.GET("/raportdetail", middleware.CheckJWT(), controller.Getraportdetail)
	}
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
