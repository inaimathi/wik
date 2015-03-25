package main

import (
	"net/http"
	"html/template"
)

func WikiHandlers (wiki *Wiki) {
	http.HandleFunc("/", ShowPage(wiki))
	http.HandleFunc("/edit/", ShowEdit(wiki))
	http.HandleFunc("/api/remove/", RemovePage(wiki))
	http.HandleFunc("/api/edit/", EditPage(wiki))
	http.HandleFunc("/api/create/", CreatePage(wiki))
}

func ShowPage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	t, _ := template.ParseFiles("static/templates/show.html")
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := wiki.Render(r.URL.Path)
		if err == nil { t.Execute(w, template.HTML(body)) }
	}
}

func ShowEdit (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	t, _ := template.ParseFiles("static/templates/edit.html")
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := wiki.Raw(r.URL.Path[len("/edit/"):])
		if err == nil { t.Execute(w, template.HTML(body)) }
	}
}

func RemovePage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		err := wiki.Remove(r.URL.Path[len("/api/remove/"):])
		if err == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func CreatePage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/api/create/"):]
		err := wiki.Create(path)
		if err == nil {
			http.Redirect(w, r, "/" + path, http.StatusFound)
		}
	}
}

func EditPage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/api/edit/"):]
		r.ParseForm()
		body := r.Form.Get("new_contents")
		err := wiki.Edit(path, []byte(body))
		if err == nil {
			http.Redirect(w, r, "/" + path, http.StatusFound)
		}
	}
}
