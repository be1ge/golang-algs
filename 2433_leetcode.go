package main

import "fmt"

func main() {
	var pref []int = []int{5, 2, 0, 3, 1}
    res := make([]int, len(pref))
    res[0] = pref[0]
    for i := 1; i < len(pref); i++ {
        res[i] = pref[i] ^ pref[i-1]
    }
	fmt.Println(res)
}