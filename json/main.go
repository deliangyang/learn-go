package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sync/singleflight"
	"strings"
)

type response struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func (r response) MarshalJSON() ([]byte, error) {
	if r.Fruits == nil {
		return []byte("{}"), nil
	}
	root := make(map[string]interface{})

	root["page"] = r.Page
	root["Fruits"] = strings.Join(r.Fruits, "+")

	return json.Marshal(root)
}


func main() {
	var a = []interface{}{1, 2, 3, 5}
	params := strings.Repeat("%P", len(a))
	fmt.Println(fmt.Sprintf(params, a...))

	var g singleflight.Group
	v, err, _ := g.Do("key", func() (interface{}, error) {
		return "bar", nil
	})
	fmt.Println(v)
	if got, want := fmt.Sprintf("%v (%T)", v, v), "bar (string)"; got != want {
		fmt.Println(fmt.Sprintf("Do = %v; want %v", got, want))
	}
	if err != nil {
		fmt.Println(fmt.Sprintf("Do error = %v", err))
	}
}
