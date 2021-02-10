package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson3/slices"
	"github.com/evgeniyv6/c_algorithms/lesson3/sorting"
	"github.com/evgeniyv6/c_algorithms/lesson8/complexalgos"
	"strconv"
	"time"
)


// Counter descr
var Counter int
var resultsTable = [][]string{
{"Название функции", "Описание","Время 100 эл. (сек.)", "Количество сравнений 100 эл.",
"Время 10000 эл. (сек.)", "Количество сравнений 10000 эл.",
"Время 1000000 эл. (сек.)", "Количество сравнений 1000000 эл."},
}

var csvDataMapa = map[string][]string{"QuickSort": {"QuickSort","Сортировка Хоара"},
"MergeSort": {"MergeSort", "Сортировка слиянием"},
"ShakerSorting": {"ShakerSorting", "Шейкерная сортировка"},
"ShakerSortingFlag":{"ShakerSortingFlag", "Шейкерная сортировка (если при проходе - нет перестановок, то стоп)"},
"ShakerSortingMemPos": {"ShakerSortingMemPos", "Шейкерная сортировка (с запоминанием позиции перестановки)"},
"BubleSort": {"BubleSort", "Пузырьковая сортировка"},
"InsertationSort": {"InsertationSort", "Сортировка вставками"}}

func main() {
	go spinner(100*time.Millisecond)
	var attemptsSlice = []int{10, 100, 100000}

	// В функциях сортировки - сортируем слайсы на месте, не копируем слайс в процедуре.
	// Операция копирования вынесена за пределы функций
	for _, attempts := range attemptsSlice {
		Counter = 0
		testSlice := slices.GenerSlice(attempts)
		duplicateTestSlice := append([]int(nil), testSlice...)
		startTime := time.Now()
		complexalgos.QuickSort(&duplicateTestSlice, 0, len(duplicateTestSlice)-1, &Counter)
		elapsedTime := time.Since(startTime)
		sec, _ := time.ParseDuration(elapsedTime.String())
		csvDataMapa["QuickSort"] = append(csvDataMapa["QuickSort"],fmt.Sprintf("%f", sec.Seconds()), strconv.Itoa(Counter))
		duplicateTestSlice = nil
		Counter = 0

		duplicateTestSlice = append([]int(nil), testSlice...)
		startTime = time.Now()
		complexalgos.MergeSort(&duplicateTestSlice, &Counter)
		elapsedTime = time.Since(startTime)
		sec, _ = time.ParseDuration(elapsedTime.String())
		csvDataMapa["MergeSort"] = append(csvDataMapa["MergeSort"], fmt.Sprintf("%f", sec.Seconds()), strconv.Itoa(Counter))
		duplicateTestSlice = nil

		evalFunc(sorting.ShakerSorting, testSlice, "ShakerSorting")
		evalFunc(sorting.ShakerSortingFlag, testSlice, "ShakerSortingFlag")
		evalFunc(sorting.ShakerSortingMemPos, testSlice, "ShakerSortingMemPos")
		evalFunc(sorting.InsertationSort, testSlice, "InsertationSort")
		evalFunc(sorting.BubleSort, testSlice, "BubleSort")
	}

	for _, v := range csvDataMapa {
		resultsTable = append(resultsTable, v)
	}
	if err := CreateCsvFile(resultsTable); err != nil {
		fmt.Println("Error from csv creator: ", err)
	}

}

func evalFunc(f func(*[]int) int, sl []int, name string) {
	duplicateTestSlice := append([]int(nil), sl...)
	startTime := time.Now()
	count :=f(&duplicateTestSlice)
	elapsedTime := time.Since(startTime)
	sec, _ := time.ParseDuration(elapsedTime.String())
	csvDataMapa[name] = append(csvDataMapa[name],fmt.Sprintf("%f", sec.Seconds()), strconv.Itoa(count))
	duplicateTestSlice = nil
}

func spinner(delay time.Duration) {
	for {
		for _, ch :=range `-\|/` {
			fmt.Printf("\r%c", ch)
			time.Sleep(delay)
		}
	}
}

