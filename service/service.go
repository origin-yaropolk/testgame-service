package service

import (
	"github.com/gin-gonic/gin"
	"github.com/origin-yaropolk/testgame-service/storage"
)

type Service struct {
	storage *storage.Storage
}

var service *Service

func RunService() {
	service = &Service{
		storage: &storage.Storage{},
	}

	router := gin.Default()

	router.GET("/tests", getTest)
	router.GET("/tests/:id", getTestByID)
	router.POST("/test", postTest)

	println(len(service.storage.Tests))

	router.Run("localhost:8080")
}
