package main

import (
	"fmt"
	"time"
)

var (
	readFormat  = "READ  %-8v: %v\n"
	writeFormat = "WRITE %-8v: %v\n"
)

// Writes to stdout read data from channel
func readFromChan(channel <-chan string, prefix string) {
	// channel <- "It'll not work"

	for data := range channel {
		fmt.Printf(readFormat, prefix, data)
	}
}

// Sends string by 1 symbol to channel
func writeToChan(channel chan<- string, data, prefix string) {
	// <-channel // "It'll not work"

	for _, symbol := range data {
		fmt.Printf(writeFormat, prefix, string(symbol))
		channel <- string(symbol)
	}
}

func main() {
	sendData := ""
	unBufferedChannel := make(chan string)
	bufferedChannel := make(chan string, 10)

	defer close(bufferedChannel)
	defer close(unBufferedChannel)

	// channel := unBufferedChannel // fatal error: all goroutines are asleep - deadlock!
	channel := bufferedChannel
	sendData = "FirstData"
	fmt.Printf(writeFormat, "First", sendData)

	channel <- sendData

	fmt.Printf(readFormat, "First", <-channel)
	fmt.Println("\n------------------------------------------------------------------\n")

	// Unbuffered channel in goroutines

	go readFromChan(unBufferedChannel, "Second")

	writeToChan(unBufferedChannel, "abcdef", "Second")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("\n------------------------------------------------------------------\n")

	// buffered channel in goroutines

	go readFromChan(bufferedChannel, "Third")

	writeToChan(bufferedChannel, "abcdef", "Third")

	time.Sleep(100 * time.Millisecond)
}
