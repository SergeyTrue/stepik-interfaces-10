package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	Merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func Merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	results1 := make([]int, n)
	results2 := make([]int, n)
	innerWg := sync.WaitGroup{}
	innerWg.Add(n * 2)
	processChan := func(ch <-chan int, results []int) {

		for i := 0; i < n; i++ {
			val := <-ch
			go func(val, index int) {
				defer innerWg.Done()
				results[i] = fn(val)
			}(val, i)
		}
		innerWg.Wait()
	}

	go processChan(in1, results1)
	go processChan(in2, results2)

	go func() {
		innerWg.Wait()
		for i := 0; i < n; i++ {
			out <- results1[i] + results2[i]
		}
		close(out)
	}()
}
