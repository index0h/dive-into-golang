package main

import (
	"fmt"
	"time"
)
// START OMIT

func Ticker(channel <-chan time.Time) {
	for _ = range channel {
		fmt.Println("tick")
	}
}

func main() {
	stopChannel := time.After(5 * time.Second)
	tickChannel := time.Tick(time.Second)

	go Ticker(tickChannel)

	<-stopChannel
	fmt.Println("Bye")
}
// END OMIT
