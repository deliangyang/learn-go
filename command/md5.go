package main

import (
	"fmt"
	"flag"
	"crypto/md5"
	"encoding/hex"
)

func main() {
	originName := flag.String("md5", "", "must set value")
	flag.Parse()
	fmt.Println(*originName)
	data := []byte("s")
	fmt.Printf("md5:%x\n", md5.Sum(data))
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	value := md5Ctx.Sum(nil)
	fmt.Printf("%x\n", value)
	fmt.Print(hex.EncodeToString(value))
}