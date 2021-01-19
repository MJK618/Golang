package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Enter a floating point number: ")
	_, _ = fmt.Scan(&a)
	fmt.Print(int(a))
}