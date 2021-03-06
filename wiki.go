package main
import (
    "io/ioutil"
    "log"
    "net/http"
    "html/template"
    "regexp"
//     "fmt"
//     "errors"
)

type Page struct {
    Title string
    Body []byte
}


func main() {
// This is the call to the main function that handles http request call to view the page
        http.HandleFunc("/view/", makeHandler(viewHandler))
        http.HandleFunc("/edit/", makeHandler(editHandler))
        http.HandleFunc("/save/", makeHandler(saveHandler))
        log.Fatal(http.ListenAndServe(":8080", nil))
//  CREATING AND READING FROM A txt file
//     p1 := &Page{Title: "WebWithGolang", Body: []byte("Go is the current tornado rocking it's space")}
//     p1.save()
//     p2, _ := loadPage("WebWithGolang")
//     fmt.Println(string(p2.Body))
}


func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage (title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// This function handle a call to an http request
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

// This function loads the page if exits or create an html construct if it dose not
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
    err :=p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, temp string, p *Page) {
    err := templates.ExecuteTemplate(w, temp + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError )
    }
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")


func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract page title from request and call the provided handler fn
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}
