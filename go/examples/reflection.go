package main

import (
	"fmt"
	"reflect"
)

type T []interface{ m() }

func (T) m() {}

func main() {
	type A = [16]byte
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())
	fmt.Println(tc.ChanDir())
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind())
	tx, ty := ta.Elem(), tb.Elem()

	fmt.Println(tx.Kind(), ty.Kind())
	fmt.Println(tx.Bits(), ty.Bits())
	fmt.Println(tx.ConvertibleTo(ty))
	fmt.Println(tb.ConvertibleTo(ta))

	fmt.Println(tb.Comparable())
	fmt.Println(tm.Elem().Comparable())
	fmt.Println(ta.Elem().Comparable())
	fmt.Println(tc.Comparable())

	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind())

	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind())

	type Person struct {
		Name string
		Age  int
	}

	var p Person
	tx = reflect.TypeOf(p)
	fmt.Println(tx.Kind())
	fmt.Println(tx.NumField())
	fmt.Println(tx.Field(1).Name)
	fmt.Println(tx.Field(0).PkgPath)
	fmt.Println(tx.Field(1).PkgPath)
	th, ti := tx.Field(0).Type, tx.Field(1).Type
	fmt.Println(th.Kind())
}
