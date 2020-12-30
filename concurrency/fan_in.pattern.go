package concurrency

import (
	"fmt"
	"time"
)

var (
// s  = rand.NewSource(time.Now().Unix())
// r  = rand.New(s)
// wg sync.WaitGroup
)

func FaninPattern() {

	done1 := make(chan bool)
	done2 := make(chan bool)
	lc := genLetter(done1)
	nc := genNumber(done2)
	out := fanin(lc, nc)
	reader(out)
	time.Sleep(1 * time.Second)
	done1 <- true
	done2 <- true
	wg.Wait()
}

func reader(in chan string) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range in {
			fmt.Printf("Received Data: %v\n", val)
		}
	}()
}

func fanin(lc chan rune, nc chan int) (out chan string) {

	out = make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case l := <-lc:
				out <- fmt.Sprintf("%c", l)
			case n := <-nc:
				out <- fmt.Sprintf("%d", n)
			default:
				close(out)
				return
			}
		}
	}()
	return
}

func genLetter(done chan bool) (out chan rune) {

	wg.Add(1)
	out = make(chan rune)
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	go func() {
		defer wg.Done()
		for {
			select {
			case val := <-done:
				if val {
					close(out)
					return
				}
			case out <- letters[r.Int()%10]:
			}
		}
	}()
	return
}

func genNumber(done chan bool) (out chan int) {

	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case val := <-done:
				if val {
					close(out)
					return
				}
			case out <- r.Int() % 10:
			}
		}
	}()
	return
}
