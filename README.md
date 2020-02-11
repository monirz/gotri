# Trie Data Struture implementation in Go  


### Usage 

```go
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
	t.Add("abc", "")

	t.Add("car", "গাড়ি")

	meaning, ok := t.Search("car")

	if !ok {
		fmt.Println("word not found")
		os.Exit(0)
	}

	fmt.Println(meaning)
}

``` 

***Get Suggestion with with searching character*** 
 
 Searching a character `a` with the previous example 
```

tr = tr.Children[97]

wordList := []string{}

resultArr := tr.GetSuggestion("a", 2)  
fmt.Println(resultArr)
``` 

Will print `["abc", "act"]` 

The last argument in the `GetSuggestion()` function is how many words do you want in the suggestion list. 

### Run Test 
```
$ go test 
```


