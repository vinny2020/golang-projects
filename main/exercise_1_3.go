package main

import (
	"fmt"
	"os"
	"strconv"
)

//	Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
//
func main() {

	for index, arg := range os.Args[1:] {

		fmt.Println(strconv.Itoa(index) + ": " + arg + "\n")

	}
}
