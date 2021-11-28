package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 2, 9, 1, 5, 3, 2, 4}
	fmt.Println("排序前", arr)
	quickSort(arr)
	fmt.Println("排序后", arr)
}

// 插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			j := i - 1
			tmp := arr[i]
			// 在 i前面找一个合适的位置插入进去
			for ; j >= 0 && arr[j] > tmp; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = tmp
		}
	}
}

// 冒泡排序
func bubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			// 一轮下来将最大的数字放在了后面
			if arr[j+1] < arr[j] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

// 快排
func quickSort(arr []int) {
	fQuickSort(arr, 0, len(arr)-1)
}

func fQuickSort(arr []int, beginIndex, endIndex int) {
	if beginIndex >= endIndex {
		return
	}
	partition := doublePartition(arr, beginIndex, endIndex)
	fQuickSort(arr, beginIndex, partition-1)
	fQuickSort(arr, partition+1, endIndex)
}

func doublePartition(arr []int, low, high int) int {
	fmt.Println("轮次", low, high, arr)
	left := low
	right := high
	pivot := arr[low]

	for left != right {
		for left < right && arr[left] <= pivot {
			left++
			fmt.Println("do left 轮次", low, high, left, right)
		}
		for left < right && arr[right] > pivot {
			right--
			fmt.Println("do right 轮次", low, high, left, right)
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[low], arr[left] = arr[left], arr[low]

	return left
}
