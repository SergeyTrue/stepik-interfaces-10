package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		var num int
		for {
			select {
			case i := <-arguments:

				num += i
			case <-done:
				out <- num
				return
			}
		}
	}()
	return out
}

func main() {
	ch1 := make(chan int, 1)
	stop := make(chan struct{})
	resultChan := calculator(ch1, stop)
	go func() {
		ch1 <- 1
		ch1 <- 3
		close(stop)
	}()

	fmt.Println(<-resultChan)
}
