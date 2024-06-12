package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	doneChan := make(chan int)

	go func() {
		defer close(doneChan)

		select {

		case <-stopChan:
			return

		case i := <-firstChan:
			doneChan <- i * i
		case i := <-secondChan:
			doneChan <- i * 3
		}

	}()
	return doneChan
}

func main() {

	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	r := calculator(firstChan, secondChan, stopChan)
	//firstChan <- 2
	stopChan <- struct{}{}
	calculator(firstChan, secondChan, stopChan)
	fmt.Println(<-r)
}
