package trie

type Trie struct {
	Children [26]*Trie
	isWord   bool
	Value    string
}

func New() *Trie {
	t := &Trie{}
	t.isWord = false

	return t
}

func (t *Trie) add(s string, value string) {
	if len(s) < 1 {
		t.isWord = true
		t.Value = value
		return
	}

	letter := s[0]

	index := letter - 'a'

	// log.Println(index)
	if t.Children[index] == nil {
		t.Children[index] = New()
	}

	t.Children[index].add(s[1:], value)
}

func (t *Trie) search(keyword string) (string, bool) {

	if t == nil {
		return "false", false
	}

	curr := t
	for i := 0; i < len(keyword); i++ {

		letter := keyword[i]
		index := letter - 'a'
		// fmt.Println("index", index)
		curr = curr.Children[index]

		if curr == nil {
			return "", false
		}
	}

	return curr.Value, curr.isWord
}

func (t *Trie) suggestion(prefix string) (string, bool) {

	if t == nil {
		return "", false
	}

	return "", false

}
