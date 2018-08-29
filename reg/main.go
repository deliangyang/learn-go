package main

import (
	"regexp"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"path/filepath"
)

func main() {

	reg := regexp.MustCompile(`\d+`)
	res := reg.Match([]byte("2018"))
	fmt.Println(res)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	b := watcher.Add("./reg")
	if b != nil {
		log.Fatal(b)
	}

	for {
		select {
			case ev := <- watcher.Events:
				fmt.Println(ev)
				if ev.Op & fsnotify.Remove == fsnotify.Remove || ev.Op & fsnotify.Write == fsnotify.Write || ev.Op & fsnotify.Create == fsnotify.Create {
					base := filepath.Base(ev.Name)
					fmt.Println(base)
				}
		}
	}
}
