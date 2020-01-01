# Trie for dictionary 


### Usage 

```go
func main() {

	t := trie.New()

	t.add("ant", "একটি")

	t.add("car", "গাড়ি")

	meaning, ok := t.search("car")

	if !ok {
		fmt.Println("word not found")
		os.Exit(0)
	}

	fmt.Println(meaning)
}

```
