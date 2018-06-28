package main

import (
	arraysort "arraysort"
	"fmt"
)

func arraysortTest() {
	arr := []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.PrintArrayInline(arr, 0, len(arr))
	arr = []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.QuickSort(arr, 0, len(arr))
	arraysort.PrintArrayInlineWithPrefix("QuickSort:", arr, 0, len(arr))
	arr = []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.BubbleSort(arr, 0, len(arr))
	arraysort.PrintArrayInlineWithPrefix("BubbleSort:", arr, 0, len(arr))
	arr = []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.SwapSort(arr, 0, len(arr))
	arraysort.PrintArrayInlineWithPrefix("SwapSort:", arr, 0, len(arr))
	arr = []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.InsertSort(arr, 0, len(arr))
	arraysort.PrintArrayInlineWithPrefix("InsertSort:", arr, 0, len(arr))
	arr = []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.MergeSort(arr)
	arraysort.PrintArrayInlineWithPrefix("MergeSort:", arr, 0, len(arr))
	arr = []int{2, 5, 8, 1, 2, 1, 1, 0, 3, 5, 6, 7, 8, 9, 1, -1}
	arraysort.HeapSort(arr)
	arraysort.PrintArrayInlineWithPrefix("HeapSort:", arr, 0, len(arr))
}

func InArrayMoveTest() {
	arr := []int{0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 0, 0, 0, 0, 0}
	arraysort.IntArrayMove(arr, 4, 5, 7)
	arraysort.PrintArrayInlineWithPrefix("InArrayMoveToEnd:", arr, 0, len(arr))

	arr = []int{0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 0, 0, 0, 0, 0}
	arraysort.IntArrayMove(arr, 4, 3, 7)
	arraysort.PrintArrayInlineWithPrefix("InArrayMoveToBegin:", arr, 0, len(arr))
}

func main() {
	arraysortTest()
	//InArrayMoveTest()
	fmt.Println("End of Main!")
}
