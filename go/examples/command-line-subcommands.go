package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmnd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmnd.Bool("enable", false, "enable")
	fooName := fooCmnd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barlevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expect foo or bar command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmnd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("    enable", *fooEnable)
		fmt.Println("    name", *fooName)
		fmt.Println("    tail", fooCmnd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("    level", *barlevel)
		fmt.Println("    tail", barCmd.Args())
	default:
		fmt.Println("expect foo or bar command")
		os.Exit(1)
	}
}
