package main

import (
	"bufio"
	"fmt"
	"imura81gt/japaneseTypingGenerator/ja"
	"os"
)

func main() {
	// get stdin
	var keyword string
	var s = bufio.NewScanner(os.Stdin)
	if s.Scan() {
		keyword = s.Text()
	}
	// get yomigana
	var typing = ja.Do(keyword)

	for _, input := range typing {
		fmt.Println(input)
	}

}
