package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	res := &response{
		Page: 1,
		Fruits: []string {"apple", "peach", "pear"},
	}
	res2, _ := json.Marshal(res)
	fmt.Println(string(res2))
}
