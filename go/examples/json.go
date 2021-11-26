package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page  int
	Fruit []string
}

type response2 struct {
	Page  int      `json:"page"`
	Fruit []string `json:"fruit"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(123)
	fmt.Println(string(intB))

	floB, _ := json.Marshal(12.34)
	fmt.Println(string(floB))

	strB, _ := json.Marshal("hello")
	fmt.Println(string(strB))

	arr := []string{"qaz", "wsx", "edc", "rfv"}
	arrB, _ := json.Marshal(arr)
	fmt.Println(string(arrB))

	m := map[string]int{"a": 1, "b": 2}
	mmB, _ := json.Marshal(m)
	fmt.Println(string(mmB))

	res1D := response1{
		Page:  1,
		Fruit: []string{"apple", "orange", "peach"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := response2{
		Page:  2,
		Fruit: []string{"banara", "grape", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13, "strs": ["a", "b", "c"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 3, "fruit":["a", "b", "c"]}`
	res := response2{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Fruit[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 1, "peach": 2}
	enc.Encode(d)
}
