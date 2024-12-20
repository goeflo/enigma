package main

import (
	"flag"
	"fmt"
)

// input rotors 'I,II,VI'

func main() {

	fmt.Printf("welcome to enigma ...\n")

	var verbose bool
	var rotors string
	var ringPos string
	var plugs string

	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.StringVar(&rotors, "rotors", "I,II,III", "list of used rotors")
	flag.StringVar(&ringPos, "ring", "16,26,08", "ring positions")
	flag.StringVar(&plugs, "plugs", "HZ,YR,IF,QT,JN,GC,AP,UX,BD,KS", "comma separated list of wires for the plugboard")

	flag.Parse()

}
