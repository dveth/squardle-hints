package squardle

import (
	"squardle-hints/internals/squardle"
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
