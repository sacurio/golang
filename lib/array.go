package lib

import (
	"fmt"
	"sort"
	"strconv"
)

//========STRUCTS========

type CustomArray struct {
	Arr []int
	Lng int //Array length
	Ptr int //Array pointer
}

//========STRUCTS========

//========INITIALIZE ARRAY========

//Exportable fuction that initalize an array
func CreateArray() *CustomArray {
	var lng int
	var ca *CustomArray
	ca = new(CustomArray)

	fmt.Println("\nInitializing...")
	ca.Lng = lng
	ca.Ptr = 0
	resizeArray(0, ca)
	return ca
}

//========INITIALIZE ARRAY========

//========RESIZE========

//Exportable function that resize an existing array
func ResizeArray(ca *CustomArray) {
	resizeArray(ca.Lng+1, ca)
}

//Function that resize an existing array
func resizeArray(lng int, ca *CustomArray) {
	newarr := make([]int, lng)
	copy(newarr, ca.Arr)
	ca.Lng = lng
	ca.Arr = newarr
}

//========RESIZE========

//========APPEND========

//Exportable function that append an element entered by user to an existing array.
func AppendToArray(ca *CustomArray) {
	var entry string
	var valid bool

	for ok := true; ok; ok = (valid == false) {
		fmt.Println("  :: Input the element to be append >> ")
		fmt.Scanf("%s", &entry)
		valid = IsNumber(entry)
		if valid {
			i, _ := strconv.Atoi(entry)
			appendToArray(i, ca)
		} else {
			fmt.Println("<The value entered is not correct. Please, try again>")
		}
	}
}

//Function that append an element entered to an existing array.
func appendToArray(element int, ca *CustomArray) {

	var slide = [1]int{element}
	var capicity = len(ca.Arr) + 1
	if ca.Ptr <= capicity {
		resizeArray(len(ca.Arr)+1, ca)
	}
	copy(ca.Arr[ca.Ptr:len(ca.Arr)], slide[0:1])
	ca.Ptr++
}

//========APPEND========

//========POP========

//Exportable function that remove the last element in an array.
func PopArray(ca *CustomArray) {
	popArray(ca)
}

//Function that remove the last element in an array.
func popArray(ca *CustomArray) {
	var lng = len(ca.Arr) - 1
	if lng >= 0 {
		resizeArray(lng, ca)
		ca.Ptr--
	} else {
		fmt.Println("<Index out of bounds>")
	}
}

//========POP========

//========FIND BY INDEX ========

//Exportable function that finds an element into the array by the index value entered by the user and print in screen the result.
func FindByIndexArray(ca *CustomArray) {
	var index string

	fmt.Println("  :: Input the insert to be search >> ")
	fmt.Scanf("%s", &index)
	if IsNumber(index) {
		indexInt, _ := strconv.Atoi(index)
		findByIndexArray(indexInt, ca)
	} else {
		fmt.Println("<The value entered is not correct. Please, try again>")
	}
}

//Function that finds an element into the array by the index and print in screen the result.
func findByIndexArray(index int, ca *CustomArray) {
	if index >= 0 && index < ca.Lng {
		for i, element := range ca.Arr {
			if i == index {
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

//Exportable function that finds an element into the array by the value value entered by the user and print in screen the result.
func FindByValueArray(ca *CustomArray) {
	var value string

	fmt.Println("  :: Input the value to be search in the Array >> ")
	fmt.Scanf("%s", &value)
	if IsNumber(value) {
		valueInt, _ := strconv.Atoi(value)
		findByValueArray(valueInt, ca)
	} else {
		fmt.Println("<The value entered is not correct. Please, try again>")
		AppendToArray(ca)
	}
}

//Function that finds an element into the array by the index value entered by the user and print in screen the result(s).
func findByValueArray(value int, ca *CustomArray) {
	var b *CustomArray
	b = new(CustomArray)
	resizeArray(0, b)
	for _, element := range ca.Arr {
		if element == value {
			appendToArray(element, b)
		}
	}
	if len(b.Arr) > 0 {
		fmt.Printf("\n:: %d elements founded ::\n\n", len(b.Arr))
		printArray(b.Arr)
	} else {
		fmt.Printf("<No values founded for value: %d\n>", value)
	}

}

//========FIND BY VALUE ========

//========INSERT AT========

//Exportable function that inserts a value in an existing array in the index position entered by the user.
func InsertAtIndex(ca *CustomArray) {

	var index, value string

	fmt.Println("  :: Input the index and value to insert >> ")
	fmt.Scanf("%s", &index)
	fmt.Println("  :: Input the value and value to insert >> ")
	fmt.Scanf("%s", &value)
	if IsNumber(index) && IsNumber(value) {
		indexInt, _ := strconv.Atoi(index)
		valueInt, _ := strconv.Atoi(value)
		insertAtIndex(indexInt, valueInt, ca)
	} else {
		fmt.Println("<Exists an error with the type of the values entereds>")
	}

}

//Function that inserts a value in an array in the index position.
func insertAtIndex(index int, value int, ca *CustomArray) {
	if index >= 0 && index < len(ca.Arr) {
		slide := make([]int, index)
		copy(slide, ca.Arr[0:index])
		slide = append(slide, value)
		slide = append(slide, ca.Arr[index:]...)
		ca.Arr = slide
		ca.Ptr++
	} else {
		fmt.Println("<The index out of bounds>")
	}
}

//========INSERT AT========

//========DELETE AT========

//Exportable function that remove an element in an existing array by the index entered by the user.
func DeleteAtIndex(ca *CustomArray) {

	var index, value string

	fmt.Println("  :: Input the index to be delete >> ")
	fmt.Scanf("%s", &index)
	if IsNumber(index) && IsNumber(value) {
		indexInt, _ := strconv.Atoi(index)
		deleteAtIndex(indexInt, ca)
	} else {
		fmt.Println("<Exists an error with the type of the values entereds>")
	}

}

//Function that inserts a value in an existing array in the index position.
func deleteAtIndex(index int, ca *CustomArray) {
	if index >= 0 && index < len(ca.Arr) {
		slide := make([]int, len(ca.Arr)-1)
		slide = append(ca.Arr[:index], ca.Arr[index+1:]...)
		ca.Arr = slide
		ca.Ptr--
	} else {
		fmt.Println("<The index out of bounds>")
	}
}

//========DELETE AT========

//========SORT=======

//Exportable function that sorts the elements of the array using the bubble sorting algorithm.
func SortIndex(ca *CustomArray) {

	sortIndex(ca)

}

//Function that sorts the elements of the array using the bubble sorting algorithm.
func sortIndex(ca *CustomArray) {
	slide := make([]int, len(ca.Arr))
	copy(slide, ca.Arr)
	sort.Ints(slide)
	ca.Arr = slide
}

//========SORT=======

//========PRINT========

//Exportable function that prints in screen the elements of the current array.
func PrintArray(ca *CustomArray) {
	printArray(ca.Arr)
}

//Function that prints in screen the elements of the current array.
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
