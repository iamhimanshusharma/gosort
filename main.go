package main

import (
	"fmt"

	"github.com/go/sort/s"
)

func main() {

	arr := []int{9, 5, 2, 56, 4, 3}
	hold := s.BubbleSort(arr)
	fmt.Println(hold)
}
