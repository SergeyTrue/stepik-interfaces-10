package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fn(i int) int {
	time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)
	return i * i
}

func Merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

}
func main() {

	chanToSlice := func(inputCh chan int, n int) []int {
		wg := sync.WaitGroup{}
		wg.Add(3)
		mu := sync.Mutex{}
		resSlice := make([]int, n)
		type indexVal struct {
			index int
			val   int
		}
		indChan := make(chan indexVal)
		calcChan := make(chan indexVal)
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				indChan <- indexVal{
					index: i,
					val:   <-inputCh,
				}
			}
			close(indChan)
		}()

		go func() {
			defer wg.Done()
			calcWG := sync.WaitGroup{}
			for item := range indChan {
				calcWG.Add(1)
				go func(item indexVal) {
					defer calcWG.Done()
					res := fn(item.val)
					calcChan <- indexVal{
						index: item.index,
						val:   res,
					}
				}(item)
			}
			calcWG.Wait()
			close(calcChan)
		}()

		go func() {
			defer wg.Done()
			for item := range calcChan {
				mu.Lock()
				resSlice[item.index] = item.val
				mu.Unlock()
			}
		}()
		wg.Wait()
		return resSlice

	}
	firstSl := []int{9, 1, 3, 5, 7}
	secondSl := []int{90, 20, 30, 40, 50}
	n := 5

	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)
	outerWg := sync.WaitGroup{}
	outerWg.Add(1)
	go func() {
		defer outerWg.Done()
		for _, val := range firstSl {
			in1 <- val
		}
		close(in1)
		for _, val := range secondSl {
			in2 <- val
		}
		close(in2)

	}()

	middleWg := sync.WaitGroup{}
	var slSl = make([][]int, 2)
	chSl := []chan int{in1, in2}

	middleWg.Add(len(chSl))
	for i, ch := range chSl {
		go func() {
			defer middleWg.Done()
			slSl[i] = chanToSlice(ch, n)
		}()
	}
	middleWg.Wait()

	lastWg := sync.WaitGroup{}
	lastWg.Add(2)
	go func() {
		defer lastWg.Done()
		for i := 0; i < n; i++ {
			out <- (slSl[0][i] + slSl[1][i])
		}
		close(out)
	}()

	go func() {
		defer lastWg.Done()
		for val := range out {
			fmt.Println(val)
		}
	}()
	lastWg.Wait()
	//вот так работает
	outerWg.Wait()
	//fmt.Println(slSl[0][0] + slSl[1][0])
	//fmt.Println(slSl[0][1] + slSl[1][1])
	//fmt.Println(slSl[0][2] + slSl[1][2])
	//fmt.Println(slSl[0][2] + slSl[1][2])
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//fmt.Println(<-out)

	//fmt.Println(time.Since(start))
}
