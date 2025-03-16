package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		ret[field]++
	}
	return ret
}

func main() {
	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	wc.Test(WordCount)
}
