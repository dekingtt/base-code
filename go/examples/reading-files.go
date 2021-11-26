package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := os.ReadFile("abc.temp")
	checkerr(err)
	fmt.Println(string(data))

	f, err := os.Open("abc.temp")
	checkerr(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	fmt.Printf("%d bytes %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(7, 0)
	checkerr(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	checkerr(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%s\n", string(b2))

	o3, err := f.Seek(7, 0)
	checkerr(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkerr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	checkerr(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	checkerr(err)
	fmt.Printf("5 bytes %s\n", string(b4))
	f.Close()
}
