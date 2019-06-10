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
		PrintOptions()
		fmt.Printf("Select an option: ")
		fmt.Scanf("%s\n", &optionStr)
		if U.IsNumber(optionStr) {

			option, _ := strconv.ParseInt(optionStr, 10, 0)
			switch option {
			case 1:
				a = U.CreateArray()
				U.PrintArray(a)
			case 2:
				U.AppendToArray(a)
				U.PrintArray(a)
			case 3:
				U.PopArray(a)
				U.PrintArray(a)
			case 4:
				U.FindByIndexArray(a)
			case 5:
				fmt.Println("5")
				//find_by_value();
			case 6:
				fmt.Println("6")
				//sort();
			case 7:
				fmt.Println("6")
				//sort();
			case 8:
				U.PrintArray(a)
			case 9:
				//SORT
			case 10:
				fmt.Println("\n<EXIT>\n\n")
				return
			default:
				fmt.Println("\n\t<Invalid option selected. Please, try again>")
				PrintOptions()
			}
		} else {
			fmt.Println("\n\t<Invalid value. Please, try again>")
			PrintOptions()
		}
	}

}

func PrintOptions() {
	fmt.Println("\n=======================")
	fmt.Println(" Arrays examples in Go ")
	fmt.Println("=======================")
	fmt.Println("\n 1.INITIALIZE\n 2.PUSH\n 3.POP\n 4.FIND BY INDEX\n 5.INSERTAT\n 6.DELETEAT\n 7.FIND BY VALUE\n 9.SORT\n 10.PRINT\n 11.EXIT \n")
}

func PrintArray(array []int64) {
	fmt.Println(array)
}
