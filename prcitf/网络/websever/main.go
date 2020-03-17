// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func findHandler(w http.ResponseWriter, r *http.Request, title string) {
print("find")
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func Must(t *template.Template, err error) *template.Template {
	if err != nil {
		print(err)
	}
	return t
}

var templates = Must(template.ParseFiles("./edit.html", "twebservice/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("/(edit|save|view|find)/(.*)")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// 找到匹配的第一字符串，并在找到的字符串中继续匹配 正则表达式的子表达式匹配的字符串
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println("url path:", r.URL.Path,"  m",m)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("127.0.0.1/",makeHandler(findHandler))
	http.HandleFunc("10.151.73.175/",makeHandler(findHandler))

	log.Fatal(http.ListenAndServe(":1211", nil))
}

func init(){

	_, err := ioutil.ReadFile("网络/websever/test.txt")
	if err != nil {
		log.Printf("err test:%v",err)
		return
	}
	m := validPath.FindStringSubmatch("/view/1liuhao2/edit/liu2")
	fmt.Println("  m",m)
	log.Println("no error")
	m=regexp.MustCompile("^\\p{Han}+").FindAllString("刘号f",10)

	//runtime.GC()
	fmt.Println(m)
}