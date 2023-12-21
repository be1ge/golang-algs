package main

import (
	"fmt"
)

func main() {
	var nums [3]int = [3]int{1,2,1};
	fmt.Println(getConcatenation(nums[:]));
}

func getConcatenation(nums []int) []int {
	arrLen := len(nums);
    for i := 0; i < arrLen; i++ {
        nums = append(nums, nums[i]);
	}
	return nums[:];
}