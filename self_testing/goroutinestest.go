package main

import (
	"fmt"
	"time"
)

func three() {
	fmt.Println("Three")
}

func main() {
	fmt.Println("One")
	fmt.Println("Two")
	go three()
	time.Sleep(1 * time.Second)
	fmt.Println("Four")
}
