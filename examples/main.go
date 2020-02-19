package main

import (
	"fmt"
	"os"

	"github.com/monirz/gotri"
)

func main() {

	t := gotri.New()

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
