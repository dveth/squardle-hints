package routes

import (
	"html/template"
	"io"
	"net/http"
	"squardle-hints/internals/squardle"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Page struct {
	PageTitle string
	Letters   []string
}

type SquardleData struct {
	Words []string
}

func RouterStart() {
	e := echo.New()
	tmpl := &Template{
		template.Must(template.ParseGlob(`./views/*.html`)),
	}
	e.Static("/css", "/css")
	e.Renderer = tmpl

	e.GET("/", indexHandler)
	e.POST("/wordlist", squardleHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

func indexHandler(c echo.Context) error {
	page := Page{
		PageTitle: "Test title 3",
		Letters:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	}
	c.Render(http.StatusOK, "index", page)
	return nil
}

func squardleHandler(c echo.Context) error {
	data := SquardleData{}
	validLetters := processFormLetters(c)
	wordBeginsSubstring := c.FormValue("wordBeginning")
	wordContainsSubstring := c.FormValue("wordContains")
	wordEndsSubstring := c.FormValue("wordEnds")
	wordLength, err := strconv.Atoi(c.FormValue("wordLength"))

	if err != nil {
		c.Render(http.StatusBadRequest, "words", data)
	}
	wordList, err := squardle.GetWordList("NWL2020.txt")
	if err != nil {
		c.Render(http.StatusInternalServerError, "words", data)
	}

	if wordBeginsSubstring != "" {
		wordList = squardle.FilterWordsBySubstring(wordList, squardle.WordBegins, wordBeginsSubstring)
	}
	if wordContainsSubstring != "" {
		wordList = squardle.FilterWordsBySubstring(wordList, strings.Contains, wordContainsSubstring)
	}
	if wordEndsSubstring != "" {
		wordList = squardle.FilterWordsBySubstring(wordList, squardle.WordEnds, wordEndsSubstring)
	}
	wordList = squardle.FilterWordsByValidLetters(wordList, []byte(validLetters))
	if wordLength != 0 {
		wordList = squardle.FilterWordsByLength(wordList, wordLength)
	}
	data.Words = wordList
	c.Render(http.StatusOK, "words", data)
	return nil
}

func processFormLetters(c echo.Context) []byte {
	lettersString := "abcdefghijlkmnopqrstuvwxyz"
	validLetters := []byte{}
	for _, letter := range lettersString {
		if c.FormValue(string(letter)) != string(letter) {
			validLetters = append(validLetters, byte(letter))
		}
	}
	return validLetters
}
