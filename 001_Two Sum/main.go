package main

import "fmt"

func main() {
	nums := []int{2, 4, 5}
	target := 3
	answer := twoSum(nums, target)
	fmt.Println(answer)
}

func twoSum(num []int, target int) []int {
	var sum []int
	for i := 0; i < 2; i++ {
		sum[i] = num[0]
	}
	return sum
}
