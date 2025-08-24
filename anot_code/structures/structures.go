package main

import "fmt"

func main() {
	var subscribers struct {
		name   string
		rate   float64
		active bool
	}
	subscribers.name = "Aman Singh"
	subscribers.rate = 4.99
	subscribers.active = true
	fmt.Println("Name:", subscribers.name)
	fmt.Println("Monthly rate:", subscribers.rate)
	fmt.Println("Active?", subscribers.active, "\n")

	var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)
}
