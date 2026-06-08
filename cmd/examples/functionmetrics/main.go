package main

import (
	"fmt"
	"time"
)

func track(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
    defer track("main")()
    fmt.Println("hello world!")
}

