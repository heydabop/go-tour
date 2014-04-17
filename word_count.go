package main

import "code.google.com/p/go-tour/wc"
import "strings"

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		num := m[words[i]]
		m[words[i]] = num + 1
	}
	return m
}

func main(){
	wc.Test(wordCount)
}
