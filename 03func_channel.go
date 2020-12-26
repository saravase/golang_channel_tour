package golang_channel_tour

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CH_CAP = 10
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// passing channel to function
	// d := make(chan int, CH_CAP)
	// if d == nil {
	// 	log.Fatal("Chennal creation failed")
	// }
	// producer(d)
	// consumer(d)

	// returning channels from functions
	// d := generator()
	// consumer(d)

	// processing pipeline - v1
	// d := generator()
	// d = counter(d)
	// consumer(d)

	// processing pipeline - v2
	d := generator()
	d = counter(d)
	d = adder(d, 5)
	consumer(d)
}

func adder(in chan int, a int) (out chan int) {
	out = make(chan int, len(in))
	for val := range in {
		out <- val + a
	}
	close(out)
	return
}

func counter(in chan int) (out chan int) {
	out = make(chan int, len(in))
	count := 0
	for val := range in {
		out <- val
		count++
	}
	close(out)
	fmt.Printf("Counted %v elements\n", count)
	return
}

func generator() (out chan int) {
	out = make(chan int, CH_CAP)
	producer(out)
	return
}

func producer(ch chan int) {
	n := r.Int()%cap(ch) + 1
	for i := 0; i < n; i++ {
		ch <- r.Int() % 200
	}
	close(ch)
}

func consumer(ch chan int) {
	for n := range ch {
		fmt.Printf("Received Value : %d\n", n)
	}
}
