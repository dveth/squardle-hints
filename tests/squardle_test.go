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
	filepath := "./NWL2020.txt"
	got, err := squardle.FilterWords(filepath, squardle.WordBegins, substring)

	if err != nil {
		t.Fatalf("In FilterWords, got error: %s", err.Error())
	}
	if !slices.Equal(want, got) {
		t.Logf("Lengths: want %d, got %d\n", len(want), len(got))
		t.Fatalf("In Filterwords, want %v, got %v, for WordBegins and substring %s", want, got, substring)
	}

	substring = "dancing"
	want = []string{"dancing", "outdancing", "ropedancing", "slamdancing", "breakdancing"}
	got, err = squardle.FilterWords(filepath, squardle.WordEnds, substring)
	if err != nil {
		t.Fatalf("In FilterWords, got error: %s", err.Error())
	}
	if !slices.Equal(want, got) {
		t.Logf("Lengths: want %d, got %d\n", len(want), len(got))
		t.Fatalf("In Filterwords, want %v, got %v, for WordEnds and substring %s", want, got, substring)
	}

	substring = "verish"
	want = []string{"feverish", "liverish", "cleverish", "feverishly", "impoverish", "feverishness", "impoverished", "impoverisher", "impoverishes", "liverishness",
		"impoverishers", "impoverishing", "feverishnesses", "impoverishment", "liverishnesses", "impoverishments"}
	got, err = squardle.FilterWords(filepath, strings.Contains, substring)
	if err != nil {
		t.Fatalf("In FilterWords, got error: %s", err.Error())
	}
	if !slices.Equal(want, got) {
		t.Logf("Lengths: want %d, got %d\n", len(want), len(got))
		t.Fatalf("In Filterwords, want %v, got %v, for WordContains and substring %s", want, got, substring)
	}
}
