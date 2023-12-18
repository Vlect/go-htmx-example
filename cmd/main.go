package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(
		w,
		name,
		data,
	)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Contact struct {
	Name  string
	Email string
}

type ContactList = []Contact

func newContact(name string, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Data struct {
    ContactList ContactList
}

func newData() Data {
    return Data {
        ContactList: ContactList{
            newContact("John", "john@gmail.com"),
            newContact("Clara", "clara@gmail.com"),
        },
    }
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	data := newData()
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
        name := c.FormValue("name")
        email := c.FormValue("email")
        data.ContactList = append(
            data.ContactList,
            newContact(name, email),
        )
		return c.Render(200, "contactList", data)
	})

	e.Logger.Fatal(e.Start(":4206"))
}
