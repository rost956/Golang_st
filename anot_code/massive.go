//Изучение массивов

package main

import "fmt"

func ret() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "si"}
	//var words string
	//var index int
	for index, words := range notes {
		fmt.Println(index, words)
	}
}
func average() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum / float64(len(numbers)))
}
func main() {
	ret()
	average()
}
