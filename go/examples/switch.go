package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	default:
		fmt.Println("Afternoon")
	}

	WhatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Println("Unknown")

		}

	}

	WhatAmI(true)
	WhatAmI(5)
	WhatAmI("me")

}
