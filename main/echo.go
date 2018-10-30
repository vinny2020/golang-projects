package main

import (
	"fmt"
	"os"
)

func main() {

	var s, t, sep, sep2 string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("print only command line args:\n" + s)

	for j := 0; j < len(os.Args); j++ {
		t += sep2 + os.Args[j]
		sep2 = " "
	}
	fmt.Println("print with program name and args:\n" + t)
}
