package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")
	name, _ = reader.ReadString('\n')
	fmt.Print("Enter Address: ")
	address, _ = reader.ReadString('\n')

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	mapB, _ := json.Marshal(m)
	fmt.Println(string(mapB))
}