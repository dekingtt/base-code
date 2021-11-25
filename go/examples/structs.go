package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(n string) Person {
	p := Person{name: n}
	p.age = 18
	return p
}

func main() {
	fmt.Println(Person{name: "Alex", age: 16})
	fmt.Println("Tom", 20)

	fmt.Println(&Person{"Will", 8})
	fmt.Println(newPerson("Job"))

	s := Person{"Go", 15}
	sp := &s
	fmt.Println(sp.name)
}
