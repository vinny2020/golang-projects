package main

import (
	"fmt"
	"os"
	"strconv"
)

//print index and arg with a line break
func main() {

	for index, arg := range os.Args[1:] {

		fmt.Println(strconv.Itoa(index) + ": " + arg + "\n")

	}
}
