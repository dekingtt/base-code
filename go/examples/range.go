package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	m := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range m {
		fmt.Printf("%v - %v\n", k, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
