package main
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Body []byte
}


func main() {
// This is the call to the main function that handles http request call to view the page
        http.HandleFunc("/view/", viewHandler)
        http.HandleFunc("/edit/", editHandler)
        http.HandleFunc("/save/", handleSave)
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
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/rainbows/"):]
    p, _ := loadPage(title)

    t, _ := template.ParseFiles("view.html")
    t.execute(w, p)
}

// This function loads the page if exits or create an html construct if it dose not
editHandler(w http.ResponseWriter, r *http.Request) {
    title := r,URL.Path[len("/edit/"):]
    p, _ := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }

    t, _ := template.ParseFiles("edit.html")
    t.execute(w, p)
}
