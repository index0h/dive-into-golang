package main

import (
	"fmt"
	"time"
)

var (
	readFormat  = "RECEIVE %-8v: %v\n"
	writeFormat = "SEND    %-8v: %v\n"
)

// Sends string by 1 symbol to channel
func writeToChan(channel chan string, data, prefix string) {
	for _, symbol := range data {
		fmt.Printf(writeFormat, prefix, string(symbol))
		channel <- string(symbol)
	}
}

func main() {
	first := make(chan string)
	second := make(chan string)

	defer close(first)
	defer close(second)

	// Send strings to channels
	go writeToChan(first, "ABCDEF", "FRIST")
	go writeToChan(second, "abcdef", "SECOND")

	for {
		time.Sleep(50 * time.Millisecond)
		select {
		case firstValue := <-first:
			fmt.Printf(readFormat, "FRIST", firstValue)
		case secondValue := <-second:
			fmt.Printf(readFormat, "SECOND", secondValue)
//		default:
//			fmt.Println("DEFAULT")
//			return
		}
	}
}
