package golang_channel_tour

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	s           = rand.NewSource(time.Now().Unix())
	r           = rand.New(s)
	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
)

func GoroutineChannel() {

	// ch := make(chan string, 5)
	// producer(1, ch)
	// consumer(ch)

	// concurrent data producer and consumer
	// ch := make(chan string, 5)
	// go producer(1, ch)
	// consumer(ch)

	// 'Unbuffered channel' concurrent data producer and consumer
	// ch := make(chan string)
	// go producer(1, ch)
	// consumer(ch)

	// 'Signalling with Unbuffered channel' concurrent data producer and consumer
	// ch := make(chan string)
	// go producer1(1, ch)
	// consumer(ch)

	// Multiple concurrent data producers and single consumer
	ch := make(chan string)
	producer2(1, ch)
	producer2(2, ch)
	producer2(3, ch)
	consumer1(1, ch) // Unbuffered channel consumer must read the data parallelly
	consumer1(2, ch)
	consumer1(3, ch)
	wgProducers.Wait()
	close(ch)
	wgConsumers.Wait() // Here consumer is goroutine. So that consumer only consumer waitgroup added
}

func producer(id int, in chan string) {
	n := r.Int() % 5
	for m := 0; m < n; m++ {
		in <- fmt.Sprintf("Producer #%d, Message #%d", id, r.Int()%1000+1)
	}
	close(in)
}

func producer1(id int, in chan string) {
	end := time.Now().Add(1000 * time.Millisecond)
	for time.Now().Before(end) {
		in <- fmt.Sprintf("Producer #%d, Message #%d", id, r.Int()%1000+1)
	}
	close(in)
}

func producer2(id int, in chan string) {
	wgProducers.Add(1)
	go func() {
		end := time.Now().Add(1000 * time.Millisecond)
		for time.Now().Before(end) {
			in <- fmt.Sprintf("Producer #%d, Message #%d", id, r.Int()%1000+1)
		}
		wgProducers.Done()
	}()
}

func consumer(out chan string) {
	c := 0
	for val := range out {
		c++
		fmt.Println("Consumer receive: %s", val)
	}

	if c == 0 {
		fmt.Println("No data received")
	}

	fmt.Printf("Total no of messages received from producers: %d\n", c)

}

func consumer1(id int, out chan string) {
	wgConsumers.Add(1)
	go func() {
		dataMap := make(map[string]int)
		fields := []string{}
		for val := range out {
			fields = strings.Split(val, ",")
			dataMap[fields[0]]++
		}
		for k, v := range dataMap {
			fmt.Printf("Total messages receiver #%d received from %s : %d\n", id, k, v)
		}
		wgConsumers.Done()
	}()
}
