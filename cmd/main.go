package main

import (
	"flag"
	"fmt"
	"nombre"
	"os"
	"strconv"
)

func main() {
	shortenOne := flag.Bool("un", false, `accorcia l'uno (esempio: "ventunmila")`)
	spaces := flag.Bool("spazi", false, `spazi tra gli ordini di grandezza (esempio: "trentamila duecento quarantre")`)
	flag.Parse()

	options := nombre.Options{
		SpaceAfterMinus:              *spaces,
		SpaceBetweenOrderOfMagnitude: *spaces,
		ShortenOne:                   *shortenOne,
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "i need at least one argument")
		os.Exit(1)
	}

	for _, arg := range flag.Args() {
		number, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\"%s\" is not a valid number\n", arg)
			os.Exit(1)
		}

		formatted := nombre.ConvertWithOptions(number, options)
		fmt.Println(formatted)
	}
}
