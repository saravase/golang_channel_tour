package main

import "fmt"

func main() {

	var ch chan int
	fmt.Printf("ch: %v len: %d capacity: %d\n", ch, len(ch), cap(ch))

	// Sending data to channel
	// ch <- 5 // Can't sent data in nil channel

	// Receving data from channel
	// <-ch // Blocked receving from nil channel

	// Making channel without capacity (Unbuffer)
	ch = make(chan int)
	fmt.Printf("ch: %v len: %d capacity: %d\n", ch, len(ch), cap(ch))

	// Sending data to unbuffer channel(Without receiver)
	// ch <- 5 // Fails while try to send data without receiver

	// Receving data from unbuffer channel(Without sender)
	// <-ch // Fails while try to receive data without sender

	// Making channel with capacity (buffer)
	ch = make(chan int, 1)
	fmt.Printf("ch: %v len: %d capacity: %d\n", ch, len(ch), cap(ch))

	// Sending data to buffered channel
	ch <- 5
	fmt.Printf("ch: %v len: %d capacity: %d\n", ch, len(ch), cap(ch))

	// Receving data from buffered channel
	v := <-ch
	fmt.Printf("ch: %v len: %d capacity: %d value: %d\n", ch, len(ch), cap(ch), v)

	// Sending data to capacity filled channel
	ch <- 5
	// ch <- 6 // Fails because the channel was reached maximum capacity

	// Making channel with more capacity (buffer)
	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("ch: %v len: %d capacity: %d\n", ch, len(ch), cap(ch))

	// Receving data from empty channel
	fmt.Printf("1 value from ch channel: %d length : %d\n", <-ch, len(ch))
	fmt.Printf("2 value from ch channel: %d length : %d\n", <-ch, len(ch))
	fmt.Printf("3 value from ch channel: %d length : %d\n", <-ch, len(ch))
	// fmt.Printf("4 value from ch channel: %d length : %d\n", <-ch, len(ch)) // Fails Because the channel is empty

	// Receive-only and send-only channels
	chs := make(chan string, 3)
	// Code that only produces data to channel
	var out chan<- string
	out = chs
	out <- "optimus"
	out <- "primz"
	// <-out // Compilation fails. Because out is send only type of variable

	// Code that only consumes data to channel
	var in <-chan string
	in = chs
	fmt.Printf("1 value from chs channel: %s length : %d\n", <-in, len(chs))
	fmt.Printf("2 value from chs channel: %s length : %d\n", <-in, len(chs))

}
