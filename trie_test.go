package gotri

import (
	"testing"
)

var testData = []struct {
	engWord  string
	bengWord string
	expected string
}{
	{"ants", "পিপীলিকা", "পিপীলিকা"},
	{"acting", "অভিনয়", "অভিনয়"},
	{"abc", "ad", "ad"},
	{"and", "এবং", "এবং"},
	{"abandon", "ত্যাগ করা", "ত্যাগ করা"},
	{"abandoned", "পরিত্যক্ত", "পরিত্যক্ত"},
	{"abyss", "অতল গহ্বর", "অতল গহ্বর"},
	{"babble", "", ""},
	{"baboon", "", ""},
	{"babbl", "", ""},
	{"book", "", ""},
	{"books", "", ""},
	// {"get up", "গাড়ি", "গাড়ি"},
	// {"work", "কাজ করা", "কাজ করা"},
	//TODO add more test data
}

var testSearchData = []struct {
	engWord  string
	bengWord string
	expected string
	found    bool
}{
	{"ants", "পিপীলিকা", "পিপীলিকা", true},
	{"acting", "অভিনয়", "অভিনয়", true},
	{"abc", "ad", "ad", true},
	{"and", "এবং", "এবং", true},
	{"abandon", "ত্যাগ করা", "ত্যাগ করা", true},
	{"abandoned", "পরিত্যক্ত", "পরিত্যক্ত", true},
	{"abyss", "অতল গহ্বর", "অতল গহ্বর", true},
	{"babble", "", "", true},
	{"baboon", "", "", true},
	{"babbl", "", "", true},
	{"book", "", "", true},
	{"books", "", "", true},
	{"okkk", "", "", false},
}

var testSuggestionData = []struct {
	searchCharacter  string
	total            int
	expectedWordList []string
	expectedCount    int
}{
	{"okkk", 1, []string{}, 0}, //test for item that doesn't exist in the tree
	{"a", 1, []string{"abandon"}, 1},
	{"a", 2, []string{"abandon", "abandoned"}, 2},
	{"b", 2, []string{"babbl", "babble"}, 2},
	{"a", 4, []string{"abandon", "abandoned", "abc", "abyss"}, 4},
	{"a", 5, []string{"abandon", "abandoned", "abc", "abyss", "acting"}, 5},
	{"book", 2, []string{"book", "books"}, 2}, //test for a complete word in the suggestion
	{"books", 1, []string{"books"}, 1},        //test for only one word in the tree
	// {"get", 1, []string{"get up"}, 1},
}

var tr = New()

func TestAdd(t *testing.T) {

	for _, v := range testData {
		tr.Add(v.engWord, v.bengWord)

		val, ok := tr.Search(v.engWord)

		if !ok {
			t.Errorf("Expected %v, got %v", true, ok)
		}

		if v.expected != val {
			t.Errorf("Expected %v, got %v", v.expected, val)
		}
	}

}

func TestSearch(t *testing.T) {

	for _, v := range testSearchData {

		val, ok := tr.Search(v.engWord)
		if v.found != ok {
			t.Errorf("Expected %v, got %v", v.found, ok)
		}

		if v.expected != val {
			t.Errorf("Expected %v, got %v", v.expected, val)
		}
	}

	//test for empty/nil tree
	var tr *Trie
	v, ok := tr.Search("keyword")

	if ok != false {
		t.Errorf("Expected %v got %v", false, ok)
	}

	if v != "" {
		t.Errorf("Expected %v got %v", "", v)
	}
}

func TestSuggestion(t *testing.T) {

	// tr = New()
	for _, val := range testSuggestionData {

		// 	//move to next position node from the searching character
		for i := 0; i < len(val.searchCharacter); i++ {

			index := val.searchCharacter[i]
			if tr.Children[index] == nil {
				return
			}

			tr = tr.Children[index]
		}

		wordList := []string{}

		_, resultArr := Suggestion(tr, val.searchCharacter, wordList, val.total)

		if len(resultArr) != val.expectedCount {
			t.Errorf("Expected %v, got %v", val.expectedCount, len(resultArr))
		}

		for k, v := range resultArr {
			if v != val.expectedWordList[k] {
				t.Errorf("Expected %v, got %v", val.expectedWordList, resultArr)
			}
		}

	}

}

func TestGetSuggestion(t *testing.T) {

	tr = New()

	//add the data to the tree and start from root node
	for _, v := range testData {
		tr.Add(v.engWord, v.bengWord)

		val, ok := tr.Search(v.engWord)

		if !ok {
			t.Errorf("Expected %v, got %v", true, ok)
		}

		if v.expected != val {
			t.Errorf("Expected %v, got %v", v.expected, val)
		}
	}

	for _, val := range testSuggestionData {

		resultArr := tr.GetSuggestion(val.searchCharacter, val.total)

		if len(resultArr) != val.expectedCount {
			t.Errorf("Expected %v, got %v", val.expectedCount, len(resultArr))
		}

		for k, v := range resultArr {
			if v != val.expectedWordList[k] {
				t.Errorf("Expected %v, got %v", val.expectedWordList, resultArr)
			}
		}

	}
}
