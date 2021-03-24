package main

import (
	"fmt"
)

func myfunc(mychnl chan string) {
	for v := 0; v < 4; v++ {
		mychnl <- "Func String"
	}
	close(mychnl)
}

func main() {

	//create channel
	c := make(chan string)

	go myfunc(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close", ok)
			break
		}
		fmt.Println("Channel Open", res, ok)
	}
}
