package main

import (
	"fmt"
)

func main() {
	var str string = "Rostik"
	fmt.Printf("About one-third: %0.2f\n", 40.0/3.0)
	fmt.Printf("Hello, %s\n", str)
	resultStr := fmt.Sprintf("I'm %s\n", str)
	fmt.Print(resultStr)

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("------------------------------")
	fmt.Printf("%12s | %5d\n", "Stamps", 50)
	fmt.Printf("%12s | %5d\n", "Paper clips", 100)
	fmt.Printf("%12s | %5d\n", "Tape", 5)
}
