package goroutineexamples

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func HeadExample() {
	code := make(chan int)
	wg := sync.WaitGroup{}
	start := time.Now()

	for index := range 10 {
		wg.Add(1)
		go func(index int) {
			pingGoogle(index, code)
			wg.Done()
		}(index)
	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for res := range code {
		fmt.Println(res)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}

func pingGoogle(index int, channel chan<- int) {
	start := time.Now()

	resp, err := http.Head("https://google.com")
	if err != nil {
		return
	}

	left := time.Since(start)
	fmt.Printf("%v - %v: status code: %v\n", index, left, resp.StatusCode)

	channel <- resp.StatusCode
}
