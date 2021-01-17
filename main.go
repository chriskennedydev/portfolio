package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"gitlab.com/kernelhax/portfolio/controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// TODO: Database logic
func main() {
	c := controllers.NewController(tpl)
	index := http.HandlerFunc(c.Index)
	resume := http.HandlerFunc(c.Resume)

	http.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(index)))
	http.Handle("/resume", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(resume)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":5000", nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/images/favicon.ico")
}
