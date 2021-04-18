package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	mystring := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(mystring))

	bs := h.Sum(nil)

	fmt.Println(mystring)
	fmt.Printf("%x\n", bs)
}
