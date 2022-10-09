package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/franktore/go-learn/pkg/mocks"
	"github.com/franktore/go-learn/pkg/models"
	"github.com/gin-gonic/gin"
)

func UpdateGreeting(c *gin.Context) {
	var updatedGreeting models.Greeting
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Print(err)
		return
	}

	// Call BindJson to bind the received JSON to updatedGreeting.
	if err := c.BindJSON(&updatedGreeting); err != nil {
		log.Print(err)
		return
	}

	if !strings.Contains(updatedGreeting.Message, "%v") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "greeting must contain %v placeholder for name"})
		return
	}

	idx, greeting, err := getGreetingById(id)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "greeting not found"})
		return
	}

	if greeting.Message == updatedGreeting.Message {
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": "greeting not modified"})
		return
	}

	mocks.Greetings[idx].Message = updatedGreeting.Message
	c.JSON(http.StatusOK, greeting)
}
