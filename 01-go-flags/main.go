package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"opcion1": "Quilmes",
	"opcion2": "Salta",
	"opcion3": "Heineken",
}

func main() {
	b := flag.Bool("beers", false, "print beers") //devuelve un puntero
	flag.Parse()
	if *b { //si queremos acceder al valor del puntero hay que poner *
		fmt.Println(beers)
	}

	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("You must specified a command beers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "find by ID")
		beersCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}

}
