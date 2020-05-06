package main

import (
	"fmt"
)

func swap(sli []int, pos int) {
	temp := sli[pos+1]
	sli[pos+1] = sli[pos]
	sli[pos] = temp
}

func main() {
	//variable declaration
	arr := make([]int, 0, 4)
	var length int
	var inputValue int
	fmt.Println("Enter the number of values in the list")
	fmt.Scan(&length)
	for i := 0; i < length; i++ {
		fmt.Printf("Enter the value: ")
		fmt.Scan(&inputValue)
		arr = append(arr, inputValue)
	}
	fmt.Println("before sorting is done:", arr)
	bubbleSort(arr)
}
func bubbleSort(arr []int) {
	for passes := 0; passes < len(arr); passes++ {
		for pos := 0; pos < len(arr)-1; pos++ {
			if arr[pos] > arr[pos+1] {
				swap(arr, pos)
			} //end of if for the swapping
		} //end of for for the sorting the adjacent values
	} //end of for the bubble sorting algorithm
	fmt.Println(arr)
}
