package main

import "fmt"

func main() {

	var input string

	fmt.Println("\nThis utility will convert string from hexadecimal\n\n")

	fmt.Println("Enter Your Input Below and Press Enter: \n")

	_, err := fmt.Scanf("%s", &input)

	if err != nil {
		fmt.Println("\n", err)
		return
	} else {
		fmt.Println("\nHexaDecimal Converted Results of Your Input : \n")
		fmt.Printf("%x\n\n", input)
	}

}
