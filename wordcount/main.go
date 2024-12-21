//go:build !solution

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	res := make(map[string]int)
	for _, f := range os.Args[1:] {
		file, err := os.Open(f)
		check(err)

		buf, err := io.ReadAll(file)
		check(err)
		check(file.Close())

		for _, line := range strings.Split(string(buf), "\n") {
			res[line] += 1
		}
	}
	for s, i := range res {
		if i > 1 {
			fmt.Printf("%d\t%s\n", i, s)
		}
	}
}
