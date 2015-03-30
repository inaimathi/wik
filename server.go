package main

import (
	"os"
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
	show, _ := template.ParseFiles("static/templates/show.html")
	create, _ := template.ParseFiles("static/templates/create.html")
	flist, _ := template.ParseFiles("static/templates/list.html")
	return func (w http.ResponseWriter, r *http.Request) {
		p, err := wiki.Local(r.URL.Path)
		if err == nil { 
			info, err := os.Stat(p)
			if err == nil && info.IsDir() {
				dir, e := wiki.GetDir(r.URL.Path)
				if e == nil { flist.Execute(w, dir) }
			} else if err == nil {
				pg, e := wiki.GetPage(r.URL.Path)
				pg.ProcessMarkdown()
				if e == nil { show.Execute(w, template.HTML(pg.Body)) }
			} else {
				create.Execute(w, r.URL.Path)
			}
		}
	}
}

func ShowEdit (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	t, _ := template.ParseFiles("static/templates/edit.html")
	return func (w http.ResponseWriter, r *http.Request) {
		pg, err := wiki.GetPage(r.URL.Path[len("/edit/"):])
		if err == nil { t.Execute(w, template.HTML(pg.Raw)) }
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
