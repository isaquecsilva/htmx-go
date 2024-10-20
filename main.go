package main

import (
	"cmp"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Name string
	Type string
}

func sanitize(input string) string {
	input = html.EscapeString(input)
	return strings.TrimSpace(input)
}

func static() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=UTF8")
		w.WriteHeader(http.StatusOK)

		name := cmp.Or(sanitize(r.URL.Query().Get("username")), "Guest")
		userType := cmp.Or(sanitize(r.URL.Query().Get("access")), "guest")

		templ, _ := template.ParseFiles("templates/index.templ")
		templ.Execute(w, User{
			Name: name,
			Type: userType,
		})
	})

	http.HandleFunc("POST /currenttime", func(w http.ResponseWriter, r *http.Request) {
		now :=time.Now().Format(time.DateTime)

		w.Header().Set("Content-Type", "text/html; charset=UTF8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<div class=\"container\">%s</div>", now)
	})

	static()

	if err := http.ListenAndServeTLS(":8000", "cert/cert.pem", "cert/key.pem", nil);
		err != nil {
			log.Fatal(err)
	}
}
