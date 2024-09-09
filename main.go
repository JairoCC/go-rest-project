package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the about page handler

func Home(w http.ResponseWriter, r *http.Request) {
	renderTempalte(w, "home.page.tmpl")
}

// About is the about page handler

func About(w http.ResponseWriter, r *http.Request) {
	renderTempalte(w, "about.page.tmpl")
}

// main is the main application function

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting application on Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

func renderTempalte(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
