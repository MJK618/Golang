package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Names struct {
	fname, lname string
}

func main() {
	ipReader := bufio.NewReader(os.Stdin)
	names := make([]Names, 0)

	fmt.Print("Enter the file name (hint: names.txt) : ")
	filename, _ := ipReader.ReadString('\n')
	filename = strings.TrimSuffix(filename, "\n")
	filename = strings.TrimSuffix(filename, "\r")

	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n Names
		s := strings.Split(scanner.Text(), " ")
		n.fname, n.lname = s[0], s[1]
		names = append(names, n)
	}
	_ = file.Close()

	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}
}