package main

import (
	"net/http"
)

func ShowPage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := wiki.Render(r.URL.Path)
		if err == nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write(body) 
		}
	}
}

func EditPage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := wiki.Raw(r.URL.Path[len("/edit/"):])
		if err == nil {
			// w.Header().Set("Content-Type")
			w.Write(body)
		}
	}
}

func RemoveAPI (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		err := wiki.Remove(r.URL.Path[len("/api/remove/"):])
		if err == nil {
			http.Redirect(w, r, "/", 303)
		}
	}
}

func CreateAPI (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/api/create/"):]
		err := wiki.Create(path)
		if err == nil {
			http.Redirect(w, r, "/" + path, 303)
		}
	}
}

func EditAPI (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/api/edit/"):]
		r.ParseForm()
		body := r.Form.Get("new_contents")
		err := wiki.Edit(path, []byte(body))
		if err == nil {
			http.Redirect(w, r, "/" + path, 303)
		}
	}
}

func WikiHandlers (wiki *Wiki) {
	http.HandleFunc("/", ShowPage(wiki))
	http.HandleFunc("/edit/", EditPage(wiki))
	http.HandleFunc("/api/remove/", RemoveAPI(wiki))
	http.HandleFunc("/api/edit/", EditAPI(wiki))
	http.HandleFunc("/api/create/", CreateAPI(wiki))
}
