package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/franktore/go-learn/pkg/mocks"
	"github.com/franktore/go-learn/pkg/models"
)

func AddGreeting(c *gin.Context) {
	var newGreeting models.Greeting

	// Call BindJson to bind the received JSON to newGreeting.
	if err := c.BindJSON(&newGreeting); err != nil {
		return
	}

	if !strings.Contains(newGreeting.Message, "%v") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "greeting must contain %v placeholder for name"})
		return
	}
	newGreeting.Id = int64(mocks.Seq_id + 1)

	mocks.Greetings = append(mocks.Greetings, newGreeting)
	mocks.Seq_id++
	c.IndentedJSON(http.StatusCreated, newGreeting)
}
