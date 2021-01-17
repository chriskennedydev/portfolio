package main

import (
	"html/template"
	"net/http"
	"os"
	"portfolio/controllers"

	"github.com/gorilla/handlers"
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

//func loggingFunc(h http.Handler) http.Handler {
//	_, err := os.OpenFile("portfolio.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
//	if err != nil {
//		log.Fatalf("failed to open log: %v", err)
//	}
//
//	return handlers.LoggingHandler(os.Stdout, h)
//}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/images/favicon.ico")
}
