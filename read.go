package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type names struct {
	fname, lname string
} //Name contains fields for first and last name

func main() {
	var fileName string
	fmt.Println("Enter the file name")
	fmt.Scanln(&fileName)
	data, _ := os.Open(fileName)
	log := []names{}
	lines := bufio.NewScanner(data)
	for i := 0; lines.Scan(); i++ {

		fullname := lines.Text()
		sepName := strings.Split(fullname, " ")
		temp := names{}
		temp.fname = sepName[0]
		temp.lname = sepName[1]
		log = append(log, temp)

	}
	for idx := range log {
		fmt.Printf("%v\n", log[idx])
	}
}
