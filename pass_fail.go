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

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)   //создаем для чтения потока ввода
	input, err := reader.ReadString('\n') //читаем до \n
	if err != nil {                       //если есть ошибка
		log.Fatal(err)
	}
	status := ""
	input = strings.TrimSpace(input)            //убираем в конце \n
	grade, err := strconv.ParseFloat(input, 64) //переводим str в float
	if grade >= 60 {                            //условие
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status) //результат
}
