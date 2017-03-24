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
	for i := 0; i < len(num); i++ {
		sum = append(sum, num[i])
	}
	return sum
}
