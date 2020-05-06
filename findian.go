package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter a string")
	var value string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() == nil {
		value = scanner.Text()
	}
	value = strings.ToLower(value)
	a := strings.HasPrefix(value, "i")
	end := strings.HasSuffix(value, "n")
	c := strings.ContainsAny(value, "a")
	//print(a, end, c, value)
	if a && end && c {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
