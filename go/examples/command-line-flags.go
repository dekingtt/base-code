package main

import (
	"flag"
	"fmt"
)

func main() {
	wordptr := flag.String("word", "foo", "a string")
	numptr := flag.Int("num", 42, "an int")
	forkpr := flag.Bool("fork", true, "a boolean")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word", *wordptr)
	fmt.Println("num", *numptr)
	fmt.Println("fork", *forkpr)
	fmt.Println("svar", svar)
	fmt.Println("tail", flag.Args())
}
