package main

import (
	"fmt"
)


func main() {
	var nums [6]int = [6]int {0,2,1,5,3,4}; 
	fmt.Println(buildArray(nums[:]))
}
func buildArray(nums []int) []int {
	newArr := []int{};
	for i := 0; i < len(nums); i++ {
		newArr = append(newArr, nums[nums[i]]);
	}
	return newArr[:]
}