package main

import (
	"github.com/gin-gonic/gin"
	route "go-backend-clean-architecture/api/route"
	"go-backend-clean-architecture/config"
	"time"
)

func main() {

	app := config.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
