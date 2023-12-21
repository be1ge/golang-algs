package main

import "fmt"

func main() {
	var operations []string = []string{"--X","X++","X++"}
	var x int = 0
	for i := 0; i < len(operations); i++ {
		if operations[i] == "--X" || operations[i] == "X--" {
			x--
		} else if operations[i] == "X++" || operations[i] == "++X" {
			x++
		}
	}
	fmt.Println(x)
}