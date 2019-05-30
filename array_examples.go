package main

import "fmt"

func main(){
	fmt.Println("Some array examples")

	intArray := []int64{1, 2, 3, 4, 5}

	PrintArray(intArray)
}

func PrintArray(array []int64) {
	fmt.Println(array)
}
