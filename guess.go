// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	//seconds := time.Now().Unix() - актуально в старых версиях go
	//rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	success := false
	reader := bufio.NewReader(os.Stdin)
	for x := 10; x > 0; x-- {
		fmt.Println("You have ", x, "attempts.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("NICE!!!!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was: ", target)
	}
}
