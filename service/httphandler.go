package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/origin-yaropolk/testgame-service/testgame"
)

func getTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.storage.Tests)
}

func getTestByID(c *gin.Context) {
	id := c.Param("id")

	for _, test := range service.storage.Tests {
		if test.ID == id {
			c.IndentedJSON(http.StatusOK, test)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "test not found"})
}

func postTest(c *gin.Context) {
	var newTest testgame.Test

	if err := c.BindJSON(&newTest); err != nil {
		return
	}

	service.storage.Tests = append(service.storage.Tests, newTest)
	c.IndentedJSON(http.StatusCreated, newTest)
}
