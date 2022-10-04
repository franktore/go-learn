package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"

	"github.com/franktore/go-learn/pkg/greetings"
)

const log_prefix string = "greetings: "

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

	message, err := greetings.Hello(name)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}

func declare_random_stuff() {
	fmt.Println("init")
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
