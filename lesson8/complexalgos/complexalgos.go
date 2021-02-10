package complexalgos

func swap(ar *[]int, i int, j int) {
	if (*ar)[i] > (*ar)[j] {
		(*ar)[i], (*ar)[j] = (*ar)[j], (*ar)[i]
	}
}

// QuickSort desc
func QuickSort(nums *[]int, begin int, end int, counter *int) {
	if begin > end {
		return
	}
	i, j, pivot := begin, end, (*nums)[(begin+end)/2]
	for i <= j {
		for (*nums)[i] < pivot {
			*counter++
			i++
		}
		for (*nums)[j] > pivot {
			*counter++
			j--
		}
		if i <= j {
			swap(nums, i, j)
			i++
			j--
		}
	}
	if i < end {
		QuickSort(nums, i, end, counter)
	}
	if begin < j {
		QuickSort(nums, begin, j, counter)
	}
}


// MergeSort descr
func MergeSort(sl *[]int, counter *int) *[]int {
	if len(*sl) < 2 {
		return sl
	}
	mid := len(*sl) /2
	lRec, rRec := (*sl)[:mid], (*sl)[mid:]
	l, r := MergeSort(&lRec, counter), MergeSort(&rRec, counter)
	return merge(l,r, counter)
}

func merge(l,r *[]int, counter *int) *[]int {
	tmpSl := make([]int, len(*l)+len(*r))
	for i:=0; len(*l)>0 || len(*r)>0; i++ {
		*counter++
		if len(*l) >0 && len(*r)>0 {
			if (*l)[0] < (*r)[0] {
				tmpSl[i] = (*l)[0]
				ll := (*l)[1:]
				l = &ll
			} else {
				tmpSl[i] = (*r)[0]
				rr:=(*r)[1:]
				r=&rr
			}
		} else if len(*l)>0 {
			tmpSl[i] = (*l)[0]
			ll:= (*l)[1:]
			l=&ll
		} else if len(*r)>0 {
			tmpSl[i] = (*r)[0]
			rr:=(*r)[1:]
			r=&rr
		}
	}
	return &tmpSl
}
