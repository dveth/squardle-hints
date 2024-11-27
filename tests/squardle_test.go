package squardle

import (
	"slices"
	"squardle-hints/internals/squardle"
	"strings"
	"testing"
)

func TestWordBegins(t *testing.T) {
	word := "beginning"
	substring := "be"

	want := true
	got := squardle.WordBegins(word, substring)

	if want != got {
		t.Fatalf("For Word Begins, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}

	substring = "can"
	want = false

	got = squardle.WordBegins(word, substring)
	if want != got {
		t.Fatalf("For Word Begins, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}

	substring = " c"
	want = false

	got = squardle.WordBegins(word, substring)
	if want != got {
		t.Fatalf("For Word Begins, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}

	substring = "deliverance test"
	want = false

	got = squardle.WordBegins(word, substring)
	if want != got {
		t.Fatalf("For Word Begins, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}
}

func TestWordEnds(t *testing.T) {
	word := "ending"
	substring := "ing"

	want := true
	got := squardle.WordEnds(word, substring)

	if want != got {
		t.Fatalf("For Word Ends, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}

	want = false
	substring = "test"
	got = squardle.WordEnds(word, substring)

	if want != got {
		t.Fatalf("For Word Ends, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}

	want = true
	substring = "ending"
	got = squardle.WordEnds(word, substring)

	if want != got {
		t.Fatalf("For Word Ends, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}

	want = false
	substring = "dynamic test"
	got = squardle.WordEnds(word, substring)

	if want != got {
		t.Fatalf("For Word Ends, wanted %v, got %v for word %s and substring %s", want, got, word, substring)
	}
}

func TestGetFirstWordFromLine(t *testing.T) {
	line := "dynamic music for 1 2 3"
	want := "dynamic"

	got := squardle.GetFirstWordFromLine(line)

	if want != got {
		t.Fatalf("For GetFirstWordFromLine, want %s, got %s for line %s", want, got, line)
	}

	line = " dynamic music for 1 2 3"

	got = squardle.GetFirstWordFromLine(line)

	if want != got {
		t.Fatalf("For GetFirstWordFromLine, want %s, got %s for line %s", want, got, line)
	}
}

func TestFilterWords(t *testing.T) {
	want := []string{"beast", "beasts", "beastie", "beastly", "beasties", "beastings", "beastlier", "beastliest", "beastliness", "beastlinesses"}
	substring := "beast"
	filename := "NWL2020.txt"
	words, err := squardle.GetWordList(filename)
	if err != nil {
		t.Fatalf("In GetWordList, got error: %s", err.Error())
	}
	got := squardle.FilterWordsBySubstring(words, squardle.WordBegins, substring)
	if !slices.Equal(want, got) {
		t.Logf("Lengths: want %d, got %d\n", len(want), len(got))
		t.Fatalf("In Filterwords, want %v, got %v, for WordBegins and substring %s", want, got, substring)
	}

	substring = "dancing"
	want = []string{"dancing", "outdancing", "ropedancing", "slamdancing", "breakdancing"}
	got = squardle.FilterWordsBySubstring(words, squardle.WordEnds, substring)
	if !slices.Equal(want, got) {
		t.Logf("Lengths: want %d, got %d\n", len(want), len(got))
		t.Fatalf("In Filterwords, want %v, got %v, for WordEnds and substring %s", want, got, substring)
	}

	substring = "verish"
	want = []string{"feverish", "liverish", "cleverish", "feverishly", "impoverish", "feverishness", "impoverished", "impoverisher", "impoverishes", "liverishness",
		"impoverishers", "impoverishing", "feverishnesses", "impoverishment", "liverishnesses", "impoverishments"}
	got = squardle.FilterWordsBySubstring(words, strings.Contains, substring)
	if !slices.Equal(want, got) {
		t.Logf("Lengths: want %d, got %d\n", len(want), len(got))
		t.Fatalf("In Filterwords, want %v, got %v, for WordContains and substring %s", want, got, substring)
	}
}

func TestWordContainsOnlyValidLetters(t *testing.T) {
	word := "anaconda"
	letters := []byte{'a', 'n', 'c', 'd', 'o'}

	want := true
	got := squardle.WordContainsOnlyValidLetters(word, letters)

	if want != got {
		t.Fatalf("In WordcontainsOnlyValidLetters, wanted %v, got %v, with %s and %v", want, got, word, string(letters))
	}

	letters = []byte{'a', 'n'}
	want = false
	got = squardle.WordContainsOnlyValidLetters(word, letters)

	if want != got {
		t.Fatalf("In WordcontainsOnlyValidLetters, wanted %v, got %v, with %s and %v", want, got, word, string(letters))
	}
}

func TestFilterWordsByValidLetters(t *testing.T) {
	words := []string{"anaconda", "beast", "valid", "test", "golang"}
	letters := []byte{'g', 'o', 'l', 'a', 'n', 'g', 't', 'e', 's'}

	want := []string{"test", "golang"}
	got := squardle.FilterWordsByValidLetters(words, letters)

	if !slices.Equal(want, got) {
		t.Fatalf("In FilterWordsByValidLetters, wanted %v, got %v, with %v and %v", want, got, words, string(letters))
	}

}

func TestFilterWordsByLength(t *testing.T) {
	words := []string{"anaconda", "beast", "valid", "test", "golang"}
	length := 5

	want := []string{"beast", "valid"}
	got := squardle.FilterWordsByLength(words, length)
	if !slices.Equal(want, got) {
		t.Fatalf("In FilterWordsByLength, wanted %v, got %v, with %v and %d", want, got, words, length)
	}
}
