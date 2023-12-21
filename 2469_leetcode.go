package main

import "fmt"

func main() {
	var celsius float64 = 36.50
	var res []float64 = []float64 {
		(celsius + 273.15),
		(celsius * 1.80 + 32.0),
	}
	fmt.Println(res)
}