package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(len(s))

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println(c)

	l := s[2:5]
	fmt.Println("sli1", l)

	l = s[:5]
	fmt.Println("sli2", l)

	l = s[2:]
	fmt.Println("sli3", l)

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD", twoD)
}
