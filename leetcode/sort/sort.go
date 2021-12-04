package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 2, 9, 1, 5, 3, 2, 4}
	fmt.Println("排序前", arr)
	insertSort(arr)
	fmt.Println("排序后", arr)
}

// 插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		tmp := arr[i]
		if arr[i] < arr[i-1] {
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
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

// 快排
func quickSort(arr []int) {
	fQuickSort(arr, 0, len(arr)-1)
}
func fQuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	partitionIndex := partition(arr, left, right)
	fQuickSort(arr, left, partitionIndex-1)
	fQuickSort(arr, partitionIndex+1, right)
}
func partition(arr []int, left, right int) int {
	pivotIndex := left
	index := pivotIndex + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivotIndex] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivotIndex], arr[index-1] = arr[index-1], arr[pivotIndex]
	return index - 1
}
