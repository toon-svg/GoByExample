package main

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(234 + <-ch)
}

func main() {
	fmt.Println("start Main function")
	//create channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End main function")
}
