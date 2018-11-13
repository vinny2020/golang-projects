package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//print only the first argument (name of the program)
	fmt.Println(strings.Join(os.Args[0:1], " "))
}
