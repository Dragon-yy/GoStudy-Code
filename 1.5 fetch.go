package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch1() {
	// fetch1
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			println("fetch1:", err)
			continue
		}
		text, err := io.ReadAll(resp.Body)

		defer resp.Body.Close()
		if err != nil {
			println("read fetched data failed:", err)
			continue
		}
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Printf("response text: %s\n", text)
	}
}

func main() {
	fetch1()
}
