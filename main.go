package main

import (
	"fmt"
	"strings"
)

func main() {
	var broken string = "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
