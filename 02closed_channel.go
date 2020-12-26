// Place 02 - Working with closed channels

package golang_channel_tour

import "fmt"

var (
	ch chan int
)

func generateChData(c, l int) {
	ch = make(chan int, c)
	for i := 0; i < l; i++ {
		ch <- i
	}
}

func ClosedChannel() {
	// closing a channel
	sc := make(chan string, 2)
	sc <- "hello"
	close(sc)
	fmt.Printf("sc: %v, len: %d, cap: %v\n", sc, len(sc), cap(sc))

	// Sending data to a closed channel
	// sc <- "optimus" // Fails, Because the channel state closed

	// Receving from closed channel
	sc = make(chan string, 2)
	sc <- "primz"
	close(sc)
	s := <-sc
	fmt.Printf("sc 1st value: %v, len: %d, cap: %v\n", s, len(sc), cap(sc))
	s = <-sc // Once channel was closed. It is always rececing zero value from closed channel
	fmt.Printf("sc 2nd value: %v, len: %d, cap: %v\n", s, len(sc), cap(sc))

	// Testing for value send from closed channel or not
	sc = make(chan string, 2)
	sc <- "optimus"
	sc <- ""
	close(sc)
	var ok bool
	s, ok = <-sc
	fmt.Printf("sc 1st value: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("sc 2nd value: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("sc 3rd value: %v, ok: %v\n", s, ok)

	// Test reading data from an empty channel without closing
	sb := make(chan bool, 3)
	sb <- true
	c, ok := <-sb
	fmt.Printf("sb 1st value: %v, ok: %v\n", c, ok)
	// c, ok = <-sb
	// fmt.Printf("sb 2nd value: %v, ok: %v\n", c, ok) // Fails, Because sb channel not have data, it is empty

	// Incorrect way of iterating channel values
	// generateChData(5, 5)
	// for i := 0; i < cap(ch); i++ {
	//   fmt.Println(<-ch)
	// }

	// generateChData(5, 3)
	// for i := 0; i < cap(ch); i++ {
	// 	fmt.Println(<-ch)
	// }

	// generateChData(5, 5)
	// for i := 0; i < len(ch); i++ {
	//	fmt.Println(<-ch)
	// }
	// Output : 0 1 2 Because, [i < len(ch)] loop condition affect the iteration
	// len(ch) value was updated each iteration

	// generateChData(5, 5)
	// l := len(ch)
	// for i := 0; i < l; i++ {
	// 	fmt.Println(<-ch)
	// }
	// Not suitable Because, if len(ch) was updatable after initialization process

	// generateChData(5, 5)
	// for val := range ch {
	// 	fmt.Println(val)
	// } // Fails, Because range loop try to get the new data from channel

	generateChData(5, 1)
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}

}
