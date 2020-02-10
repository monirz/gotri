package trie

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
	// {"get up", "গাড়ি", "গাড়ি"},
	// {"work", "কাজ করা", "কাজ করা"},
	//TODO add more test data
}

var testSuggestionData = []struct {
	searchCharacter  string
	total            int
	expectedWordList []string
	expectedCount    int
}{
	{"a", 1, []string{"abandon"}, 1},
	{"a", 2, []string{"abandon", "abandoned"}, 2},
	{"a", 3, []string{"abandon", "abandoned", "abc"}, 3},
	{"a", 4, []string{"abandon", "abandoned", "abc", "abyss"}, 4},
	{"a", 5, []string{"abandon", "abandoned", "abc", "abyss", "and"}, 5},
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
	for _, v := range testData {

		val, ok := tr.Search(v.engWord)
		// fmt.Println(val)
		if !ok {
			t.Errorf("Expected %v, got %v", true, ok)
		}

		if v.expected != val {
			t.Errorf("Expected %v, got %v", "গাড়ি", val)
		}
	}
}

// func TestPrintSuggestion(t *testing.T) {
// 	tr.PrintSuggestion("a")
// }

func TestSuggestion(t *testing.T) {

	for _, val := range testSuggestionData {

		// 	//move to next position node from the searching character
		for i := 0; i < len(val.searchCharacter); i++ {
			index := val.searchCharacter[i]
			// fmt.Println("index", index)
			index = index - 'a'
			if tr.Children[index] == nil {
				return
			}

			tr = tr.Children[index]
		}

		wordList := []string{}

		_, resultArr := Suggestion(tr, "a", &wordList, val.total)

		if len(*resultArr) != val.expectedCount {
			t.Errorf("Expected %v, got %v", val.expectedCount, len(*resultArr))
		}

		for k, v := range *resultArr {
			if v != val.expectedWordList[k] {
				t.Errorf("Expected %v, got %v", val.expectedWordList, resultArr)
			}
		}

	}
	// tr.PrintSuggestion("a")
}
