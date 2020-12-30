package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// A function that launches goroutine to produce data on a channel it returns

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func GeneratePattern() {

	done := make(chan bool)
	in := initGen(done)

	for i := 0; i < 10; i++ {
		fmt.Printf("Received Data : %d\n", <-in)
	}
	done <- true
}

func initGen(done chan bool) (out chan int) {
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
