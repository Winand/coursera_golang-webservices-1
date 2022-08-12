package main

import "fmt"

func main() {
	option1()
}

func option1() {
	nums := []int{16, 8, 42, 4, 23, 15}
	mx := 0
	for i, v := range nums {
		if v > nums[mx] {
			mx = i
		}
	}
	fmt.Println(nums[mx])
}

// lecturer's solution
func option2() {
	nums := []int{16, 8, 42, 4, 23, 15}
	mx := nums[0]
	for _, v := range nums[1:] {
		if v > mx {
			mx = v
		}
	}
	fmt.Println(mx)
}
