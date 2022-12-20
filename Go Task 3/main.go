package main

import (
	"log"
	"net/http"
	"regexp"
)

// Path Validation using Regex
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Handler Function for Router
func makeHandler(handler func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Returns a slice of strings after parsing the URL
		// Sample [/view/home view home]
		urlmatches := validPath.FindStringSubmatch(req.URL.Path)
		if urlmatches == nil {
			http.NotFound(res, req)
			return
		}
		// urlmatches[2] defines the title of the page that we require
		// It calls the corrosponding function accordingly
		handler(res, req, urlmatches[2])
	}
}

func main() {
	// Declare Routes
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// Start Server at 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
