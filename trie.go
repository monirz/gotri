package gotri

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
		return "", false
	}

	curr := t
	for i := 0; i < len(keyword); i++ {

		letter := keyword[i]
		index := letter

		curr = curr.Children[index]

		if curr == nil {
			return "", false
		}
	}

	return curr.Value, curr.isWord
}

//no node has value
//basically end of the node
func isLastNode(t *Trie) bool {

	for i := 0; i < 128; i++ {
		if t.Children[i] != nil {
			return false
		}
	}

	return true

}

//TODO return the value to make it testable
func (t *Trie) GetSuggestion(query string, total int) []string {

	var result []string
	//move to next position node from the searching character
	for i := 0; i < len(query); i++ {
		index := query[i]

		if t.Children[index] == nil {
			return result
		}

		t = t.Children[index]
	}

	//it's a word and has no child node
	//return the search keyword in the array
	if t.isWord && isLastNode(t) {
		result = append(result, query)

		return result
	}

	wordList := []string{}

	//check if the searching word already is a word
	//tappend if true
	if t.isWord {
		wordList = append(wordList, query)
		total--
	}

	if !isLastNode(t) {
		_, result = Suggestion(t, query, wordList, total)
	}

	return result
}

func Suggestion(t *Trie, prefix string, wordList []string, repeat int) (int, []string) {

	if isLastNode(t) {
		return repeat, wordList
	}

	for i := 0; i < 128; i++ {
		if repeat < 1 {
			return repeat, wordList
		}
		r := t
		if t.Children[i] != nil {

			l := i
			prefix += string(l)

			r = r.Children[i]

			if r.isWord {
				wordList = append(wordList, prefix)
				repeat--
			}

			repeat, wordList = Suggestion(r, prefix, wordList, repeat)

			// fmt.Println(prefix)
			prefix = prefix[0 : len(prefix)-1]
		}

	}

	return repeat, wordList

}
