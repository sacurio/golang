package lib

import (
	"fmt"
)

func CreateArray() []int {

	var l int
	fmt.Println("Input the length of the array: ")
	fmt.Scanf("%d", &l)

	return ResizeArray([]int{}, l)

}

func ResizeArray(a []int, l int) []int {

	return make([]int, len(a)+l)

}

func PrintArray(a []int) {
	fmt.Println(printArray(a))
}

func printArray(a []int) string {
	//var b bytes.Buffer
	for i, element := range a {
		fmt.Println(i, ":", element)
	}
	return "b.String()"
}
