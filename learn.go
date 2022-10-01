package main

import (
	"fmt"
	"os"

	"source/go-learn/greetings"
)

func main() {
	argsWithoutProg := os.Args[1:]
	name := "T"
	if len(argsWithoutProg) > 0 {
		name = argsWithoutProg[0]
	}

	message := greetings.Hello(name)
	fmt.Println(message)
}
