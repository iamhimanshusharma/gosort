package main

import (
	"fmt"

	"github.com/package/allsorting/s"
)

func main() {

	arr := []int{3, 2, 5, 7, 3}

	hold := s.SelectionSort(arr)

	fmt.Println(hold)
}
