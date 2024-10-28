package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error : %v", err)
		return
	}
	fmt.Fprintf(w, "Post request sucessful")
	name := r.FormValue("name")
	fmt.Fprintf(w, "\nName %s", name)
}

func helloFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusNotAcceptable)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fmt.Printf("Hello")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloFunction)

	fmt.Print("statring server at 8080....")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
