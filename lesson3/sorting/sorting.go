package sorting

// ShakerSorting description
// Реализация по алгоритму без оптимизаций
func ShakerSorting(sl *[]int) (counter int) {
	var startIndx int
	var endIndx = len(*sl) - 1

	for startIndx < endIndx {
		for i:=startIndx; i < endIndx; i++ {
			counter++
			if (*sl)[i] > (*sl)[i+1] {
				counter++
				(*sl)[i], (*sl)[i+1] = (*sl)[i+1], (*sl)[i]
			}
		}
		endIndx--

		for i:= endIndx; i > startIndx;i-- {
			counter++
			if (*sl)[i] < (*sl)[i-1] {
				counter++
				(*sl)[i], (*sl)[i-1] = (*sl)[i-1], (*sl)[i]
			}
		}
		startIndx++
	}
	return
}

// ShakerSortingFlag description (for linter)
// Для оптимизации введем флаг swappedFlag, если в проход нет перестановок, выходим из цикла
func ShakerSortingFlag(sl *[]int) (counter int) {
	var swappedFlag = true  // флаг проверки наличия перестановок
	var startIndx int  // default - 0
	var endIndx = len(*sl) - 1

	for swappedFlag {
		swappedFlag = false
		for i := startIndx; i < endIndx; i++ {
			counter++
			if (*sl)[i] > (*sl)[i+1] {
				counter++
				(*sl)[i], (*sl)[i+1] = (*sl)[i+1], (*sl)[i]
				swappedFlag = true
			}
		}

		if !swappedFlag {break}
		swappedFlag = false
		endIndx--

		for i:= endIndx; i > startIndx;i-- {
			counter++
			if (*sl)[i-1] > (*sl)[i] {
				counter++
				(*sl)[i], (*sl)[i-1] = (*sl)[i-1], (*sl)[i]
				swappedFlag = true
			}
		}
		startIndx++
	}
	return
}

// ShakerSortingMemPos description
// Смещаться не на 1 каждый раз, а на индекс последней перестановки
func ShakerSortingMemPos(sl *[]int) (counter int) {
	var swapped = true
	var startIndx int; var startSwap int
	var endIndx, endSwap = len(*sl) - 1, len(*sl) - 1

	for swapped {
		swapped = false
		for i:=startIndx; i < endIndx; i++ {
			counter++
			if (*sl)[i] > (*sl)[i+1] {
				counter++
				(*sl)[i], (*sl)[i+1] = (*sl)[i+1], (*sl)[i]
				swapped = true
				endSwap = i
			}
		}
		if !swapped {
			break
		} else {
			swapped = false
			endIndx = endSwap
		}

		for i:= endIndx; i > startIndx;i-- {
			counter++
			if (*sl)[i] < (*sl)[i-1] {
				counter++
				(*sl)[i], (*sl)[i-1] = (*sl)[i-1], (*sl)[i]
				swapped = true
				startSwap = i
			}
		}
		startIndx = startSwap
	}
	return
}
