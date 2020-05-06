package main
import "fmt"

func main(){
	var num float32
	var y int
	fmt.Println("enter a floating point number")
	fmt.Scan(&num)
	y = int(num)
	fmt.Printf("%d", y)
}