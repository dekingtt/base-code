package main

import (
	"errors"
	"fmt"
)

func f1(v int) (int, error) {
	if v == 42 {
		return -1, errors.New("f1 error")
	}
	return v + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (a argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}
func f2(v int) (int, error) {
	if v == 42 {
		return -1, argError{v, "f2 error"}
	}
	return v + 5, nil
}

func main() {
	for _, n := range []int{7, 42} {
		if r, e := f1(n); e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(r)
		}
	}

	for _, n := range []int{7, 42} {
		if r, e := f2(n); e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
