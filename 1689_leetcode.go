package main

import (
	"fmt"
)

func main() {
	var n string = "42332"
    maxDigit := 0
    for i := 0; i < len(n); i++ {
        digit := int(n[i] - '0')
        if digit > maxDigit {
            maxDigit = digit
        }
    }
    fmt.Println(maxDigit)
}