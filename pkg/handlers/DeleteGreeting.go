package handlers

import (
	"net/http"
	"strconv"

	"github.com/franktore/go-learn/pkg/mocks"
	"github.com/gin-gonic/gin"
)

func DeleteGreeting(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return
	}

	idx, _, err := getGreetingById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "greeting not found"})
	}

	mocks.Greetings, err = removeByIdx(mocks.Greetings, idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting greeting"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "greeting deleted"})
}
