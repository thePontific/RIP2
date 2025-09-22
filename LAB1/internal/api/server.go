package api

import (
	"LAB1/internal/app/handler"
	"LAB1/internal/app/repository"
	"LAB1/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	log.Println("Starting server")

	repo, err := repository.NewRepository()
	if err != nil {
		logrus.Error("ошибка инициализации репозитория")
	}

	minioService := service.NewMinioService()
	h := handler.NewHandler(repo, minioService)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./resources")

	r.GET("/Andromeda", h.GetStars)
	r.GET("/Andromeda/star/:id", h.GetStarDetails)

	r.GET("/cart/:id", h.GetCartDetails)

	r.Run()
	log.Println("Server down")
}
