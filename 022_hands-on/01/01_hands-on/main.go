package main

import (
	"fmt"
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from home handler!")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog!")
}

func me(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form.Get("name")
	if name == "" {
		name = "you"
	}
	io.WriteString(res, fmt.Sprintf("Hello %s!", name))
}

func main() {

	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
