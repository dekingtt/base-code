package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("contains", s.Contains("test", "es"))
	p("count", s.Count("test", "t"))
	p("hasprefix", s.HasPrefix("test", "te"))
	p("hassuffix", s.HasSuffix("test", "st"))
	p("index", s.Index("test", "s"))
	p("join", s.Join([]string{"a", "b"}, "-"))
	p("repeat", s.Repeat("x", 5))
	p("replace", s.Replace("foo", "o", "0", -1))
	p("replace", s.Replace("foo", "o", "0", 1))
	p("split", s.Split("t-e-s-t", "-"))
	p("toupper", s.ToUpper("Test"))
	p("tolower", s.ToLower("Test"))

	p("length", len("hello"))
	p("Char", "hello"[1])
}
