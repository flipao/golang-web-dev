package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "dog.gohtml", nil)
}

func me(res http.ResponseWriter, req *http.Request) {
	var postForm url.Values
	fmt.Println(req.Method)
	if req.Method == "POST" {
		req.ParseForm()
		postForm = req.PostForm
		fmt.Print(postForm)
	}

	tpl.ExecuteTemplate(res, "me.gohtml", postForm)
}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
