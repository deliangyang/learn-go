package main

import (
	"fmt"
	"crypto/md5"
	"encoding/base64"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("hello world")
	fmt.Printf("hello world's md5:%x\n", md5.Sum(data))

	base64String := base64.StdEncoding.EncodeToString(data)
	fmt.Println(base64String)

	data, _ = base64.StdEncoding.DecodeString(base64String)
	fmt.Println(string(data))

	err := ioutil.WriteFile("./test11", []byte("hello world"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	content, err := ioutil.ReadFile("./test11")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
