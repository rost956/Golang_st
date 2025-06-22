// pass_fail сообщает, сдал ли пользователь экзамен.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func main() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 { //условие
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status) //результат
}
