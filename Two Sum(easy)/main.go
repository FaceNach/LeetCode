package main

import "fmt"

func twoSum(nums []int, target int) []int {

	out := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				out = append(out, i)
				out = append(out, j)
			}
		}
	}

	return out
}

func main() {

	nums := []int{2, 7, 11, 15}

	result := twoSum(nums, 13)

	for _, index := range result {
		fmt.Printf("The index is: %v\n", index)
	}

}
