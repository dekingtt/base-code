package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(p *int) {
	*p = 0
}

func main() {
	i := 1
	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer", &i)
}
