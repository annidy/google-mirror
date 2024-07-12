package handler

import (
	"google-mirror/pkg/model"
	"html/template"
	"log"
	"net/http"
	nurl "net/url"
	"sync"
)

type ListItem struct {
	Name    string
	Mirrors []string
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

	cfg := model.MustConfig(model.ReloadConfig())
	h.Data.Items = nil
	wg := sync.WaitGroup{}
	for i := 0; i < len(cfg.Mirrors); i++ {
		wg.Add(1)
		go func(i int) {
			m := make(map[string]any)
			defer wg.Done()
			urls := model.ExtractURLs(cfg.Mirrors[i])
			for j := 0; j < len(urls); j++ {
				m[urls[j]] = struct{}{}
			}
			item := ListItem{
				Name: cfg.Mirrors[i],
			}
			for k := range m {
				item.Mirrors = append(item.Mirrors, k)
			}
			item.Mirrors = item.Mirrors[0:3] // TODO: debug
			h.Data.Items = append(h.Data.Items, item)
		}(i)
	}
	wg.Wait()

	log.Println("handler list")
	tmpl := template.New("list.html").Funcs(template.FuncMap{
		"host": func(url string) string {
			if u, err := nurl.Parse(url); err != nil {
				return url
			} else {
				return u.Host
			}
		}})
	tmpl = template.Must(tmpl.ParseFiles("resource/tmpl/list.html"))

	tmpl.Execute(w, h.Data)
}
