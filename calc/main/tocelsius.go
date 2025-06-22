// преобразует температуру в градусах по Фаренгейту в градусы по цельсию
package main

import (
	"fmt"
	"log"

	"example.com/calc/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celcius)
}
