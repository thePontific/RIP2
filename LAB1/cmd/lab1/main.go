package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"LAB1/internal/app/config"
	"LAB1/internal/app/dsn"
	"LAB1/internal/app/handler"
	"LAB1/internal/app/repository"
	"LAB1/internal/pkg"
	"LAB1/internal/service"
)

func main() {
	router := gin.Default()
	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("error loading config: %v", err)
	}

	postgresString := dsn.FromEnv()
	fmt.Println(postgresString)

	rep, errRep := repository.New(postgresString)
	if errRep != nil {
		logrus.Fatalf("error initializing repository: %v", errRep)
	}

	minioService := service.NewMinioService()
	hand := handler.NewHandler(rep, minioService)

	application := pkg.NewApp(conf, router, hand)
	application.RunApp()
}
