package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year := now.Second()
	fmt.Println(year)
}
