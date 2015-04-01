package main

import (
	"os"
	"net/http"
	"html/template"
	"path/filepath"
	"strings"
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
				if e == nil { show.Execute(w, pg) }
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
		if err == nil { t.Execute(w, pg) }
	}
}

func RemovePage (wiki *Wiki) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		err := wiki.Remove(r.URL.Path[len("/api/remove/"):])
		if err == nil {
			path := r.URL.Path[len("/api/remove"):]
			http.Redirect(w, r, filepath.Dir(path), http.StatusFound)
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

type Crumb struct {
	Name string
	URI string
}

type Trail struct {
	Links []Crumb
	Name string
}

// Breadcrumbs takes a URI and returns the Trail of breadcrumbs that lead to it. 
//   "/"		=> {[{"home" "/"}] ""} // this one should probably be {[] ""} instead
//   "/test.md"		=> {[{"home" "/"}] "test.md"}
//   "/a/b/test.md"	=>   {[{"home" "/"} {"a" "/a"} {"b" "/a/b"}] "test.md"}
//   "/a/b/c/test.md"	=> {[{"home" "/"} {"a" "/a"} {"b" "/a/b"} {"c" "/a/b/c"}] "test.md"}
func Breadcrumbs(uri string) Trail {
	split := strings.Split(uri, "/")
	links := make([]Crumb, 0, len(split)+1)
	links = append(links, Crumb{Name: "home", URI: "/"})
	for ix := range split[:len(split)-1] {
		if split[ix] != "" {
			links = append(links, Crumb{Name: split[ix], URI: strings.Join(split[0:ix+1], "/")})
		}
	}
	if len(split) > 1 {
		return Trail{ Links: links, Name: split[len(split)-1]}
	} else {
		return Trail{ Links: links, Name: ""}
	}
}

// Helper method for templates. Calls Breadcrumbs on the URI of a *Page
func (pg *Page) CrumbsOf() Trail {
	return Breadcrumbs(pg.URI)
}
