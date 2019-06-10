package lib

import (
	"fmt"
	"strconv"
)

//========STRUCTS========

type A struct {
	Arr []int
	L   int
	P   int
}

//========STRUCTS========

//========INITIALIZE ARRAY========

func CreateArray() *A {
	var l int
	var a *A
	a = new(A)

	fmt.Println("\nInitializing...")
	a.L = l
	a.P = 0
	resizeArray(0, a)
	return a
}

//========INITIALIZE ARRAY========

//========RESIZE========

func ResizeArray(a *A) {
	resizeArray(a.L+1, a)
}

func resizeArray(l int, a *A) {
	r := make([]int, l)
	copy(r, a.Arr)
	a.L = l
	a.Arr = r
}

//========RESIZE========

//========APPEND========

func AppendToArray(a *A) {
	var e string

	fmt.Println("Input the element to be append: ")
	fmt.Scanf("%s", &e)
	if IsNumber(e) {
		i, _ := strconv.Atoi(e)
		appendToArray(i, a)
	} else {
		fmt.Println("<The value entered is not correct. Please, try again>")
		AppendToArray(a)
	}
}

func appendToArray(element int, a *A) {

	var s = [1]int{element}
	var c = len(a.Arr) + 1
	if a.P <= c {
		resizeArray(len(a.Arr)+1, a)
	}
	copy(a.Arr[a.P:len(a.Arr)], s[0:1])
	a.P++
}

//========APPEND========

//========POP========

func PopArray(a *A) {
	popArray(a)
}

func popArray(a *A) {
	var l = len(a.Arr) - 1
	if l >= 0 {
		resizeArray(l, a)
		a.P--
	} else {
		fmt.Println("<Index out of bounds>")
	}
}

//========POP========

//========FIND BY INDEX ========

func FindByIndexArray(a *A) {
	var i string

	fmt.Println("Input the insert to be search: ")
	fmt.Scanf("%s", &i)
	if IsNumber(i) {
		j, _ := strconv.Atoi(i)
		appendToArray(j, a)
	} else {
		fmt.Println("<The value entered is not correct. Please, try again>")
		AppendToArray(a)
	}
}

func findByIndexArray(i int, a *A) {

}

//========FIND BY INDEX ========

//========INSERT AT========

//========INSERT AT========

//========DELETE AT========

//========DELETE AT========

//========PRINT========

func PrintArray(a *A) {
	printArray(a.Arr)
}

func printArray(ar []int) {
	if len(ar) > 0 {
		for i, element := range ar {
			fmt.Printf("arr[%d] = %d\n", i, element)
		}
	} else {
		fmt.Println("<Array is empty>")
	}
}

//========PRINT========
