package searching

// BinSearch description
func BinSearch(sl *[]int, el int) (bool, int) {
	var med int
	begin, end := 0, len(*sl) - 1
	for begin <= end {
		med = (begin + end) / 2
		switch {
		case el == (*sl)[med]:
			return true, med
		case (*sl)[med] < el:
			begin = med + 1
		case (*sl)[med] > el:
			end = med - 1
		}
	}
	return false, -1
}