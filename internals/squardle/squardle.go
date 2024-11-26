package squardle

import (
	"bufio"
	"os"
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

func FilterWords(filename string, filter WordFilter, substring string) ([]string, error) {
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
