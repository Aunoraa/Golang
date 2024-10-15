// @title API Documentation
// @version 1.0
// @description This is a sample API documentation for the Go Backend Clean Architecture project.
// @host localhost:8080
// @BasePath /
package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	route "go-backend-clean-architecture/api/route"
	"go-backend-clean-architecture/cmd/docs"
	"go-backend-clean-architecture/configs"
	"time"
)

// @title API Documentation
// @version 1.0
// @description This is a sample API documentation for the Go Backend Clean Architecture project.
// @host localhost:8080
// @BasePath /

func main() {

	app := configs.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

}
