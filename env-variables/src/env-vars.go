package main

import (
	"fmt"
	"os"
)

func main() {
	// This looping show you all environment variables
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
