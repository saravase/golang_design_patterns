package concurrency

import (
	"fmt"
	"time"
)

// A function that launches a goroutine to produce & consume data

var (
// s  = rand.NewSource(time.Now().Unix())
// r  = rand.New(s)
// wg sync.WaitGroup
)

func ProcessorPattern() {

	done := make(chan bool)
	in := initGen2(done)
	in = throughCounter(in)
	in = filter(in, func(val int) bool {
		if val%2 == 0 {
			return true
		}
		return false
	})
	counter1(in)
	time.Sleep(1 * time.Second)
	done <- true
	wg.Wait()
}

func filter(in chan int, cmp func(int) bool) (out chan int) {

	wg.Add(1)
	out = make(chan int)

	if cmp == nil {
		cmp = func(int) bool {
			return true
		}
	}

	go func() {
		defer wg.Done()
		for val := range in {
			if cmp(val) {
				out <- val
			}
		}
		close(out)
	}()
	return
}

func initGen2(done chan bool) (out chan int) {

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

func counter1(in chan int) {

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

func throughCounter(in chan int) (out chan int) {

	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
		fmt.Println("ThroughCounter counting process begin....")
		c := 0
		for val := range in {
			c++
			out <- val
		}
		fmt.Printf("ThroughCounter received data count: %d\n", c)
		close(out)
	}()
	return
}
