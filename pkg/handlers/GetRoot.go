package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type Post struct {
	Title   string
	Content template.HTML
}

func GetRoot(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Home Page"})
}

func GetRootMd(c *gin.Context) {
	var posts []string

	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusServiceUnavailable, "error.tmpl.html", nil)
		c.Abort()
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			fmt.Println(file.Name())
			posts = append(posts, file.Name())
		}
	}

	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"posts": posts,
	})
}

func GetMarkdown(c *gin.Context) {
	postName := c.Param("postName")

	mdfile, err := ioutil.ReadFile("./" + postName)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
		c.Abort()
		return
	}

	postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(mdfile)))

	post := Post{Title: postName, Content: postHTML}

	c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
		"Title":   post.Title,
		"Content": post.Content,
	})
}
