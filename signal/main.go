package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	_ = 1 << (1 + iota)
	KiB
	MiB
	X
)

func main() {
	fmt.Println(KiB, MiB, X)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, os.Kill)
	fmt.Println("start")
	s := <-c
	fmt.Println("end", s)
}
