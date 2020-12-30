package concurrency

import (
	"fmt"
	"time"
)

// A function that launches goroutine to produce data on a channel it returns

var (
// s  = rand.NewSource(time.Now().Unix())
// r  = rand.New(s)
// wg sync.WaitGroup
)

func FanoutPattern() {

	done := make(chan bool)
	in := initGen3(done)
	c1, c2, c3 := fanout(in)
	counter2(c1)
	counter2(c2)
	counter2(c3)
	time.Sleep(1 * time.Second)
	done <- true
	wg.Wait()
}

func initGen3(done chan bool) (out chan int) {

	out = make(chan int)
	go func() {
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

func fanout(in chan int) (out1 chan int, out2 chan int, out3 chan int) {

	wg.Add(1)
	out1 = make(chan int)
	out2 = make(chan int)
	out3 = make(chan int)
	go func() {
		defer wg.Done()
		for val := range in {
			select {
			case out1 <- val:
			case out2 <- val:
			case out3 <- val:
			}
		}
		close(out1)
		close(out2)
		close(out3)
	}()
	return
}

func counter2(in chan int) {

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
