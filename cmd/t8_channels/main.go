package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	var c = make(chan string) // Channel: Channels are used to communicate between goroutines
	// c <- "Hello"              // Send data to the channel
	// The above line will cause a deadlock as there is no goroutine to receive the data from the channel
	// To fix this, we need to run the above line in a goroutine
	// go process(c)
	// fmt.Println("Received data from channel: ", <-c) // Receive data from the channel

	go processLoop(c)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Received data from channel: ", i) // Receive data from the channel
	// }
	// The above approach isn't effective as we don't know how many times the data will be sent to the channel

	// To fix this, we use range to iterate over the channel
	fmt.Println("Before listening 1")
	for msg := range c {
		fmt.Println("Received data from channel: ", msg) // Receive data from the channel
		// time.Sleep(2 * time.Second)
	}
	fmt.Println("After listening 1") // This line will be blocked until the channel is closed

	// Buffer channels: Channels can be buffered to store multiple values
	cb := make(chan string, 5) // Create a buffered channel with a capacity of 5
	// With above approach, we can send 5 values to the channel without being blocked by the receiver
	go processLoop(cb)
	fmt.Println("Before listening 2")
	for msg := range cb {
		fmt.Println("Received data from channel: ", msg) // Receive data from the channel
		time.Sleep(1 * time.Second)
	}
	fmt.Println("After listening 2") // This line will be blocked until the channel is closed


}

func process(c chan string) {
	c <- "Hello" // Send data to the channel
}

func processLoop(c chan string) {
	defer close(c) // Close the channel after writing all the data is necessary to avoid deadlock
	for i := 0; i < 5; i++ {
		// time.Sleep(1 * time.Second)
		fmt.Println("Before writing: ", i)
		c <- fmt.Sprintf("Hello %d", i) // Send data to the channel
		fmt.Println("After writing: ", i)
	}
	// Close the channel after writing all the data is necessary to avoid deadlock
	// close(c) // Close the channel - Above defer statement can also be used and is recommended
}

func listen(c chan string) {
	fmt.Println("Received data from channel: ", <-c) // Receive data from the channel
	wg.Done()
}
