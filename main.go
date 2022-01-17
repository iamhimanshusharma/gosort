package main

import (
	"fmt"
	"sorting/s"
)

func main() {

	arr := []int{5, 3, 6, 7, 8}
	shift := s.InsertionSort(arr)
	fmt.Println(shift)
}
