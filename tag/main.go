package main

import (
	"reflect"
	"fmt"
)

type tag struct {
	name string `test:"a" b:"c"`
}

func main() {
	tg := tag{}
	st := reflect.TypeOf(tg)

	field := st.Field(0)
	fmt.Println(field.Tag.Get("test"), field.Tag.Get("b"))

	// callback function
	t("hello world", func(message string) {
		fmt.Println(message)
	})

	message := "hello world"
	defer func(message string) {
		fmt.Println(message)
	}(message)
}

type test func(message string)

func t(s string, tt test) {
	tt(s)
}

func bb(message string)  {
	fmt.Println(message)
}
