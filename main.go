package main

import (
	"fmt"
)

func main() {
	g := getGreeting()
	fmt.Println(g)
}

func getGreeting() string {
	return "Hello world!"
}