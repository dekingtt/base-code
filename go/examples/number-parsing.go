package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0xFDC3", 0, 64)
	fmt.Printf("%X\n", d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("123")
	fmt.Println(k)

	_, err := strconv.Atoi("Wat")
	fmt.Println(err)
}
