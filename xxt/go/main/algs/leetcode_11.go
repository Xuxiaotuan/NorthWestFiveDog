package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	minLen := 0
	area := 0
	for l < r {
		if height[l] > height[r] {
			minLen = height[r]
		} else {
			minLen = height[l]
		}
		temp := minLen * (r - l)
		if temp > area {
			area = temp
		}
		if height[l] <= height[r] { // 移动较小的那一端
			l++
		} else {
			r--
		}
	}
	return area
}

func main() {
	// 第一个测试用例
	test1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("maxArea(%v) = %d\n", test1, maxArea(test1))

	// 第二个测试用例
	test2 := []int{1, 1}
	fmt.Printf("maxArea(%v) = %d\n", test2, maxArea(test2))

	// 第三个测试用例
	test3 := []int{4, 3, 2, 1, 4}
	fmt.Printf("maxArea(%v) = %d\n", test3, maxArea(test3))
}
