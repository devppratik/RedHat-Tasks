package main

import (
	"net/http"
	"text/template"
)

// Describes the model that we use
type Page struct {
	Title string
	Body  []byte
}

// view route handler
func viewHandler(res http.ResponseWriter, req *http.Request, title string) {
	// Load view pages order than home
	if title != "home" {
		page, err := loadPage(title)
		if err != nil {
			http.Redirect(res, req, "/edit/"+title, http.StatusFound)
			return
		}
		renderTemplate(res, "view", page)
		return
	}
	// Load the homepage
	var page, _ = homePage()
	renderTemplate(res, "home", page)
}

// Edit route handler
func editHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(res, "edit", page)
}

// Save route handler
func saveHandler(res http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	err := page.save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/view/"+title, http.StatusFound)
}

// Load all the templates from the template dir
var templates = template.Must(template.ParseFiles("./templates/edit.html", "./templates/view.html", "./templates/home.html"))

// Rendering dynamic templates
func renderTemplate(res http.ResponseWriter, tmpl string, page *Page) {
	err := templates.ExecuteTemplate(res, tmpl+".html", page)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
