package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You requested: %s \n", r.URL.Path)

	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Home page")
	case "/about":
		fmt.Fprint(w, "About Us")
	case "/login":
		fmt.Fprint(w, "Login")
	}
}

func main() {
	http.HandleFunc("/", welcomePage)
	http.ListenAndServe("", nil)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
