package main

import (
	"fmt"
	"strconv"

	U "./lib"
)

func main() {
	Menu()
}

//Function for handle the main menu
func Menu() {

	var option int
	var optionStr string
	var a *U.A

	for ok := true; ok; ok = (option != 11) {
		PrintOptions(a)
		fmt.Printf(":: Select an option >> ")
		fmt.Scanf("%s\n", &optionStr)
		if U.IsNumber(optionStr) {

			option, _ := strconv.ParseInt(optionStr, 10, 0)
			switch option {
			case 1:
				a = U.CreateArray()
				U.PrintArray(a)
			case 2:
				if ControlMenu(a) {
					U.AppendToArray(a)
					U.PrintArray(a)
				}
			case 3:
				if ControlMenu(a) {
					U.PopArray(a)
					U.PrintArray(a)
				}
			case 4:
				if ControlMenu(a) {
					U.FindByIndexArray(a)
				}
			case 5:
				if ControlMenu(a) {
					U.FindByValueArray(a)
				}
			case 6:
				if ControlMenu(a) {
					U.InsertAtIndex(a)
					U.PrintArray(a)
				}
			case 7:
				if ControlMenu(a) {
					U.DeleteAtIndex(a)
					U.PrintArray(a)
				}
				//sort();
			case 8:
				if ControlMenu(a) {
					U.SortIndex(a)
					U.PrintArray(a)
				}
			case 9:
				if ControlMenu(a) {
					U.PrintArray(a)
				}
			case 10:
				fmt.Println("\n<EXIT>\n\n")
				return
			default:
				fmt.Println("\n\t<Invalid option selected. Please, try again>")
				PrintOptions(a)
			}
		} else {
			fmt.Println("\n\t<Invalid value. Please, try again>")
			PrintOptions(a)
		}
	}

}

func PrintOptions(a *U.A) {
	fmt.Println("\n=======================")
	fmt.Println(" Arrays examples in Go ")
	fmt.Println("=======================")
	if a == nil {
		fmt.Println("\n  1. INITIALIZE\n 10. EXIT\n")
	} else {
		fmt.Println("  2. PUSH\n  3. POP\n  4. FIND BY INDEX\n  5. FIND BY VALUE\n  6. INSERTAT\n  7. DELETEAT\n  8. SORT\n  9. PRINT\n 10. EXIT \n")
	}
}

func ControlMenu(a *U.A) bool {
	if a == nil {
		fmt.Println("<Invalid menu option>")
		return false
	}
	return true
}
