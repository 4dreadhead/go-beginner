package goroutineexamples

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func PingExample() {
	path := "goroutineexamples/url.txt"
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	urls := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)

	for _, url := range urls {
		go ping(url, respCh, errCh)
	}
	for range len(urls) {
		select {
		case res := <-respCh:
			fmt.Println(res)
		case errRes := <-errCh:
			fmt.Println(errRes.Error())
		}
	}
}

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Head(url)
	if err != nil {
		errCh <- fmt.Errorf("can't ping %s: %v", url, err)
		return
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		errCh <- fmt.Errorf("can't ping %s: status code %v", url, resp.StatusCode)
		return
	}
	respCh <- resp.StatusCode
}
