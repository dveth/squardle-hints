package squardle

// Function that takes in a word and return whether it matches the specified conditions
type WordFilter interface {
	func(word string, substring string) bool
}

func WordBegins(word string, substring string) bool {
	if len(substring) > len(word) {
		return false
	}

	return word[:len(substring)] == substring
}
