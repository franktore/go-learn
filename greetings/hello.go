package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Println(message)

	ret := quote.Go()
	return ret
}
