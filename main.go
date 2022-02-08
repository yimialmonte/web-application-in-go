package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	s := "<h1>Hello, Yelmi Almonte Gau Gau</h1><p>You Age: 200</p><ul><li>Santiago</li><li>Puerto Plata</li><li>Samana</li><li>San Pedro</li></ul>"
	fmt.Fprint(w, s)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To ge in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>FAQ Pages</h1>")
}

func noFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Not Found Page </h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(noFound)

	http.ListenAndServe(":3000", r)
}
