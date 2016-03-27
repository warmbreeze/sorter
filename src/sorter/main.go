package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "Unsorted data file")
var outfile *string = flag.String("o", "outfile", "Sorted data file")
var algorithm *string = flag.String("a", "quick", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	// read unsort file

}
