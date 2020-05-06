package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type info struct {
		name, addr string
	}
	object := new(info)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter your name")
	scanner.Scan()
	object.name = scanner.Text()
	fmt.Println("enter your address")
	scanner.Scan()
	object.addr = scanner.Text()
	mapD := map[string]string{object.name: object.addr}
	data, _ := json.Marshal(mapD)
	fmt.Println(string(data))
}
