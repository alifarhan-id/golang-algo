package main

import (
	"fmt"

	"github.com/alifarhan1230/golang-algo/sorting/bubble_sort"
)

func main() {
	fmt.Println("welcome to data structure with go, your here with Ali Farhan")
	number := [8]int{101, 8, 4, 2, 10, 87, 98, 100}
	fmt.Println(bubble_sort.BubbleSort(number))

}
