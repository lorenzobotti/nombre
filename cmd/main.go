package main

import (
	"fmt"
	"nombre"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "i need at least one argument")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		number, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\"%s\" is not a valid number\n", arg)
			os.Exit(1)
		}

		formatted := nombre.Convert(number)
		fmt.Println(formatted)
	}
}
