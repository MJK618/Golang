package main

import (
	"fmt"
	s "strings"
)

func main() {
	var x string
	fmt.Print("Enter a string: ")
	_, _ = fmt.Scan(&x)
	x = s.ToLower(x)
	fmt.Println(x)
	if s.HasPrefix(x, "i") &&
		x[len(x)-1:] == "n" &&
		s.Contains(x, "a") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}