package main

import (
	"os" 
	"time"
	"io/ioutil"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Page struct {
	Path string
	ModTime time.Time
	Body []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Path, p.Body, 0600)
}

func (p *Page) render() []byte {
	unsafe := blackfriday.MarkdownCommon(p.Body)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}

func (p *Page) check_file() {
	file, _ := os.Open(p.Path)
	defer file.Close()
	fstat, _ := file.Stat()
	if fstat.ModTime().After(p.ModTime) {
		var dst []byte
		file.Read(dst)
		p.Body = dst
		p.ModTime = fstat.ModTime()
	}
}

func loadPage(title string) (*Page, error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	
	return &Page{Path: title, Body: body}, nil
}
