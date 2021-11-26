package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkerr(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("aaa.temp", d1, os.ModeAppend)
	checkerr(err)

	f, err := os.Create("aab.temp")
	checkerr(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkerr(err)
	fmt.Printf("write %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkerr(err)
	fmt.Printf("write string %d\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkerr(err)
	fmt.Printf("write buffed %d\n", n4)
	w.Flush()
}
