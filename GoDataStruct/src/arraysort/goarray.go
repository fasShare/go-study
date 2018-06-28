package arraysort

import (
	"fmt"
)

// 数组的开闭原则:前闭后开

// 查找start到end内最小的元素下标
func FindMinElemIndex(arr []int, start, end int) int {
	mini := start
	for index := start + 1; index < end; index++ {
		if arr[mini] > arr[index] {
			mini = index
		}
	}
	return mini
}

// 交换数组总right和left下标所指向的元素
func Swap(arr []int, left, right int) {
	tmp := arr[left]
	arr[left] = arr[right]
	arr[right] = tmp
}

// 查找elem在有序数组中的位置，二分查找找到的未必是第一次出现的位置
func BinaryFind(arr []int, start, end int, elem int) (int, bool) {
	for start < end {
		mid := (start + end) / 2
		if arr[mid] == elem {
			return mid, true
		} else if arr[mid] > elem {
			end = mid
		} else {
			start = mid
		}
	}
	return -1, false
}

// 查找第一个大于等于elem的元素在arr中所在的位置
func FindFirstMaxThanElem(arr []int, start, end int, elem int) int {
	for index := start; index < end; index++ {
		if arr[index] >= elem {
			return index
		}
	}
	return end
}

func IntArrayMove(arr []int, from, to, len int) {
	// 数组内的拷贝，按有无重叠去分大体上能够分两种情况：
	// (1)源slice和目标slice有重叠，这种情况又能分为三种小情况
	//  a. to < from   元素从前往后一个一个的拷贝
	//  b. to > from   元素从后向前一个一个的拷贝
	//  c. to == from  直接返回
	// (2)源slice和目标slice无重叠，可以使用有重的叠任何一种拷贝方法

	// 按照from，to的大小关系也能分为三种情况，并且能够兼容元素重叠
	// a. to < from   元素从前往后一个一个的拷贝
	// b. to > from   元素从后向前一个一个的拷贝
	// c. to == from  直接返回

	if to < from {
		for pos := 0; pos < len; pos++ {
			arr[to+pos] = arr[from+pos]
		}
	} else if to > from {
		for pos := len - 1; pos >= 0; pos-- {
			arr[to+pos] = arr[from+pos]
		}
	}
}

func BuildMaxHeap(arr []int) {
	for index := len(arr)/2 - 1; index >= 0; index-- {
		// 判断树index处的节点是否有两个子类
		maxChild := index
		if (2*index + 2) >= len(arr) {
			// 只有一个子类
			maxChild = index*2 + 1
		} else {
			// 两个子类
			if arr[index*2+1] >= arr[index*2+2] {
				maxChild = index*2 + 1
			} else {
				maxChild = index*2 + 2
			}
		}
		if arr[index] < arr[maxChild] {
			// 与较大的子节点交换一下位置
			Swap(arr, index, maxChild)
		}
	}
}

// 堆排序，非稳定排序
func HeapSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		BuildMaxHeap(arr[:i+1])
		Swap(arr, 0, i)
	}
}

func __mergeSort(arr, tmp []int) {
	// 必须是len(arr) > 1，不然会导致死循环
	if len(arr) > 1 {
		mid := len(arr) / 2
		__mergeSort(arr[0:mid], tmp[0:mid])
		__mergeSort(arr[mid:], tmp[mid:])
		// merge
		left, right := 0, mid
		pos := 0
		for (left < mid) && (right < len(arr)) && (pos < len(arr)) {
			if arr[left] <= arr[right] {
				tmp[pos] = arr[left]
				left++
			} else {
				tmp[pos] = arr[right]
				right++
			}
			pos++
		}
		// 直接copy剩下的元素
		for left < mid {
			tmp[pos] = arr[left]
			pos++
			left++
		}
		for right < len(arr) {
			tmp[pos] = arr[right]
			pos++
			right++
		}
		// 必须备份一下
		for i := range arr {
			arr[i] = tmp[i]
		}
	}
}

// 归并排序，稳定排序算法
func MergeSort(arr []int) {
	iarr := make([]int, len(arr))
	__mergeSort(arr, iarr)
}

// 插入排序，稳定性排序算法
func InsertSort(arr []int, start, end int) {
	for sorted := start + 1; sorted < end; sorted++ {
		index := FindFirstMaxThanElem(arr, 0, sorted, arr[sorted])
		if index == sorted {
			continue
		}
		tmp := arr[sorted]
		IntArrayMove(arr, index, index+1, sorted-index)
		arr[index] = tmp
	}
}

// 交换排序，稳定排序算法
func SwapSort(arr []int, start, end int) {
	for sortindex := start; sortindex < end-1; sortindex++ {
		// 在sortindex - end之间查找最小的元素放在sortindex位置处
		mini := FindMinElemIndex(arr, sortindex, end)
		Swap(arr, mini, sortindex)
	}
}

// 快速排序
func FindPivot(arr []int, start, end int) int {
	pivot := arr[start]
	end--
	for start < end {
		for (start < end) && (arr[end] >= pivot) {
			end--
		}
		// 跳出上面那个循环有两种可能:
		// (1)arr[end] =< pivot
		// (2) start == end
		if start == end {
			break
		}
		arr[start] = arr[end]
		start++

		for (start < end) && (arr[start] <= pivot) {
			start++
		}

		// 跳出上面那个循环有两种可能:
		// (1) arr[start] > pivot
		// (2) start == end
		if start == end {
			break
		}
		arr[end] = arr[start]
		end--
	}
	arr[start] = pivot
	return start
}

// 快速排序，非稳定性排序算法
func QuickSort(arr []int, start, end int) {
	if start < end {
		pivot := FindPivot(arr, start, end)
		QuickSort(arr, start, pivot)
		QuickSort(arr, pivot+1, end)
	}
}

// 冒泡排序，稳定排序算法
func BubbleSort(arr []int, start, end int) {
	for usort := end; usort > start; usort-- {
		// 一趟比较
		for vi := start + 1; vi < usort; vi++ {
			if arr[vi] < arr[vi-1] {
				Swap(arr, vi, vi-1)
			}
		}
	}
}

func PrintArrayInlineWithPrefix(prefix string, arr []int, start, end int) {
	fmt.Printf("%s", prefix)
	PrintArrayInline(arr, start, end)
}

func PrintArrayInline(arr []int, start, end int) {
	for i := start; i < end; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("")
}
