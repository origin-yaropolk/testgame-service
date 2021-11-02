package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "")
}
