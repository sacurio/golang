package main

import "fmt"

func main(){
	
	Menu();
	intArray := []int64{1, 2, 3, 4, 5}
	PrintArray(intArray)	
}

func Menu(){

	var option int	
	PrintOptions()

	for ok := true; ok; ok = (option!=7) {
		fmt.Printf("Select the option");		
		fmt.Scanf("%d\n", &option);
		switch option {
        case 1:
        	fmt.Println("1")
			//push();
        case 2:
        	fmt.Println("2")
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
            fmt.Println("\n<EXIT>\n\n")
        default:
            fmt.Println("\n\t<Invalid option selected>")
		}
	}

}

func PrintOptions(){
	fmt.Println("\n=======================")
	fmt.Println(" Arrays examples in Go ")
	fmt.Println("=======================")
	fmt.Println("\n 1.PUSH\n 2.POP\n 3.DISPLAY\n 4.FIND BY INDEX\n 5.FIND BY VALUE\n 6.SORT\n 7.EXIT \n")
}

func PrintArray(array []int64) {
	fmt.Println(array)
}
