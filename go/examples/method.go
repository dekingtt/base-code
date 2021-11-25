package main

import "fmt"

type rectangle struct {
	width, long int
}

func (r *rectangle) area() int {
	return r.width * r.long
}

func (r *rectangle) perim() int {
	return r.width*2 + r.long*2
}

func main() {
	r := rectangle{width: 10, long: 20}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
