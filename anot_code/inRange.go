// inRange выводит сегмент чисел, которые входят в заданный диапазон
package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -12.5, 3.2, 45, 1, 50, 160, -100, 0))
}
