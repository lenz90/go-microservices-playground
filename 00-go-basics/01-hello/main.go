package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, Go!")
	if len(os.Args) > 1 {
		fmt.Println("Args:", strings.Join(os.Args[1:], ", "))
	} else {
		fmt.Println("Args: none")
	}
}
