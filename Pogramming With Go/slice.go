package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var x string
	xSlice := make([]int, 3, 9999)
	length := len(xSlice)
	m := 0
	for true {
		_, _ = fmt.Scan(&x)
		number, err := strconv.Atoi(x)
		if x == "X" {
			break
		}
		if err == nil {
			if m < length {
				xSlice[m] = number
			} else {
				xSlice = append(xSlice, number)
			}
			m++
			sort.Ints(xSlice)
			fmt.Print(xSlice)
		}
	}
}