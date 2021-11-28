package main

import (
	"fmt"
)

func main() {
	question1()
}

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，
如果目标值存在返回下标，否则返回 -1

示例1：
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例2：
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

*/
func question1() {
	tmpArr := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(find(tmpArr, 9))
}
func find(arr []int, target int) (int, bool) {
	fmt.Println(arr)
	if len(arr) <= 0 {
		return 0, false
	}
	middle := len(arr) / 2
	if arr[middle] == target {
		return arr[middle], true
	}
	find(arr[:middle], target)
	find(arr[middle:], target)
	return 0, false
}
