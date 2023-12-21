package main 

import (
	"fmt"
	"strings"
)

func main() {
    str := "address: \"1.1.1.1\""
    newStr := strings.Split(str, "address: ");
	str = newStr[1];
	newStr = strings.Split(str, "\"")
	str = newStr[1]
	str = strings.Replace(str, ".", "[.]", 4);
	fmt.Println(str)
}