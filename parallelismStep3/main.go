package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(5 * time.Second)
}

func main() {
	fmt.Println("Main starting")
	done := make(chan (struct{}))
	go func() {
		work()
		close(done)
	}()

	<-done
	fmt.Println("Main end")
}
