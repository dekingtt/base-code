package main

import "fmt"

type base struct {
	num int
}

func (b base) desc() string {
	return fmt.Sprintf("base with num=%d", b.num)
}

type container struct {
	base
	name string
}

func main() {
	co := container{
		base: base{
			num: 5,
		},
		name: "abc",
	}

	fmt.Println(co)
	fmt.Println(co.desc())

	type describer interface {
		desc() string
	}

	var d describer = co
	fmt.Println(d.desc())
}
