// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

//func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//	t, _ := template.ParseFiles(tmpl + ".html")
//	t.Execute(w, p)
//}

//	func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//		t, err := template.ParseFiles(tmpl + ".html")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		err = t.Execute(w, p)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//	}
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/view/"):]
//	p, _ := loadPage(title)
//	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
//}

//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/view/"):]
//	p, _ := loadPage(title)
//	renderTemplate(w, "view", p)
//}

//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/view/"):]
//	p, err := loadPage(title)
//	if err != nil {
//		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
//		return
//	}
//	renderTemplate(w, "view", p)
//}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

//	func editHandler(w http.ResponseWriter, r *http.Request) {
//		title := r.URL.Path[len("/edit/"):]
//		p, err := loadPage(title)
//		if err != nil {
//			p = &Page{Title: title}
//		}
//		fmt.Fprintf(w, "<h1>Editing %s</h1>"+
//			"<form action=\"/save/%s\" method=\"POST\">"+
//			"<textarea name=\"body\">%s</textarea><br>"+
//			"<input type=\"submit\" value=\"Save\">"+
//			"</form>",
//			p.Title, p.Title, p.Body)
//	}
//func editHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/edit/"):]
//	p, err := loadPage(title)
//	if err != nil {
//		p = &Page{Title: title}
//	}
//	renderTemplate(w, "edit", p)
//}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//func saveHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/save/"):]
//	body := r.FormValue("body")
//	p := &Page{Title: title, Body: []byte(body)}
//	p.save()
//	http.Redirect(w, r, "/view/"+title, http.StatusFound)
//}

//func saveHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/save/"):]
//	body := r.FormValue("body")
//	p := &Page{Title: title, Body: []byte(body)}
//	err := p.save()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	http.Redirect(w, r, "/view/"+title, http.StatusFound)
//}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
