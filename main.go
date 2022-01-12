package main

import (
	"fmt"

	"github.com/package/allsorting/s"
)

func main() {

	arr := []int{9, 5, 2, 56, 4, 3}
	hold := s.BubbleSort(arr)
	fmt.Println(hold)
}
