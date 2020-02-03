package trie

var (
	suggestionList = []string{}
)

type Trie struct {
	Children [128]*Trie
	isWord   bool
	Value    string
}

func New() *Trie {

	t := &Trie{}
	t.isWord = false
	return t
}

func (t *Trie) Add(s string, value string) {
	if len(s) < 1 {
		t.isWord = true
		t.Value = value
		return
	}

	letter := s[0]

	index := letter

	if t.Children[index] == nil {
		t.Children[index] = New()
	}

	t.Children[index].Add(s[1:], value)
}

func (t *Trie) Search(keyword string) (string, bool) {

	if t == nil {
		return "false", false
	}

	curr := t
	for i := 0; i < len(keyword); i++ {

		letter := keyword[i]
		// fmt.Println(letter)
		index := letter
		// fmt.Println("index", index)
		curr = curr.Children[index]

		if curr == nil {
			return "", false
		}
	}

	return curr.Value, curr.isWord
}

func isLastNode(t *Trie) bool {

	for i := 0; i < 128; i++ {
		if t.Children[i] != nil {
			return false
		}
	}

	return true

}

//TODO return the value to make it testable
func (t *Trie) PrintSuggestion(query string) {

	//move to next position node from the searching character
	for i := 0; i < len(query); i++ {
		index := query[i]

		if t.Children[index] == nil {
			return
		}

		t = t.Children[index]
	}

	if t.isWord && isLastNode(t) {
		return
	}

	wordList := []string{}

	if !isLastNode(t) {
		Suggestion(t, query, &wordList, 4)
	}

}

func Suggestion(t *Trie, prefix string, wordList *[]string, repeat int) *[]string {

	if isLastNode(t) {
		return wordList
	}

	if repeat == 0 {
		return wordList
	}

	for i := 0; i < 128; i++ {

		r := t
		if t.Children[i] != nil {

			l := i
			prefix += string(l)

			r = r.Children[i]

			if r.isWord {
				// fmt.Println(prefix)
				*wordList = append(*wordList, prefix)
			}

			Suggestion(r, prefix, wordList, repeat)
			repeat--

			prefix = prefix[0 : len(prefix)-1]

		}

	}

	return wordList

}
