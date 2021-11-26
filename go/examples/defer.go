package main

import (
	"fmt"
	"os"
)

func createFile(s string) *os.File {
	fmt.Println("Creating file")
	f, err := os.Create(s)
	if nil != err {
		panic(err.Error())
	}
	return f
}

func writeFile(f *os.File, s string) {
	fmt.Println("Writing to file")
	fmt.Fprintln(f, s)
}

func closeFile(f *os.File) {
	fmt.Println("Closing file")
	err := f.Close()
	if nil != err {
		panic(err.Error())
	}
}

func main() {
	f := createFile("./abc.temp")
	defer closeFile(f)
	writeFile(f, "Hello World!")
}
