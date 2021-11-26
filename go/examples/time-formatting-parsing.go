package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	p(now.Format(time.RFC3339))

	t1, err := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	if err != nil {
		panic(err)
	}
	p(t1)

	p(now.Format("5.03PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "5 03 AM"

	t2, err := time.Parse(form, "8 12 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(ansic, "8:41PM")
	p(err)
}
