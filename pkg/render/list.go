package render

import (
	"google-mirror/pkg/model"
	"html/template"
	"log"
	"net/http"
	nurl "net/url"
)

type ListItem struct {
	Name    string
	Mirrors []model.Mirror
}
type ListData struct {
	Items []ListItem
}

type ListHandler struct {
	Data ListData
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("render list")
	tmpl := template.New("list.html").Funcs(template.FuncMap{
		"host": func(url string) string {
			if u, err := nurl.Parse(url); err != nil {
				return url
			} else {
				return u.Host
			}
		}})
	tmpl = template.Must(tmpl.ParseFiles("resource/tmpl/list.html"))

	// tmpl := template.Must(template.ParseFiles("resource/tmpl/list.html"))
	tmpl.Execute(w, h.Data)
}
