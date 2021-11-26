package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Oragin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%d, name=%s, oragin=%v", p.Id, p.Name, p.Oragin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "coffee"}
	coffee.Oragin = []string{"Ethiopia", "brazil"}

	out, _ := xml.MarshalIndent(coffee, "", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 10, Name: "Tomato"}
	tomato.Oragin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plant   []*Plant `xml:"parent>child>plant"`
	}

	nesting := Nesting{}
	nesting.Plant = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
