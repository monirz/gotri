# Trie Data Struture implementation in Go  


### Usage 

```go
func main() {

	t := trie.New()

	t.Add("ant", "পিপীলিকা")
	t.Add("act", "অভিনয় করা")

	t.Add("car", "গাড়ি")

	meaning, ok := t.search("car")

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

resultArr := Suggestion(tr, "a", &wordList, 3)  
fmt.Println(resultArr)
``` 

Will print `["act", "ant"]` 

The last argument in the `Suggestion()` function is how many words do you want in the suggestion list. 

### Run Test 
```
$ go test 
```


