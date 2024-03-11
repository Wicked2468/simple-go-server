package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post is a Sucess\n")
	name := r.FormValue("Name")
	address := r.FormValue("Address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not allowed", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Writing HTML response with decoration
	fmt.Fprintf(w, "<!DOCTYPE html>")
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<head>")
	fmt.Fprintf(w, "<title>Hello Page</title>")
	fmt.Fprintf(w, "</head>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
	fmt.Fprintf(w, "<p>Welcome to the Hello Page.</p>")
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Staring Server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
