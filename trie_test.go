package trie

import "testing"

var testData = []struct {
	engWord  string
	bengWord string
	expected string
}{
	{"ant", "পিপীলিকা", "পিপীলিকা"},
	{"car", "গাড়ি", "গাড়ি"},
	//TODO add more test data
}

var trie = New()

func TestAdd(t *testing.T) {

	for _, v := range testData {
		trie.Add(v.engWord, v.bengWord)

		val, ok := trie.Search(v.engWord)

		if !ok {
			t.Errorf("Expected %v, got %v", true, ok)
		}

		if v.expected != val {
			t.Errorf("Expected %v, got %v", "গাড়ি", val)
		}
	}
}

func TestSearch(t *testing.T) {
	for _, v := range testData {

		val, ok := trie.Search(v.engWord)

		if !ok {
			t.Errorf("Expected %v, got %v", true, ok)
		}

		if v.expected != val {
			t.Errorf("Expected %v, got %v", "গাড়ি", val)
		}
	}
}
