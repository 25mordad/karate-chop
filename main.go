package main

import (
	"cycloid-challenge/midsearch"
	"cycloid-challenge/recursive"
	"fmt"
)

///
func main() {
	items := []int{1, 2, 20, 31, 45, 55, 63, 70, 100}

	fmt.Println("START midsearch")
	fmt.Println("Items: ", items)
	fmt.Println("FIND: ", midsearch.Chop(100, items))

	fmt.Println("START recursive")
	fmt.Println("Items: ", items)
	fmt.Println("FIND: ", recursive.Chop(100, items))

}
