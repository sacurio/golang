package lib

import (
	"fmt"
)

type A struct {
	Arr []int
	L   int
}

func CreateArray() *A {

	var l int
	var a *A

	a = new(A)

	fmt.Println("Input the length of the array: ")
	fmt.Scanf("%d", &l)

	fmt.Println("\nInitializing...")
	a.L = l
	ResizeArray(a)
	return a
}

func ResizeArray(a *A) {
	r := make([]int, a.L)
	copy(r, a.Arr)
	a.Arr = r
}

func PrintArray(a *A) {
	for i, element := range a.Arr {
		fmt.Printf("arr[%d] = %d\n", i, element)
	}
}
