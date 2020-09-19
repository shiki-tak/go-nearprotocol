package main

import (
	"fmt"

	near "github.com/shiki-tak/go-nearprotocol"
)

func main() {
	near, err := near.NewWasmer()
	if err != nil {
		fmt.Println(err)
	}
	near.Greet("Go-NearProtocol")
}
