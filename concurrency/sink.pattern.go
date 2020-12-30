package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// A function that launches goroutine to consume data on a channel

var (
	// s  = rand.NewSource(time.Now().Unix())
	// r  = rand.New(s)
	wg sync.WaitGroup
)

func SinkPattern() {

	done := make(chan bool)
	in := initGen1(done)
	counter(in)
	time.Sleep(1 * time.Second)
	done <- true
	wg.Wait()
}

func initGen1(done chan bool) (out chan int) {

	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-done:
				if v {
					close(out)
					return
				}
			case out <- (r.Int()%100 + 1):
			}
		}
	}()
	return
}

func counter(in chan int) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Counting process begin....")
		c := 0
		for range in {
			c++
		}
		fmt.Printf("Received data count: %d\n", c)
	}()
}
