package squardle

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

// Function that takes in a word and return whether it matches the specified conditions
type WordFilter func(word string, substring string) bool

func WordBegins(word string, substring string) bool {
	if len(substring) > len(word) {
		return false
	}

	return word[:len(substring)] == substring
}

func WordEnds(word string, substring string) bool {
	if len(substring) > len(word) {
		return false
	}
	return word[len(word)-len(substring):] == substring
}

func GetFirstWordFromLine(line string) string {
	line = strings.TrimSpace(line)
	return strings.Split(line, " ")[0]
}

func FilterWordsBySubstring(filename string, filter WordFilter, substring string) ([]string, error) {
	file, err := os.Open("NWL2020.txt")
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	words := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firstWord := strings.ToLower(GetFirstWordFromLine(scanner.Text()))
		if filter(firstWord, substring) {
			words = append(words, firstWord)
		}
	}
	return words, nil
}

func FilterWordsByValidLetters(words []string, letters []byte) []string {
	filteredWords := []string{}
	for _, word := range words {
		if WordContainsOnlyValidLetters(word, letters) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}

func WordContainsOnlyValidLetters(word string, letters []byte) bool {
	valid := true
	for i := 0; i < len(word); i++ {
		if !slices.Contains(letters, word[i]) {
			valid = false
		}
	}
	return valid
}
