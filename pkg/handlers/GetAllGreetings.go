package handlers

import (
	"net/http"

	"github.com/franktore/go-learn/pkg/mocks"
	"github.com/franktore/go-learn/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllGreetings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAllMockGreetings())
}

func getAllMockGreetings() []models.Greeting {
	return mocks.Greetings
}
