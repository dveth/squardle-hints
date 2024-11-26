package routes

import (
	"html/template"
	"io"
	"net/http"

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
}

func RouterStart() {
	e := echo.New()
	tmpl := &Template{
		template.Must(template.ParseGlob(`./views/*.html`)),
	}
	e.Renderer = tmpl

	e.GET("/", indexHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

func indexHandler(c echo.Context) error {
	page := Page{
		PageTitle: "Test title 3",
	}
	c.Render(http.StatusOK, "index", page)
	return nil
}
