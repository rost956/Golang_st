// average предназначен для подсчета среднего из введенных чисел при запуске программы
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sum := 0.0
	for _, str := range os.Args[1:] {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	fmt.Printf("Average: %0.2f\n", sum/float64(len(os.Args[1:])))
}
