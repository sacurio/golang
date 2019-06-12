package lib

import (
	"fmt"
	"sort"
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

	fmt.Println("  :: Input the element to be append >> ")
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

	fmt.Println("  :: Input the insert to be search >> ")
	fmt.Scanf("%s", &i)
	if IsNumber(i) {
		j, _ := strconv.Atoi(i)
		findByIndexArray(j, a)
	} else {
		fmt.Println("<The value entered is not correct. Please, try again>")
	}
}

func findByIndexArray(j int, a *A) {
	if j >= 0 && j < a.L {
		for i, element := range a.Arr {
			if i == j {
				fmt.Printf("  :: Element found arr[%d] = %d :: \n", i, element)
				return
			}
		}
		fmt.Println("<The index was not found>")
	} else {
		fmt.Println("<The index out of bounds>")
	}
}

//========FIND BY INDEX ========

//========FIND BY VALUE ========

func FindByValueArray(a *A) {
	var i string

	fmt.Println("  :: Input the value to be search in the Array >> ")
	fmt.Scanf("%s", &i)
	if IsNumber(i) {
		j, _ := strconv.Atoi(i)
		findByValueArray(j, a)
	} else {
		fmt.Println("<The value entered is not correct. Please, try again>")
		AppendToArray(a)
	}
}

func findByValueArray(j int, a *A) {
	var b *A
	b = new(A)
	resizeArray(0, b)
	for _, element := range a.Arr {
		if element == j {
			appendToArray(element, b)
		}
	}
	if len(b.Arr) > 0 {
		fmt.Printf("\n:: %d elements founded ::\n\n", len(b.Arr))
		printArray(b.Arr)
	} else {
		fmt.Printf("<No values founded for value: %d\n>", j)
	}

}

//========FIND BY VALUE ========

//========INSERT AT========

func InsertAtIndex(a *A) {

	var i, v string

	fmt.Println("  :: Input the index and value to insert >> ")
	fmt.Scanf("%s", &i)
	fmt.Println("  :: Input the value and value to insert >> ")
	fmt.Scanf("%s", &v)
	if IsNumber(i) && IsNumber(v) {
		i, _ := strconv.Atoi(i)
		v, _ := strconv.Atoi(v)
		insertAtIndex(i, v, a)
	} else {
		fmt.Println("<Exists an error with the type of the values entereds>")
	}

}

func insertAtIndex(i int, v int, a *A) {
	if i >= 0 && i < len(a.Arr) {
		l := make([]int, i)
		copy(l, a.Arr[0:i])
		l = append(l, v)
		l = append(l, a.Arr[i:]...)
		a.Arr = l
		a.P++
	} else {
		fmt.Println("<The index out of bounds>")
	}
}

//========INSERT AT========

//========DELETE AT========

func DeleteAtIndex(a *A) {

	var i, v string

	fmt.Println("  :: Input the index to be delete >> ")
	fmt.Scanf("%s", &i)
	if IsNumber(i) && IsNumber(v) {
		i, _ := strconv.Atoi(i)
		deleteAtIndex(i, a)
	} else {
		fmt.Println("<Exists an error with the type of the values entereds>")
	}

}

func deleteAtIndex(i int, a *A) {
	if i >= 0 && i < len(a.Arr) {
		l := make([]int, len(a.Arr)-1)
		l = append(a.Arr[:i], a.Arr[i+1:]...)
		a.Arr = l
		a.P--
	} else {
		fmt.Println("<The index out of bounds>")
	}
}

//========DELETE AT========

//========SORT=======

func SortIndex(a *A) {

	sortIndex(a)

}

func sortIndex(a *A) {
	s := make([]int, len(a.Arr))
	copy(s, a.Arr)
	sort.Ints(s)
	a.Arr = s
}

//========SORT=======

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
