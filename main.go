package main

import (
	"context"
	"log"

	"github.com/VictorBion/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/VictorBion/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/VictorBion/meu-primeiro-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to dabase, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}