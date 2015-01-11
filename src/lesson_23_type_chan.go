package main

import (
	"fmt"
	"time"
)

var (
	readFormat  = "RECEIVE %-8v: %v\n"
	writeFormat = "SEND    %-8v: %v\n"
)

// Receives data from channel
func receive(channel <-chan string, prefix string) {
	// channel <- "It'll not work"

	for data := range channel {
		fmt.Printf(readFormat, prefix, data)
	}
}

// Sends string by 1 symbol to channel
func send(channel chan <- string, data, prefix string) {
	// <-channel // "It'll not work"

	for _, symbol := range data {
		fmt.Printf(writeFormat, prefix, string(symbol))
		channel <- string(symbol)
	}
}

func main() {
	var unBufferedChannel chan string
	unBufferedChannel = make(chan string)
	bufferedChannel := make(chan string, 10)

	defer close(bufferedChannel)
	defer close(unBufferedChannel)

	// channel := unBufferedChannel // fatal error: all goroutines are asleep - deadlock!
	channel := bufferedChannel
	// defer close(channel) // panic: runtime error: close of closed channel
	fmt.Printf(writeFormat, "First", "FirstData")

	channel <- "FirstData"

	fmt.Printf(readFormat, "First", <-channel)
	fmt.Println("\n------------------------------------------------------------------\n")

	// Unbuffered channel in goroutines

	go receive(unBufferedChannel, "Second")

	send(unBufferedChannel, "abcdef", "Second")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("\n------------------------------------------------------------------\n")

	// buffered channel in goroutines

	go receive(bufferedChannel, "Third")

	send(bufferedChannel, "abcdef", "Third")

	time.Sleep(100 * time.Millisecond)
}
