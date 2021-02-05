package controllers

import (
	"html/template"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		c.tpl.ExecuteTemplate(w, "index.gohtml", nil)
	} else {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found!"))
		return
	}
}

func (c Controller) Resume(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/resume" {
		c.tpl.ExecuteTemplate(w, "resume.gohtml", nil)
	} else {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found!"))
		return
	}
}

func (c Controller) Blog(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/blog" {
		c.tpl.ExecuteTemplate(w, "blog.gohtml", nil)
	} else {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found!"))
		return
	}
}

func (c Controller) Latest(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/latest" {
		c.tpl.ExecuteTemplate(w, "latest.gohtml", nil)
	} else {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found!"))
		return
	}
}

func (c Controller) NotFound(w http.ResponseWriter, req *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		return
	}
}
