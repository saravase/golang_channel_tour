package golang_channel_tour

import (
	"sync"
	"time"

	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	BITS_LEN = 100000
)

var (
	wgProducer sync.WaitGroup
	wgConsumer sync.WaitGroup
)

func SelectionChannel() {
	// fmt.Printf("Message 1 at : %v\n ", time.Now())
	// sleep(1 * time.Second)
	// fmt.Printf("Message 2 at : %v\n ", time.Now())

	// fmt.Printf("Message 1 at : %v\n ", time.Now())
	// time.Sleep(1 * time.Second)
	// fmt.Printf("Message 2 at : %v\n ", time.Now())

	// fmt.Printf("Message 1 at : %v\n ", time.Now())
	// alarm := notifyAfter(1 * time.Second)
	// <-alarm
	// fmt.Printf("Message 2 at : %v\n ", time.Now())

	// fmt.Printf("Message 1 at : %v\n ", time.Now())
	// <-time.After(1 * time.Second)
	// fmt.Printf("Message 2 at : %v\n ", time.Now())

	// var ch chan int
	// select {
	// case <-ch: // ch is nil
	// 	log.Info("Data received from channel ch")
	// default:
	// 	log.Warn("No data received from channel ch") // run
	// }

	// var ch1, ch2 chan int
	// select {
	// case <-ch1: // ch1 is nil
	// 	log.Info("Data received from channel ch1")
	// case ch2 <- 10: // ch2 is nil
	// 	log.Info("Data send from channel ch2")
	// default:
	// 	log.Warn("No communication between ch1 & ch2") // run
	// }

	// Not suitable for million or billion bits generation
	// bits := generateBits(10)
	// fmt.Print("Random bit stream : ")
	// for bit := range bits {
	// 	fmt.Print(bit)
	// }
	// fmt.Println()

	// Best way
	// bits := generateBits1(100000)
	// fmt.Print("Random bit stream : ")
	// for bit := range bits {
	// 	fmt.Print(bit)
	// }
	// fmt.Println()

	// bitsMap := make(map[int8]int)
	// bits := generateBits1(BITS_LEN)
	// fmt.Println("Random bit stream : ")
	// for bit := range bits {
	// 	bitsMap[bit]++
	// }
	// for k, v := range bitsMap {
	// 	fmt.Printf("bit %d occupied %.2f%% \n", k, (float32(v)/BITS_LEN)*100)
	// }

	in := producer()
	consumer(in)
	wgConsumer.Wait()
}

func consumer(in chan int) {
	wgConsumer.Add(1)
	go func() {
		defer wgConsumer.Done()
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				fmt.Println(v)
			case <-time.After(3 * time.Millisecond):
				fmt.Println("Stopped")
				return
			}
		}
	}()
}

func producer() (out chan int) {
	out = make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}

		time.Sleep(2 * time.Millisecond)
		for i := 10; i < 20; i++ {
			out <- i
		}
		close(out)

	}()
	return
}

func generateBits1(l int) (out chan int8) {
	out = make(chan int8)
	go func() {
		for m := 0; m < l; m++ {
			select {
			case out <- 0:
			case out <- 1:
			case out <- 1:
			case out <- 1:
			}
		}
		close(out)
	}()
	return
}

func generateBits(l int) (out chan int8) {
	out = make(chan int8, l)
	defer close(out)

	for {
		select {
		case out <- 0:
		case out <- 1:
		default:
			return
		}
	}
}

func notifyAfter(delay time.Duration) (out chan time.Time) {
	out = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		out <- time.Now()
		close(out)
	}()
	return out
}

func sleep(delay time.Duration) {
	end := time.Now().Add(delay)
	for time.Now().Before(end) {
		log.Info("Typing......")
	}
}
