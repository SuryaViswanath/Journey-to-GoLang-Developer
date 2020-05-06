package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := make([]int, 3)
	var input string
	cond := true
	for cond {
		fmt.Println("enter the interger; press X after completetion")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			break
		} else {
			temp, _ := strconv.Atoi(input)
			arr = append(arr, temp)
		}

	}
	sort.Sort(sort.IntSlice(arr))
	fmt.Println(arr[3:])
}
