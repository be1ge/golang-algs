package main

import (
	"fmt"
)


func main() {
	var nums []int = []int {1,2,3,1,1,3}
	fmt.Println(numIdenticalPairs(nums[:]))
}
func numIdenticalPairs(nums []int) int {
	var result uint16 = 0;
    for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
			} else if i < j {
				if nums[i] == nums[j]{
					result++;
				}
			}
		}
	}
	return int(result)
}