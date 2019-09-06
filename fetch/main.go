package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	urls := []string{"http://www.baidu.com", "http://jd.com", "http://alipay.com"}
	// ch make channel
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs finished", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d %d	%s", secs, nBytes, resp.StatusCode, url)
}
