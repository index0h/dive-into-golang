package main

import (
	"fmt"
	"time"
	"sync"
	"strconv"
)

// Current function works 1 second
func Work(name string) {
	fmt.Printf("Work: %-5v start\n", name)
	time.Sleep(time.Second)
	fmt.Printf("Work: %-5v end\n", name)
}

// Simple demonstration of goroutine
func Sleep() {

	go Work("simple")

	time.Sleep(2 * time.Second)
}

// Demonstration of sync.WaitGroup
func WaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			Work("wg-" + strconv.Itoa(id))
		}(i)
	}

	wg.Wait()
}

func main() {
	fmt.Println("main: start")

	Sleep()

	WaitGroup()

	fmt.Println("main: end")
}
