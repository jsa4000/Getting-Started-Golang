package main

import (
	"context"
	"fmt"
)

// SendRequest to manage a request ignoring the context
// _ is used so the compiler does not complain about the non usage of the variable
func SendRequest(ctx context.Context, arg interface{}) error {
	return nil
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {
	printHeader("Context")

	ctx = context.Background()

	go SendRequest(c)
}
