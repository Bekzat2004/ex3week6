package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Title        template.HTML
			BusinessName string
			Slogan       string
		}{
			Title:        template.HTML("Ex3_week3 &verbar; AI & GPT"),
			BusinessName: "Business.",
			Slogan:       "we get things done!",
		}

		err := templates.ExecuteTemplate(w, "index.html", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", index())

	port := "8066"
	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
