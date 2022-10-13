package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"

	"github.com/franktore/go-learn/pkg/greetings"
	"github.com/franktore/go-learn/pkg/handlers"
	"github.com/franktore/go-learn/pkg/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const log_prefix string = "greetings: "

var config = Configuration{}
var CONFIGDIR = "./"

type Configuration struct {
	WORKDIR string
	HOST    string
	PORT    string
	AUTH    bool
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("WorkDir: " + path)

	config, _ = GetConfig()
}

func GetConfig(params ...string) (Configuration, error) {
	configuration := Configuration{}
	var env string
	if _, err := os.Stat(CONFIGDIR + "prod_conf.json"); err == nil {
		fmt.Printf("PROD ENV\n")
		env = "prod"
	} else {
		fmt.Printf("DEV ENV\n")
		env = "dev"
	}
	// gin.SetMode(gin.ReleaseMode)
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf(CONFIGDIR+"%s_conf.json", env)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &configuration); err != nil {
		log.Println("unable to marshal data")
	}

	return configuration, err
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix(log_prefix)
	log.SetFlags(0)

	// declare name variable
	name := ""

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		name = argsWithoutProg[0]
	}

	message, _ := greetings.Hello(name)

	// If an error was returned, print it to the console and
	// exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	handlers.Name = name

	var router *gin.Engine
	if _, err := os.Stat(CONFIGDIR + "creds.json"); err == nil {
		fmt.Printf("setup router with authent\n")
		router = setup_router_auth()
	} else {
		fmt.Printf("setup router\n")
		router = setup_router()
	}
	if err := router.Run(":" + config.PORT); err != nil {
		log.Fatal(err)
	}
}

func setup_router() *gin.Engine {
	router := gin.Default()
	router.Delims("{{", "}}")
	token, _ := handlers.RandToken(64)
	store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(sessions.Sessions("authsession", store))

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(static.Serve("/assets", static.LocalFile(config.WORKDIR+"assets", false)))
	router.LoadHTMLGlob(config.WORKDIR + "templates/*.html")
	router.GET("/greetings", handlers.GetAllGreetings)
	router.POST("/greetings", handlers.AddGreeting)
	router.GET("/greetings/:id", handlers.GetGreetingById)
	router.PATCH("/greetings/:id", handlers.UpdateGreeting)
	router.DELETE("/greetings/:id", handlers.DeleteGreeting)

	router.GET("/", handlers.GetRootMd)
	router.GET("/:postName", handlers.GetMarkdown)
	return router
}

func setup_router_auth() *gin.Engine {
	// Creates a router without any middleware by default
	router := gin.New()
	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}
	store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(sessions.Sessions("authsession", store))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Delims("{{", "}}")
	router.LoadHTMLGlob(config.WORKDIR + "templates/*.html")
	router.Use(static.Serve("/assets", static.LocalFile(config.WORKDIR+"assets", false)))

	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)
	router.GET("/", handlers.GetRootMd)
	router.GET("/:postName", handlers.GetMarkdown)

	router.Use(middleware.AuthorizeRequest())

	// Authorization group
	// authorized := r.Group("/greetings", AuthRequired())
	// exactly the same as:
	authorized := router.Group("/greetings")

	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthorizeRequest())
	{
		router.GET("/greetings", handlers.GetAllGreetings)
		router.POST("/greetings", handlers.AddGreeting)
		router.GET("/greetings/:id", handlers.GetGreetingById)
		router.PATCH("/greetings/:id", handlers.UpdateGreeting)
		router.DELETE("/greetings/:id", handlers.DeleteGreeting)
	}

	return router
}

func declare_random_stuff() {
	// some common ways to declare variables
	// you wont get far without them
	var a = "a"
	log.Print(a)
	var b, c int = 1, 2
	log.Print(b, c)
	var d = true
	log.Print(d)
	var e int
	log.Print(e)
	h := "h"
	log.Print(h)

	// just like var, one may use const to declare entities
	const f rune = 'f'
	log.Print("my constant rune: ", f)
	const n = 500000000
	const j = 3e20 / n
	log.Print("my constant j: ", j)
	log.Print("my constant j type: ", reflect.TypeOf(j))

	// a number can be given a type by using it in a context that requires one
	// math.Sin expects a float64, n is implicitly cast
	log.Print("my constant n type: ", reflect.TypeOf(n))
	log.Print(math.Sin(n))
}
