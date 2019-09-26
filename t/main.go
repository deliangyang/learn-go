package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	A map[string]interface{} `json:"a"`
}

type TB map[string]interface{}

type TS struct {
	B TB
}

func main() {
	t := &T{}
	fmt.Println(t.A)
	if t.A == nil {
		fmt.Println("xxxxxbbbbbbbxx")
	}

	tb := TB{}
	fmt.Println(tb)

	var tt TB
	fmt.Println(tt)

	ts := TS{}
	if ts.B == nil {
		fmt.Println("xxxxxxx")
	}
	fmt.Println(ts.B)

	var a map[string]interface{}
	fmt.Println(a)

	s, err := json.Marshal(t)
	if err == nil {
		fmt.Println(string(s))
	}

	t.A = map[string]interface{}{
		"test": "bbbb",
	}
	s, err = json.Marshal(t)
	if err == nil {
		fmt.Println(string(s))
	}

}