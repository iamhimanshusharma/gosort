package s

func BubbleSort(listToSort []int) []int {

	for i := 0; i < len(listToSort)-1; i++ {
		for j := 0; j < len(listToSort)-1; j++ {
			fIndex := j
			lIndex := j + 1
			if listToSort[fIndex] > listToSort[lIndex] {
				temp := listToSort[fIndex]
				listToSort[fIndex] = listToSort[lIndex]
				listToSort[lIndex] = temp
			}
		}
	}
	return listToSort
}
