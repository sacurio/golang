package main

import (
	"fmt"
	"strconv"

	U "./lib"
)

func main() {
	renderMenu()
}

//Function for handle the main menu
func renderMenu() {

	var option int
	var optionStr string
	var ca *U.CustomArray

	for ok := true; ok; ok = (option != 11) {
		PrintOptions(ca)
		fmt.Printf(":: Select an option >> ")
		fmt.Scanf("%s\n", &optionStr)
		if U.IsNumber(optionStr) {

			option, _ := strconv.ParseInt(optionStr, 10, 0)
			switch option {
			case 1:
				ca = U.CreateArray()
				U.PrintArray(ca)
			case 2:
				if ControlMenu(ca) {
					U.AppendToArray(ca)
					U.PrintArray(ca)
				}
			case 3:
				if ControlMenu(ca) {
					U.PopArray(ca)
					U.PrintArray(ca)
				}
			case 4:
				if ControlMenu(ca) {
					U.FindByIndexArray(ca)
				}
			case 5:
				if ControlMenu(ca) {
					U.FindByValueArray(ca)
				}
			case 6:
				if ControlMenu(ca) {
					U.InsertAtIndex(ca)
					U.PrintArray(ca)
				}
			case 7:
				if ControlMenu(ca) {
					U.DeleteAtIndex(ca)
					U.PrintArray(ca)
				}
				//sort();
			case 8:
				if ControlMenu(ca) {
					U.SortIndex(ca)
					U.PrintArray(ca)
				}
			case 9:
				if ControlMenu(ca) {
					U.PrintArray(ca)
				}
			case 10:
				fmt.Println("\n<EXIT>\n\n")
				return
			default:
				fmt.Println("\n\t<Invalid option selected. Please, try again>")
				PrintOptions(ca)
			}
		} else {
			fmt.Println("\n\t<Invalid value. Please, try again>")
			PrintOptions(ca)
		}
	}

}

//Function that prints the menu options.
func PrintOptions(a *U.CustomArray) {
	fmt.Println("\n=======================")
	fmt.Println(" Arrays examples in Go ")
	fmt.Println("=======================")
	if a == nil {
		fmt.Println("\n  1. INITIALIZE\n 10. EXIT\n")
	} else {
		fmt.Println("  2. PUSH\n  3. POP\n  4. FIND BY INDEX\n  5. FIND BY VALUE\n  6. INSERTAT\n  7. DELETEAT\n  8. SORT\n  9. PRINT\n 10. EXIT \n")
	}
}

//Function for controlling if a CustomArray struct exists in order to select the correct menu option
func ControlMenu(ca *U.CustomArray) bool {
	if ca == nil {
		fmt.Println("<Invalid menu option>")
		return false
	}
	return true
}
