package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 2)
	go worker(done)

	fmt.Println("waiting for the worker to finish")

	<-done
	go worker(done)
	fmt.Println("worker finished, exiting")

	<-done
}
