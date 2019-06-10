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

	for ok := true; ok; ok = (option != 8) {
		PrintOptions()
		fmt.Printf("Select an option: ")
		fmt.Scanf("%s\n", &optionStr)
		if U.IsNumber(optionStr) {

			option, _ := strconv.ParseInt(optionStr, 10, 0)
			switch option {
			case 1:
				a = U.CreateArray()
				fmt.Println(a)
				U.PrintArray(a)
				//push();
			case 2:
				fmt.Println("2")
				fmt.Println("*********************\n")
				U.PrintArray(a)
				//pop();
			case 3:
				fmt.Println("3")
				//print();
			case 4:
				fmt.Println("4")
				//find_by_index();
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
	fmt.Println("\n 1.INITIALIZE\n 2.PUSH\n 3.POP\n 4.DISPLAY\n 5.FIND BY INDEX\n 6.FIND BY VALUE\n 7.SORT\n 8.EXIT \n")
}

func PrintArray(array []int64) {
	fmt.Println(array)
}
