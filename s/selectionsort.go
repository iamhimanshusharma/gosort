package s

func SelectionSort(listToSort []int) []int {
	for i := 0; i < len(listToSort)-1; i++ {
		minIndex := i
		for j := i; j < len(listToSort); j++ {
			if listToSort[minIndex] > listToSort[j] {
				minIndex = j
			}
		}
		temp := listToSort[minIndex]
		listToSort[minIndex] = listToSort[i]
		listToSort[i] = temp
	}
	return listToSort
}
