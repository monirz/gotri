package main

import (
	"fmt"
	"os"

	"github.com/monirz/trie"
)

func main() {

	t := trie.New()

	t.Add("ant", "পিপীলিকা")
	t.Add("act", "অভিনয় করা")

	t.Add("car", "গাড়ি")

	meaning, ok := t.Search("car")

	if !ok {
		fmt.Println("word not found")
		os.Exit(0)
	}

	fmt.Println(meaning)
}
