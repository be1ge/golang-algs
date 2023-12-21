package main

import (
	"fmt"
)

func main() {
	var nums [4]int = [4]int {2,7,11,15};
	var target int = 9;
	fmt.Println(twoSum(nums[:], target));
}

func twoSum(nums []int, target int) []int {
	var result [2]int; 
    for i := 0; i < len(nums); i++ {
		for g := 0; g < len(nums); g++ {
			if g == i {
				break;
			} else if nums[g] + nums[i] == target {
				result[0] = g;
				result[1] = i;
			}
		}
	}
	return result[:];
}