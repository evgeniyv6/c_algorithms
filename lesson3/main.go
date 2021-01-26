package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson3/searching"
	"github.com/evgeniyv6/c_algorithms/lesson3/slices"
	"github.com/evgeniyv6/c_algorithms/lesson3/sorting"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	arr:= slices.GenerSlice(100)
	fmt.Printf("Исходный слайс - %v\n", arr)
	// fmt.Printf("Шейкерная сортировка. Количество операций: %d\n",sorting.ShakerSorting(&arr))
	// fmt.Printf("Шейкерная сортировка с флагом. Количество операций: %d\n",sorting.ShakerSortingFlag(&arr))
	fmt.Printf("Шейкерная сортировка с флагом и смещением на позицию перестановки. Количество операций: %d\n",sorting.ShakerSortingMemPos(&arr))
	fmt.Printf("Отсортированный слайс: %v\n",arr)

	fmt.Println()
	// binary search
	// Подаем на вход отсортированный выше массив arr
	var num = 9
	fmt.Println("Бинарный поиск")
	if ok, idx := searching.BinSearch(&arr, num); ok {
		fmt.Printf("Число %d есть в массиве. Индекс: %d", num, idx)
	} else {
		fmt.Printf("Числа %d нет в массиве. Результат: %d", num, idx)
	}
}
