package main

import (
	"fmt"
)

var GitCommit, GitState, Version string

func main() {
	fmt.Println("hi")
	fmt.Printf("%v, %v, %v", GitCommit, GitState, Version)
}
