// A program for calculate value for painting
package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, heigth float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("A width of %0.2f if invalid", width)
	}
	if heigth < 0 {
		return 0, fmt.Errorf("A heigth of %0.2f if invalid", heigth)
	}
	area := width * heigth
	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
}
