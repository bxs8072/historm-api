package main

import (
	"fmt"
	"historm_api/controllers"
	"historm_api/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(middlewares.Hash(os.Getenv("JWT_SECRET")))
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)

	var app *gin.Engine = gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middlewares.ApiMiddleware())

	app.POST("/api/v1/anime", controllers.RetrieveAnimeList)
	app.POST("/api/v1/latest-anime", controllers.RetrieveLatestAnimeList)
	app.POST("/api/v1/anime-detail", controllers.RetrieveAnimeDetail)
	app.POST("/api/v1/video", controllers.RetrieveVideo)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Printf("Server is running in address http://%s:%s", host, port)
	app.Run(host + ":" + port)

}
