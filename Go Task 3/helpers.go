package main

import (
	"os"
)

var HOME = "./data/homepage.txt"

// Function to Save Files
func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	// Append File Name to Homepage.txt to mainayin list of articles
	f, err := os.OpenFile(HOME, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("<li><a href=\"" + p.Title + "\">" + p.Title + "</a></li>")
	if err != nil {
		return (err)
	}

	// Write contents into the file
	return os.WriteFile(filename, p.Body, 0600)
}

// Function to Read and Load Files
func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Function to Load Home Page
func homePage() (*Page, error) {
	body, err := os.ReadFile(HOME)
	if err != nil {
		return nil, err
	}
	return &Page{Title: "List of all Articles", Body: body}, nil
}
