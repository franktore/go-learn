package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Name string = "World"

func GetGreetingById(c *gin.Context) {
	content := getMockGreeting(c)
	c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
		"Title":   "",
		"Content": content,
	})
}

func getMockGreeting(c *gin.Context) string {
	var message string
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Print(err)
		return fmt.Sprintf("Something went wrong %v. "+
			"Did you specify `id` as a positive integer?", Name)
	}

	session := sessions.Default(c)
	v := session.Get("user-id")

	if v == nil {
		name, exists := c.GetQuery("name")
		if exists {
			Name = name
		}
	} else {
		Name = v.(string)
	}

	_, greeting, err := getGreetingById(id)
	if err != nil {
		log.Print(err)
		message = "Sorry %v, no greeting for you"
	} else {
		message = greeting.Message
	}

	if message == "" {
		log.Print("No greeting found for id: ", id)
		message = "Sorry %v, no greeting for you"
	}

	message = fmt.Sprintf(message, Name)
	return message
}
