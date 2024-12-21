//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchAndPrint(url string) {
	resp, err := http.Get(url)
	check(err)
	buf, err := io.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(buf))
}

func main() {
	for _, url := range os.Args[1:] {
		fetchAndPrint(url)
	}
}
