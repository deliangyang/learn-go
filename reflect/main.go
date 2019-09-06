package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `user:"age" xxx:"xxxxxxx" size:"1"`
}

func main() {

	value := reflect.ValueOf(User{
		Name: "lily",
		Age:  1,
	})
	fmt.Println(value)
	types := reflect.TypeOf(User{})
	fmt.Println(types.NumField())
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		fmt.Println(field.Name)
		fmt.Println(field.Tag)
		fmt.Println(field.Type)
	}
	if userType, ok := types.FieldByName("Age"); ok {
		fmt.Println(userType.Tag.Get("user"))
		fmt.Println(userType.Tag.Get("xxx"))
		fmt.Println(userType.Tag.Get("size"))
	}
}
