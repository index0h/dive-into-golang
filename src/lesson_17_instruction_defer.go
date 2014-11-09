package main

import (
	"fmt"
	"time"
)

func first() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Call: first")
}

func second() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Call: second")
}

func third() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Call: third")
}

func fourth() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Call: fourth")
}

func fifth() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Call: fifth")
}


func main() {
	fmt.Println("Call: main - begin")

	defer first()
	defer second()
	defer third()
	defer fourth()
	defer fifth()

	fmt.Println("Call: main - end")
}
