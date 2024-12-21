//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchAndPrint(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: url: %s, err: %v\n", url, err)
		return
	}
	buf, err := io.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(buf))
}

func main() {
	wg := sync.WaitGroup{}
	counter := atomic.Int32{}
	urls := os.Args[1:]
	for idx, url := range urls {
		wg.Add(1)
		go func(idx int, url string) {
			fetchAndPrint(url)
			counter.Add(1)
			wg.Done()
			fmt.Printf("Done %d%%\n", int(float64(counter.Load())*100/float64(len(urls))))
		}(idx, url)
	}
	wg.Wait()
}
