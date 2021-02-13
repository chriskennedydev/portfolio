package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// TODO: Database logic
func main() {
	http.HandleFunc("/", serveTemplate)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Printf("Server listening on port 5000\n")
	http.ListenAndServe(":5000", nil)

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/images/favicon.ico")
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
		if err != nil {
			log.Println(err.Error())
		}
		err = tpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if file not found
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 for dir request
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
