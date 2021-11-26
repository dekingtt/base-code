package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p := Point{5, 6}
	fmt.Printf("struct1 %v\n", p)
	fmt.Printf("struct2 %+v\n", p)
	fmt.Printf("struct3 %#v\n", p)

	fmt.Printf("type %T\n", p)
	fmt.Printf("bool %t\n", true)
	fmt.Printf("int %d\n", 123)

	fmt.Printf("bin %b\n", 123)
	fmt.Printf("char %c\n", 33)
	fmt.Printf("hex %x\n", 123)

	fmt.Printf("float %f\n", 45.67)
	fmt.Printf("float2 %e\n", 123456.789456)
	fmt.Printf("float3 %E\n", 98765.12348)

	fmt.Printf("string 1 %s\n", "\"string\"")
	fmt.Printf("string 2 %q\n", "\"string\"")
	fmt.Printf("string 3 %x\n", "hex this")

	fmt.Printf("pointer %p\n", &p)

	fmt.Printf("width1 |%6d|%6d|\n", 12, 123)
	fmt.Printf("width2 |%6.2f|%6.2f|\n", 1.2, 1.34)
	fmt.Printf("width3 |%-6.2f|%-6.2f|\n", 1.2, 1.34)
	fmt.Printf("width4 |%6s|%6s|\n", "foo", "so")
	fmt.Printf("width5 |%-6s|%-6s|\n", "foo", "so")

	s := fmt.Sprintf("sprints a %s", "string")
	fmt.Println(s)

}
