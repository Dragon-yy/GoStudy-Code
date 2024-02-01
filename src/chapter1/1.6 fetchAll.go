package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch(url string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	ch <- fmt.Sprintf("%s response status: %s time:%0.2fs \n", url, resp.Status, time.Since(startTime).Seconds())

}

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
