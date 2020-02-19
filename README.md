# Gotri 

[![Coverage](https://gocover.io/_badge/github.com/monirz/gotri)](https://gocover.io/github.com/monirz/trie) [![Actions Status](https://github.com/monirz/gotri/workflows/Build/badge.svg)](https://github.com/monirz/gotri/actions)


**Gotri** is an Unicode character based [Trie/prefix tree](https://en.wikipedia.org/wiki/Trie) implementation in Go, with the suggestion/auto-complete feature for character searching. Since it supports Unicode characters, so word like `café` with the latin `é` would also work for the insertion and searching.      


### Usage 

```go
package main

import (
	"fmt"
	"os"

	"github.com/monirz/trie"
)

func main() {

	t := gotri.New()

	t.Add("ant", "পিপীলিকা")
	t.Add("abc", "letters") 
	t.Add("act", "to behave in the stated way")
	t.Add("and", "used to join two words")
	t.Add("café", "a restaurant where simple and usually quite cheap meals are served")

	t.Add("car", "গাড়ি")

	meaning, ok := t.Search("car")

	if !ok {
		fmt.Println("word not found")
		os.Exit(0)
	}

	fmt.Println(meaning)
}

``` 

## Get Suggestion list with the prefix character 
 
 Searching the character `a` with the inserted words from the previous example. 

```go
resultArr := t.GetSuggestion("a", 2)  
fmt.Println(resultArr)
``` 

Will return an array like this: `["abc", "act"]` 

The last argument in the `GetSuggestion()` function is how many words you want in the suggestion list. 

### Run Test 
```
$ go test -v .
```


