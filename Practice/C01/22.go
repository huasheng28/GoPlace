package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
