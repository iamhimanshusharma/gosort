package s

func InsertionSort(listToSort []int) []int {
	emptyList := []int{}
	for i := 0; i < len(listToSort)-1; i++ {
		for j := 0; j < i; j++ {
			if listToSort[i] < listToSort[j] {
				emptyList = finalLister(listToSort, i, j)
			}
		}
	}
	return emptyList
}

func Deleter(listToInsert []int, placeFrom int) []int { //Deletes element from the slice
	prefixList := listToInsert[0:placeFrom]
	shiftList := listToInsert[placeFrom+1:]
	newlist := []int{}
	newlist = append(newlist, prefixList...)
	newlist = append(newlist, shiftList...)
	return newlist
}

func Inserter(listToInsert []int, placeFrom int, placeTo int) []int { //Adds the element to the slice
	prefixList := listToInsert[0:placeTo]
	shiftList := listToInsert[placeTo:]
	newlist := []int{}
	newlist = append(newlist, prefixList...)
	newlist = append(newlist, listToInsert[placeFrom])
	newlist = append(newlist, shiftList...)
	return newlist
}

func finalLister(listToInsert []int, placeFrom int, placeTo int) []int { //returns final list after positioning the element at right place
	insertedList := Inserter(listToInsert, placeFrom, placeTo)
	finalList := Deleter(insertedList, placeFrom+1)
	return finalList
}

// Features need to add
// If passed soreted list then returns empty list : Need to fix
