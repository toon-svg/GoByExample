package main

import (
	"fmt"
	"sort"
)

func main() {

	mystrings := []string{"c", "a", "b"}
	sort.Strings(mystrings)
	fmt.Println("Strings:", mystrings)

	myintegers := []int{7, 2, 4}
	sort.Ints(myintegers)
	fmt.Println("Ints:  ", myintegers)

	s := sort.IntsAreSorted(myintegers)
	fmt.Println("Sorted: ", s)
}
