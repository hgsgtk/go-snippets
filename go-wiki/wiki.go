package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// TODO: template engine/parse imporovement
var templates = template.Must(template.ParseFiles("edit.html", "view.html", "index.html", "admin.html"))

// TODO: routing validation imporovement
var validPath = regexp.MustCompile("^/(edit|save|view|index|admin)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func setSession(userName string, w http.ResponseWriter) error {
	v := map[string]string{
		"name": userName,
	}
	encoded, err := cookieHandler.Encode("session", v)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	return nil
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getUserName(r *http.Request) string {
	var userName string
	if cookie, err := r.Cookie("session"); err != nil {
		val := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &val); err == nil {
			return ""
		}
		println(val)
		userName = val["name"]
	}
	return userName
}

// wiki page
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// wiki page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// wiki page
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &Page{})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		pass := r.FormValue("password")
		rt := "/"
		if name != "" && pass != "" {
			if err := setSession(name, w); err != nil {
				http.Error(w, "Failed to set session", http.StatusInternalServerError)
				return
			}
			rt = "/admin/"
		}
		http.Redirect(w, r, rt, 302)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
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

func adminIndexHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		renderTemplate(w, "admin", &Page{Title: userName})
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/admin/", adminIndexHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
